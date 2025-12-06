package mimic

import (
	"context"

	"github.com/ustav/umimic/config"
	"github.com/ustav/umimic/models"
	"github.com/ustav/umimic/openrouter"
)

const (
	openRouterBaseURL = "https://openrouter.ai/api/v1"
)

func SendMessage(ctx context.Context, message string, historyContext []models.Message) (string, error) {
	client := openrouter.NewClient(openRouterBaseURL, openrouter.WithAuth(config.GetOpenRouterAPIKey()))
	return client.ChatCompletion(ctx, message, historyContext)
}
