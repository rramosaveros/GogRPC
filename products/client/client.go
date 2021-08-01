package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"../productpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := productpb.NewProductServiceClient(cc)

	// Creating Product
	fmt.Println("-----Creating Product-----\n")

	product := &productpb.Product{
		Name:  "Smartphone YY",
		Price: 20500.50,
	}

	createdProduct, err := c.CreateProduct(context.Background(), &productpb.CreateProductRequest{
		Product: product,
	})

	if err != nil {
		log.Fatalf("Failed to create product %v", err)
	}

	fmt.Printf("Product created %v", createdProduct)

	// Creating Product
	fmt.Println("-----Getting Product-----\n")

	productID := createdProduct.GetProduct().GetId()

	getProductReq := &productpb.GetProductRequest{ProductId: productID}

	getProductRes, getProductErr := c.GetProduct(context.Background(), getProductReq)

	if getProductErr != nil {
		log.Fatalf("Failed to getting product", getProductErr)
	}

	fmt.Printf("Product gotten: %v", getProductRes)

	//update Product
	fmt.Println("-----updating Product-----")

	newProduct := &productpb.Product{
		Id:    productID,
		Name:  "New name : Smartphone",
		Price: 40500,
	}

	updateRes, updateErr := c.UpdateProduct(context.Background(), &productpb.UpdateProductRequest{Product: newProduct})

	if updateErr != nil {
		fmt.Printf("Error happened while getting: %v \n", updateErr)
	}

	fmt.Printf("Product was updated: %v", updateRes)

	//delete Product
	fmt.Println("-----deleting Product-----")

	deleteRes, deleteErr := c.DeleteProduct(context.Background(), &productpb.DeleteProductRequest{ProductId: productID})

	if deleteErr != nil {
		fmt.Printf("Error deleting the product %v", deleteErr)
	}

	fmt.Printf("Product deleted %v: \n", deleteRes.GetProductId())

	//list products
	stream, err := c.ListProduct(context.Background(), &productpb.ListProductRequest{})
	if err != nil {
		log.Fatalf("Error calling list product %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving product %v", err)
		}
		fmt.Println(res.GetProduct())
	}
}
