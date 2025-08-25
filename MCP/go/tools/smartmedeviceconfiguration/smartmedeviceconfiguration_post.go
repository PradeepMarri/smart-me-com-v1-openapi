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

func Smartmedeviceconfiguration_postHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.SmartMeDeviceConfigurationContainer
		
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
		url := fmt.Sprintf("%s/api/SmartMeDeviceConfiguration", cfg.BaseURL)
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
		var result map[string]interface{}
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

func CreateSmartmedeviceconfiguration_postTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_SmartMeDeviceConfiguration",
		mcp.WithDescription("Sets the configuration of a smart-me device. The device needs to be online."),
		mcp.WithArray("InputConfiguration", mcp.Description("Input parameter: The configuration for the intput outputs")),
		mcp.WithString("UploadInterval", mcp.Description("Input parameter: Number of seconds the device will upload the data. For smaller values maybe a professional license is needed.")),
		mcp.WithString("DeviceEncryptionKey", mcp.Description("Input parameter: The encryption key used to decrypt messages received from an external meter (used only for the smart-me modules)")),
		mcp.WithString("DevicePinCode", mcp.Description("Input parameter: PIN code to enter on a external meter (e.g. for the FNN meters)")),
		mcp.WithString("DnsUpdateState", mcp.Description("Input parameter: Configuration of the dynamic DNS service. More information: http://wiki.smart-me.com/index.php/Dynamisches_DNS")),
		mcp.WithString("Id", mcp.Description("Input parameter: The ID of the device")),
		mcp.WithArray("OutputConfiguration", mcp.Description("Input parameter: The configuration for the external outputs")),
		mcp.WithBoolean("ShowReactiveEnergy", mcp.Description("Input parameter: Shows the reactive energy values (if the meter supports it).")),
		mcp.WithBoolean("EnableModbusTcp", mcp.Description("Input parameter: Enables or disables Modbus TCP (if the meter supports it).")),
		mcp.WithArray("SwitchConfiguration", mcp.Description("Input parameter: The configuration for the phase switches")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Smartmedeviceconfiguration_postHandler(cfg),
	}
}
