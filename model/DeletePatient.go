package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeletePatient(client *mongo.Client, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("crud").Collection("patients")

	result, err := collection.DeleteOne(ctx, bson.M{"patientid": id})

	if result.DeletedCount == 0 {
		fmt.Println("Invalid patient id")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Patient %v deleted successfully", id)
	}
}
