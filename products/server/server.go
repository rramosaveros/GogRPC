package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"../productpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price float64            `bson:"price"`
}

type server struct{}

func (*server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	// parse content and save to mongo
	prod := req.GetProduct()

	data := product{
		Name:  prod.GetName(),
		Price: prod.GetPrice(),
	}

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert OID: %v", err),
		)
	}

	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			Id:    oid.Hex(),
			Name:  prod.GetName(),
			Price: prod.GetPrice(),
		},
	}, nil
}

func (*server) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	productId := req.GetProductId()

	oid, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create empty struct
	data := &product{}

	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find the product: %v", err),
		)
	}

	return &productpb.GetProductResponse{
		Product: dbToProductPb(data),
	}, nil
}

func dbToProductPb(data *product) *productpb.Product {
	return &productpb.Product{
		Id:    data.ID.Hex(),
		Name:  data.Name,
		Price: data.Price,
	}
}

func (*server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	fmt.Println("Update product request")

	prod := req.GetProduct()

	oid, err := primitive.ObjectIDFromHex(prod.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create empty struct
	data := &product{}
	filter := bson.M{"_id": oid}

	//search the product in DB
	res := collection.FindOne(context.Background(), filter)
	if res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find the product with the id %v", err),
		)
	}

	// update the internal struct product
	data.Name = prod.GetName()
	data.Price = prod.GetPrice()

	//update in DB
	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update the product %v", updateErr),
		)
	}
	return &productpb.UpdateProductResponse{
		Product: dbToProductPb(data),
	}, nil
}

func (*server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	fmt.Println("Delete product request")

	oid, err := primitive.ObjectIDFromHex(req.GetProductId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete product %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find the product %v", err),
		)
	}

	return &productpb.DeleteProductResponse{
		ProductId: req.GetProductId(),
	}, nil
}

func (*server) ListProduct(req *productpb.ListProductRequest, stream productpb.ProductService_ListProductServer) error {
	fmt.Println("List products")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &product{}

		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error decoding data product: %v", err),
			)
		}

		stream.Send(&productpb.ListProductResponse{
			Product: dbToProductPb(data),
		})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return nil
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Conexión con mongo db
	fmt.Println("Connecting to Mongo DB")
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017")) //cuando se tiene con usuario y contraseña
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error creating the client DB %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting the client DB %v", err)
	}
	//Fin conexion con mongo db

	collection = client.Database("productsdb").Collection("products")

	fmt.Println("Product service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to liste %v", err)
	}

	s := grpc.NewServer()

	productpb.RegisterProductServiceServer(s, &server{})

	go func() {
		fmt.Println("Strating server ")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// wait for ctrl + x to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block until we get the signal
	<-ch
	fmt.Println("Stoping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing mongo client connection")
	client.Disconnect(context.TODO())
	fmt.Println("Goodbye :D .....")
}
