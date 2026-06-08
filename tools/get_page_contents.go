package tools

import (
	"context"
	"errors"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetPageContentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, err := request.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result, err := GetPageContent(url)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(result), nil
}

func GetPageContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Content not found")
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertReader(resp.Body)
	if err != nil {
		return "", err
	}

	return markdown.String(), err
}
