package main

import (
	"context"

	"main/api"
	config "main/config"
)

func main() {
	//Get firestore client
	config.CTX = context.Background()
	//Initializes the account for firebase
	config.CLIENT = config.InitializeServiceAccountForFirebase()

	api.HandleRoutes()

}
