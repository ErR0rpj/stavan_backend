package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var CLIENT *firestore.Client
var CTX context.Context

// Initialzies the firebase with service account details
func InitializeServiceAccountForFirebase() *firestore.Client {
	fmt.Println("Initializing service account for Firebase.")

	sa := option.WithCredentialsFile("service-account-credentials.json")
	app, err := firebase.NewApp(CTX, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	CLIENT, err := app.Firestore(CTX)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Service account initialized.")
	fmt.Println("")

	return CLIENT
}
