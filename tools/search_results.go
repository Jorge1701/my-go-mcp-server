package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

var (
	excludeDomains = []string{
		"truthsocial.com",
		"youtube.com",
		"facebook.com",
		"instagram.com",
		"x.com",
	}
)

type SearchResult struct {
	Url     string  `json:"url"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Score   float32 `json:"score"`
}

type byScore []SearchResult

func (sr byScore) Len() int           { return len(sr) }
func (sr byScore) Swap(i, j int)      { sr[i], sr[j] = sr[j], sr[i] }
func (sr byScore) Less(i, j int) bool { return sr[i].Score > sr[j].Score }

type SearchResponse struct {
	Query   string         `json:"query"`
	Results []SearchResult `json:"results"`
}

type FinalResult struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func SearchResultsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	limit := request.GetInt("limit", 0)
	if limit < 0 {
		return mcp.NewToolResultError("Limit cannot be negative"), nil
	}

	result, err := GetSearchResults(query, limit)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(result), nil
}

func GetSearchResults(query string, limit int) (string, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("format", "json")

	base, err := url.Parse("http://localhost:8080/search")
	if err != nil {
		return "", err
	}
	fullURL := base.RawQuery + params.Encode()

	fullURL = fmt.Sprintf("%s?%s", base.String(), params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var sr SearchResponse
	err = json.Unmarshal([]byte(body), &sr)
	if err != nil {
		return "", err
	}

	var results []SearchResult
	for _, r := range sr.Results {
		if !shouldIgnore(r.Url) {
			results = append(results, r)
		}
	}

	sort.Sort(byScore(results))
	results = results[:limit]

	var finalResults []FinalResult
	for _, r := range results {
		finalResults = append(finalResults, FinalResult{
			Url:     r.Url,
			Title:   r.Title,
			Content: r.Content,
		})
	}

	json, err := json.Marshal(finalResults)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

func shouldIgnore(url string) bool {
	for _, domainToIgnore := range excludeDomains {
		if strings.Contains(url, domainToIgnore) {
			return true
		}
	}
	return false
}
