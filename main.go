package main

import (
	"context"

	api "main/api"
	config "main/config"
)

func main() {
	//Get firestore client
	config.CTX = context.Background()
	//Initializes the account for firebase
	config.CLIENT = config.InitializeServiceAccountForFirebase()

	api.HandleRoutes()
}
