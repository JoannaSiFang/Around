package main

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
)

func annotate(uri string) (float32, error) {
	ctx := context.Background()
	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return 0.0, err
	}
	defer client.Close()

	image := vision.NewImageFromURI(uri)

	faceAnnotations, err := client.DetectFaces(ctx, image, nil, 1)
	if err != nil {
		return 0.0, err
	}
	if len(faceAnnotations) == 0 {
		fmt.Println("No faces found.")
		return 0.0, nil
	}

	return faceAnnotations[0].DetectionConfidence, nil
}
