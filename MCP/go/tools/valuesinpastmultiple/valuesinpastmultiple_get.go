package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/smart-me/mcp-server/config"
	"github.com/smart-me/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Valuesinpastmultiple_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startDate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startDate=%v", val))
		}
		if val, ok := args["endDate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endDate=%v", val))
		}
		if val, ok := args["interval"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("interval=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/ValuesInPastMultiple/%s%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []ValuesData
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateValuesinpastmultiple_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_ValuesInPastMultiple_id",
		mcp.WithDescription("Gets multiple values of a device. This call needs a smart-me professional licence."),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the device")),
		mcp.WithString("startDate", mcp.Required(), mcp.Description("The date when the first value should start")),
		mcp.WithString("endDate", mcp.Required(), mcp.Description("The date when the last value should start")),
		mcp.WithNumber("interval", mcp.Required(), mcp.Description("The interval in minutes betwenn the values. 0 means as fast as possible. Only 1000 values can be get in one call.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Valuesinpastmultiple_getHandler(cfg),
	}
}
