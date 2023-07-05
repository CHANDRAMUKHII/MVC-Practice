package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Patient struct {
	FirstName       string
	PatientID       string
	LastName        string
	DOB             string
	Gender          string
	ContactNumber   string
	MedicalHistory  string
	DateOfDischarge time.Time
}

func NewPatient(client *mongo.Client, patient *Patient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("crud").Collection("patients")

	_, err := collection.InsertOne(ctx, bson.M{
		"patientid":       patient.PatientID,
		"firstname":       patient.FirstName,
		"lastname":        patient.LastName,
		"dob":             patient.DOB,
		"gender":          patient.Gender,
		"contactnumber":   patient.ContactNumber,
		"medicalhistory":  patient.MedicalHistory,
		"dateofdischarge": patient.DateOfDischarge,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created new patient record successfully!")
	}

}
