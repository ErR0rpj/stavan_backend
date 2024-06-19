package adminConfig

import (
	"fmt"
	"log"
	"main/config"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Initialzies the firebase with service account details
func InitializeServiceAccountForFlutterFirebase() *firestore.Client {
	fmt.Println("Initializing service account for Flutter firebase.")

	sa := option.WithCredentialsFile("flutter-service-account-credentials.json")
	app, err := firebase.NewApp(config.CTX, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(config.CTX)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Service account initialized.")
	fmt.Println("")

	return client
}
