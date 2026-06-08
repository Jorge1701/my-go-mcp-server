package tools

import (
	"context"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func TodayMetadataHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := GetTodayMetadata()
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(result), nil
}

func GetTodayMetadata() (string, error) {
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	return formattedTime, nil
}
