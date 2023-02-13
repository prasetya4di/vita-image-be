package network

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func GetGoogleVision() (*vision.ImageAnnotatorClient, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(os.Getenv("VISIONCREDENTIALFILE")))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, err
}
