package dd

import (
	"errors"
	"github.com/zorkian/go-datadog-api"
	"log"
	"os"
)

var Client *datadog.Client

func Init() {
	// Ensure API and APP key are set.
	// They should be exported as DD_API_KEY and DD_APP_KEY respectively.
	apiKey, appKey, err := getKeys()
	if err != nil {
		log.Fatalf("[hound] ERROR: %s\n", err)
	}

	// Create new datadog client
	Client = datadog.NewClient(apiKey, appKey)
}

func getKeys() (string, string, error) {
	apiKey := os.Getenv("DD_API_KEY")
	appKey := os.Getenv("DD_APP_KEY")

	if len(apiKey) == 0 {
		return "", "", errors.New("missing DD_API_KEY environment variable. Export DataDog API key")
	}

	if len(appKey) == 0 {
		return "", "", errors.New("missing DD_APP_KEY environment variable. Export DataDog APP key")
	}
	return apiKey, appKey, nil
}
