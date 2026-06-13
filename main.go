package main

import (
	"fmt"
	"shorsh-mcp/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	//fmt.Println(tools.GetTodayMetadata())
	//fmt.Println(tools.GetSearchResults("golan tutorial", 3))
	//fmt.Println(tools.GetPageContent("github.com/mark3labs/mcp-go/mcp"))
	//fmt.Println(tools.ExecuteCurl("GET", "http://localhost:56789/verb/exercise?tense=SIMPLE_PRESENT&pronoun=FIRST_PERSON_SINGULAR&random_limit=1"))
	prepareServer()
}

func prepareServer() {
	s := server.NewMCPServer(
		"Shorsh MCP server tools",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	todayMetadataTool := mcp.NewTool("today_metadata",
		mcp.WithDescription("Returns current date and time"),
	)

	searchResultsTool := mcp.NewTool("search_results",
		mcp.WithDescription("Returns web results"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Query to execute"),
		),
		mcp.WithInteger("limit",
			mcp.Description("How many results to return"),
		),
	)

	getPageContentsTool := mcp.NewTool("get_page_contents",
		mcp.WithDescription("Returns the contents of a web page"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("Exact url of the page"),
			mcp.DefaultNumber(5),
		),
	)

	curlApiTool := mcp.NewTool("curl_api",
		mcp.WithDescription("Performs an API call"),
		mcp.WithString("method",
			mcp.Required(),
			mcp.Description("Must be GET, POST, PUT, DELETE, PATCH, HEAD, or OPTIONS"),
			mcp.DefaultString("GET"),
		),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("Url with http or https, domain, and query params"),
		),
	)

	s.AddTool(todayMetadataTool, tools.TodayMetadataHandler)
	s.AddTool(searchResultsTool, tools.SearchResultsHandler)
	s.AddTool(getPageContentsTool, tools.GetPageContentHandler)
	s.AddTool(curlApiTool, tools.CurlAPIHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
