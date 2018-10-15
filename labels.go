package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func getLabels(path string) map[string][]string {
	result := make(map[string][]string)
	ctx := context.Background()

	client := initClient(ctx)
	filename := extractFileName(path)
	result[filename] = callVisionAPI(ctx, path, client)

	return result
}

// callVisionAPI process Google Vision API call and return labels
func callVisionAPI(ctx context.Context, path string, client *vision.ImageAnnotatorClient) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 5)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	var result []string
	for _, label := range labels {
		result = append(result, label.Description)
	}

	return result
}

// initClient create the client
func initClient(ctx context.Context) *vision.ImageAnnotatorClient {
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

// extractFileName return file name
func extractFileName(path string) string {
	filename := filepath.Base(path)
	result := strings.Split(filename, ".")
	return result[0]
}

// isPathDir return true if the specified path is a directory
func isPathDir(path string) bool {
	var result bool

	fi, err := os.Stat(path)
	if err != nil {
		log.Fatalf("Failed to retrieve file info: %v", err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		result = true
	case mode.IsRegular():
		result = false
	}

	return result
}
