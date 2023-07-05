package main

import (
	"fmt"
	"mvcPractice/controller"
	"mvcPractice/model"
)

func main() {
	client, err := model.Connection()
	if err != nil {
		fmt.Println("Failed to connect to database!")
	}
	defer model.DisconnectDB(client)
	controller.SwitchOperation(client)

}
