package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdatePatient(client *mongo.Client) {
	var patientID string
	fmt.Print("Enter the patient ID of the patient to update: ")
	fmt.Scan(&patientID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("crud").Collection("patients")

	filter := bson.M{"patientid": patientID}

	fmt.Print("Enter the field to update (e.g., firstname, lastname, dob, gender, contactnumber, medicalhistory): ")
	var fieldToUpdate string
	fmt.Scan(&fieldToUpdate)

	fmt.Print("Enter the new value: ")
	var newValue string
	fmt.Scan(&newValue)

	update := bson.M{"$set": bson.M{fieldToUpdate: newValue}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("Error updating patient details:", err)
		return
	}

	if result.ModifiedCount == 0 {
		fmt.Println("No matching patient found for the given ID.")
	} else {
		fmt.Println("Patient details updated successfully.")
	}
}
