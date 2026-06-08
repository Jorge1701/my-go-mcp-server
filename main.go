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

	s.AddTool(todayMetadataTool, tools.TodayMetadataHandler)
	s.AddTool(searchResultsTool, tools.SearchResultsHandler)
	s.AddTool(getPageContentsTool, tools.GetPageContentHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
