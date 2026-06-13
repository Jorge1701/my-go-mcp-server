package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

var (
	validMethods = map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"HEAD":    true,
		"OPTIONS": true,
	}
)

func CurlAPIHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	method, err := request.RequireString("method")
	if err != nil {
		return mcp.NewToolResultError("Missing required 'method' parameter"), nil
	}

	if !isValidMethod(method) {
		return mcp.NewToolResultError("Invalid HTTP method, must be GET, POST, PUT, DELETE, PATCH, HEAD, or OPTIONS"), nil
	}

	url, err := request.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError("Missing required 'url' parameter"), nil
	}

	if !isValidURL(url) {
		return mcp.NewToolResultError("Invalid URL"), nil
	}

	result, err := ExecuteCurl(method, url)
	if err != nil {
		return mcp.NewToolResultError("Error executing call: " + err.Error()), nil
	}

	return mcp.NewToolResultText(result), nil
}

func ExecuteCurl(method, url string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "%s\n", resp.Status)
	fmt.Fprintf(&sb, "%s", respBody)
	return sb.String(), nil
}

func isValidMethod(method string) bool {
	return validMethods[strings.ToUpper(method)]
}

func isValidURL(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	return err == nil
}
