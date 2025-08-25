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

func Oauth_authorizeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_id=%v", val))
		}
		if val, ok := args["redirect_uri"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("redirect_uri=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["scope"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("scope=%v", val))
		}
		if val, ok := args["client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_secret=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/oauth/authorize%s", cfg.BaseURL, queryString)
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

func CreateOauth_authorizeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_oauth_authorize",
		mcp.WithDescription(""),
		mcp.WithString("client_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("redirect_uri", mcp.Required(), mcp.Description("")),
		mcp.WithString("state", mcp.Required(), mcp.Description("")),
		mcp.WithString("scope", mcp.Description("")),
		mcp.WithString("client_secret", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Oauth_authorizeHandler(cfg),
	}
}
