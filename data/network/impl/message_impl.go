package impl

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"log"
	"os"
	"vita-image-service/data/entity"
	"vita-image-service/data/entity/image"
	"vita-image-service/data/network"
)

type messageService struct{}

func NewMessageService() network.MessageService {
	return &messageService{}
}

func (m *messageService) SendImageMessage(message entity.Message) []image.Possibility {
	ctx := context.Background()
	client, err := network.GetGoogleVision()
	if err != nil {
		log.Fatalf("error when init google vision : %v", err)
		return nil
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	filename := os.Getenv("BASEURL") + "public/images/" + message.Message

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	localImage, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	res, err := client.AnnotateImage(ctx, &pb.AnnotateImageRequest{
		Image: localImage,
		Features: []*pb.Feature{
			{Type: pb.Feature_LANDMARK_DETECTION, MaxResults: 25},
			{Type: pb.Feature_DOCUMENT_TEXT_DETECTION, MaxResults: 25},
			{Type: pb.Feature_TEXT_DETECTION, MaxResults: 25},
			{Type: pb.Feature_LOGO_DETECTION, MaxResults: 25},
			{Type: pb.Feature_OBJECT_LOCALIZATION, MaxResults: 25},
		},
	})

	if err != nil {
		log.Fatalf("failed to detect object : %v", err)
	}

	var possibilities []image.Possibility

	if landmarkNotation := res.LandmarkAnnotations[0]; landmarkNotation != nil {
		possibilities = append(possibilities, image.Possibility{
			Type:        pb.Feature_Type_name[int32(pb.Feature_LANDMARK_DETECTION)],
			Description: landmarkNotation.Description,
		})
	}

	if fullTextNotation := res.FullTextAnnotation; fullTextNotation != nil {
		possibilities = append(possibilities, image.Possibility{
			Type:        pb.Feature_Type_name[int32(pb.Feature_TEXT_DETECTION)],
			Description: fullTextNotation.Text,
		})
	}

	if logoNotation := res.LogoAnnotations[0]; logoNotation != nil {
		possibilities = append(possibilities, image.Possibility{
			Type:        pb.Feature_Type_name[int32(pb.Feature_LOGO_DETECTION)],
			Description: logoNotation.Description,
		})
	}

	return possibilities
}
