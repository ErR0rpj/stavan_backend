package admin

import (
	"fmt"
	"main/config"
	"main/internal/models"
)

// This function will write many songs in web firestore with batch function
func WriteMultipleSongs(songs []models.Song) error {
	fmt.Println("Writing multiple songs in batches")
	// Get a new write batch.
	batch := config.CLIENT.BulkWriter(config.CTX)

	for _, song := range songs {
		sfRef := config.CLIENT.Collection("songs").Doc(song.Id)
		//Writing each song in the batch will be written altogether in the end
		_, err := batch.Create(sfRef, song)
		if err != nil {
			return err
		}
	}

	// Commit the batch and write all the changes
	batch.End()
	return nil
}
