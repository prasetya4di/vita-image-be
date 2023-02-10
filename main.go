package main

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"log"
	"os"
)

func init() {
	//Get environment data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(os.Getenv("VISIONCREDENTIALFILE")))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	filename := "upload/images/eiffel.jpg"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	res, err := client.AnnotateImage(ctx, &pb.AnnotateImageRequest{
		Image: image,
		Features: []*pb.Feature{
			{Type: pb.Feature_LANDMARK_DETECTION, MaxResults: 25},
			{Type: pb.Feature_TEXT_DETECTION, MaxResults: 25},
			{Type: pb.Feature_LOGO_DETECTION, MaxResults: 25},
		},
	})

	if err != nil {
		log.Fatalf("failed to detect object : %v", err)
	}

	log.Println(res)

}
