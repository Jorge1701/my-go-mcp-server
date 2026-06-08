# my-go-mcp-server

A MCP server written in go using [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go). This project provides basic tools for retrieving date/time metadata, searching web results, and fetching content from a given URL.

---

## Features

- **Today Metadata Tool**: Returns the current date and time.
- **Search Results Tool**: Retrieves web search results based on a query using calling searX locally on port 8080.
- **Get Page Contents Tool**: Fetches the contents of a specific web page by its URL.

---

## Build

Run `./build.sh` to generate a `shorsh-mcp` binary.

## Provide

Configure your `mcp.json` file for your local models like so:

```
{
  "mcpServers": {
    "shorsh-mcp": {
      "command": "/full/path/to/binary/shorsh-mcp",
      "args": []
    }
  }
}
```

What if the models are not local? I've no idea :D
