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

func Picoloadmanagementsetdynamiccurrent_postHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		serialVal, ok := args["serial"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serial"), nil
		}
		serial, ok := serialVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serial"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["current"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("current=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/pico/loadmanagementgroup/current/%s%s", cfg.BaseURL, serial, queryString)
		req, err := http.NewRequest("POST", url, nil)
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
		var result models.Object
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

func CreatePicoloadmanagementsetdynamiccurrent_postTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_pico_loadmanagementgroup_current_serial",
		mcp.WithDescription("Sets the dynamic current of a load management group or a single station."),
		mcp.WithNumber("serial", mcp.Required(), mcp.Description("The serial number can be any pico serial in the group (e.g. 700001)")),
		mcp.WithNumber("current", mcp.Required(), mcp.Description("The dynamic current of the group (in mA)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Picoloadmanagementsetdynamiccurrent_postHandler(cfg),
	}
}
