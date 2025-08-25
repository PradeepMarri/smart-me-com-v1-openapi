package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/smart-me/mcp-server/config"
	"github.com/smart-me/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Foldersettings_postHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.FolderSettings
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/api/folder/settings/%s", cfg.BaseURL, id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.FolderMenuItem
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

func CreateFoldersettings_postTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_folder_settings_id",
		mcp.WithDescription("Add or edit a folder or a meter. To add a new folder use and empty ID"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the folder or meter to edit. Use and empty ID to add a new folder")),
		mcp.WithString("ValueCorrectionParentFolder", mcp.Description("Input parameter: The value correction on all parent folders. but not on the meter itself")),
		mcp.WithBoolean("Enable", mcp.Description("Input parameter: Flag if the meter is enabled (folder not supported yet)")),
		mcp.WithString("Name", mcp.Description("Input parameter: The Name of the folder or meter")),
		mcp.WithString("ParentFolderId", mcp.Description("Input parameter: The parent folder ID of this item")),
		mcp.WithString("VisualizationName", mcp.Description("Input parameter: The name of the visualization of the folder")),
		mcp.WithString("Description", mcp.Description("Input parameter: The Description of the folder or meter")),
		mcp.WithString("FolderType", mcp.Description("Input parameter: The Type of the folder")),
		mcp.WithNumber("SerialNumber", mcp.Description("Input parameter: The serial number (meter only)")),
		mcp.WithBoolean("UseableForVirtualBillingMeters", mcp.Description("Input parameter: Flag if the meter is usable for virtual billing meters (e.g. washroom)")),
		mcp.WithString("ValueCorrection", mcp.Description("Input parameter: The value correction on this meter")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Foldersettings_postHandler(cfg),
	}
}
