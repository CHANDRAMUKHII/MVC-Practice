package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchPatientByID(client *mongo.Client) {
	var patientID string
	fmt.Print("Enter the patient ID to search: ")
	fmt.Scan(&patientID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("crud").Collection("patients")

	filter := bson.M{"patientid": patientID}

	var patient Patient
	err := collection.FindOne(ctx, filter).Decode(&patient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching patient found for the given ID.")
		} else {
			fmt.Println("Error searching for patient:", err)
		}
		return
	}

	fmt.Println("Patient Details:")
	fmt.Println("Patient ID:", patient.PatientID)
	fmt.Println("First Name:", patient.FirstName)
	fmt.Println("Last Name:", patient.LastName)
	fmt.Println("DOB:", patient.DOB)
	fmt.Println("Gender:", patient.Gender)
	fmt.Println("Contact Number:", patient.ContactNumber)
	fmt.Println("Medical History:", patient.MedicalHistory)
	fmt.Println("Date of Discharge:", patient.DateOfDischarge)
}
