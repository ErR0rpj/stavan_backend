package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Starts the server and starts APIs
func AdminRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/admin/transferDataBetweenFirestore", transferDataBetweenFirestore).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

// Fetches the song and its details
func transferDataBetweenFirestore(w http.ResponseWriter, r *http.Request) {
	err := TransferDataFromFlutterFirestoreToWebFirestore()

	if err != nil {
		log.Default().Println("Throwing 500 Internal Server Error:", err)
		//It happens when the there might be an error in code or the data from the database is not interpreted properly.
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status 200"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
