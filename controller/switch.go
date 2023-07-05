package controller

import (
	"fmt"

	"mvcPractice/model"
	"mvcPractice/view"

	"go.mongodb.org/mongo-driver/mongo"
)

func SwitchOperation(client *mongo.Client) (flag bool) {
	operation := view.GetUserChoice()
	switch operation {
	case "1":
		patient := view.PatientDetails()
		modelPatient := model.Patient{
			FirstName:       patient.FirstName,
			PatientID:       patient.PatientID,
			LastName:        patient.LastName,
			DOB:             patient.DOB,
			Gender:          patient.Gender,
			ContactNumber:   patient.ContactNumber,
			MedicalHistory:  patient.MedicalHistory,
			DateOfDischarge: patient.DateOfDischarge,
		}

		model.NewPatient(client, &modelPatient)
	case "2":
		model.UpdatePatient(client)
	case "3":
		model.SearchPatientByID(client)
	case "4":
		fmt.Print("Delete")
		id := view.GetId()
		model.DeletePatient(client, id)
	case "5":
		return false
	default:
		fmt.Print("Invalid choice!!")
	}
	return true
}
