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

	//Enable this and disable above to use admin APIs
	// admin.AdminRoutes()
}
