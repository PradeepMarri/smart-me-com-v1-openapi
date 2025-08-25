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

func Devices_postHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DeviceToPost
		
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
		url := fmt.Sprintf("%s/api/Devices", cfg.BaseURL)
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
		var result models.DeviceToPost
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

func CreateDevices_postTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_Devices",
		mcp.WithDescription("Creates or updates a Device or updates it's values."),
		mcp.WithString("MeterSubType", mcp.Description("Input parameter: The Sub Type of this Meter.")),
		mcp.WithString("Id", mcp.Description("Input parameter: The ID of the device")),
		mcp.WithString("VoltageL3", mcp.Description("Input parameter: The Voltage Phase L3 (in V)")),
		mcp.WithString("PowerFactorL3", mcp.Description("Input parameter: The Power Factor (cos phi) Phase L3. Range: 0 - 1")),
		mcp.WithString("VoltageL1", mcp.Description("Input parameter: The Voltage Phase L1 (in V)")),
		mcp.WithString("CounterReadingExport", mcp.Description("Input parameter: The Meter Counter Reading only export")),
		mcp.WithString("CounterReadingT1", mcp.Description("Input parameter: The Meter Counter Reading Tariff 1 in kWh or m3.")),
		mcp.WithString("Name", mcp.Description("Input parameter: The Name of the Device")),
		mcp.WithString("Temperature", mcp.Description("Input parameter: The Temperature (in degree celsius)")),
		mcp.WithString("CounterReadingExportT1", mcp.Description("Input parameter: The Meter Counter Reading only export (Tariff 1)")),
		mcp.WithString("ActivePower", mcp.Description("Input parameter: The Active Power or current flow rate. In kW or m3/h")),
		mcp.WithString("PowerFactor", mcp.Description("Input parameter: The Power Factor (cos phi). Range: 0 - 1")),
		mcp.WithNumber("Serial", mcp.Description("Input parameter: The Serial number")),
		mcp.WithString("Voltage", mcp.Description("Input parameter: The Voltage (in V)")),
		mcp.WithString("CurrentL3", mcp.Description("Input parameter: The Current Phase L3 (in A)")),
		mcp.WithString("Current", mcp.Description("Input parameter: The Current (in A)")),
		mcp.WithString("CurrentL1", mcp.Description("Input parameter: The Current Phase L1 (in A)")),
		mcp.WithString("CurrentL2", mcp.Description("Input parameter: The Current Phase L2 (in A)")),
		mcp.WithString("PowerFactorL2", mcp.Description("Input parameter: The Power Factor (cos phi) Phase L2. Range: 0 - 1")),
		mcp.WithString("ValueDate", mcp.Description("Input parameter: The Date of the Value (in UTC). If this is null the Server Time is used.")),
		mcp.WithString("PowerFactorL1", mcp.Description("Input parameter: The Power Factor (cos phi) Phase L1. Range: 0 - 1")),
		mcp.WithString("CounterReading", mcp.Description("Input parameter: The Meter Counter Reading (Total Energy used) in kWh or m3.")),
		mcp.WithString("CounterReadingT2", mcp.Description("Input parameter: The Meter Counter Reading Tariff 2 in kWh or m3.")),
		mcp.WithString("VoltageL2", mcp.Description("Input parameter: The Voltage Phase L2 (in V)")),
		mcp.WithString("CounterReadingExportT2", mcp.Description("Input parameter: The Meter Counter Reading only export (Tariff 2)")),
		mcp.WithString("DeviceEnergyType", mcp.Description("Input parameter: The Energy Type of this device")),
		mcp.WithBoolean("DigitalInput1", mcp.Description("Input parameter: The digital input number 1")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Devices_postHandler(cfg),
	}
}
