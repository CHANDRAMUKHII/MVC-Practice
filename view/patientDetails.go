package view

import (
	"fmt"
	"time"
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

func PatientDetails() Patient {
	var patient Patient
	fmt.Print("Enter the first name of the patient : ")
	fmt.Scan(&patient.FirstName)
	fmt.Print("Enter the last name of the patient : ")
	fmt.Scan(&patient.LastName)
	fmt.Print("Enter the patient id :")
	fmt.Scan(&patient.PatientID)
	fmt.Print("Enter the DOB :")
	fmt.Scan(&patient.DOB)
	fmt.Print("Enter the gender :")
	fmt.Scan(&patient.Gender)
	fmt.Print("Enter the Contact number : ")
	fmt.Scan(&patient.ContactNumber)
	fmt.Print("Enter the medical history : ")
	fmt.Scan(&patient.MedicalHistory)
	var days int8
	fmt.Print("Enter the number of days to be hospitalized :")
	fmt.Scan(&days)
	date := time.Now().AddDate(0, 0, int(days))
	patient.DateOfDischarge = date
	return patient
}

func GetId() string {
	fmt.Println("Enter the patient id : ")
	var id string
	fmt.Scan(&id)
	return id
}
