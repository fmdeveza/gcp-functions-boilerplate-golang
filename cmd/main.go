package main

import (
	"context"
	"fmt"
	"log"
	"os"

	f "fmdeveza/gcp-functions-boilerplate-golang"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/joho/godotenv/autoload"
)

var PORT string = os.Getenv("PORT")
var GOOGLE_CLOUD_AUTH string = os.Getenv("GOOGLE_CLOUD_AUTH")

func main() {
	LoadGoogleCredentials()
	ctx := context.Background()
	funcframework.RegisterHTTPFunctionContext(ctx, "/", f.Trigger)
	if err := funcframework.Start(PORT); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func LoadGoogleCredentials() {
	root_path, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/google_cloud_auth.json", root_path)
	err := os.WriteFile(filePath, []byte(GOOGLE_CLOUD_AUTH), 0644)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
