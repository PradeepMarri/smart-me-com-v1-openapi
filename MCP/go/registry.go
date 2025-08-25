package main

import (
	"github.com/smart-me/mcp-server/config"
	"github.com/smart-me/mcp-server/models"
	tools_picosettings "github.com/smart-me/mcp-server/tools/picosettings"
	tools_picocharging "github.com/smart-me/mcp-server/tools/picocharging"
	tools_devicesbyenergy "github.com/smart-me/mcp-server/tools/devicesbyenergy"
	tools_virtualmetercalculateformula "github.com/smart-me/mcp-server/tools/virtualmetercalculateformula"
	tools_virtualtariffsstatusforproperty "github.com/smart-me/mcp-server/tools/virtualtariffsstatusforproperty"
	tools_virtualtariff "github.com/smart-me/mcp-server/tools/virtualtariff"
	tools_customdevice "github.com/smart-me/mcp-server/tools/customdevice"
	tools_devices "github.com/smart-me/mcp-server/tools/devices"
	tools_valuesinpast "github.com/smart-me/mcp-server/tools/valuesinpast"
	tools_mbus "github.com/smart-me/mcp-server/tools/mbus"
	tools_accesstoken "github.com/smart-me/mcp-server/tools/accesstoken"
	tools_meterfolderinformation "github.com/smart-me/mcp-server/tools/meterfolderinformation"
	tools_foldersettings "github.com/smart-me/mcp-server/tools/foldersettings"
	tools_usertofolderassign "github.com/smart-me/mcp-server/tools/usertofolderassign"
	tools_metervalues "github.com/smart-me/mcp-server/tools/metervalues"
	tools_virtualtariffconsumption "github.com/smart-me/mcp-server/tools/virtualtariffconsumption"
	tools_actions "github.com/smart-me/mcp-server/tools/actions"
	tools_virtualbillingmeteractive "github.com/smart-me/mcp-server/tools/virtualbillingmeteractive"
	tools_picoloadmanagementgroup "github.com/smart-me/mcp-server/tools/picoloadmanagementgroup"
	tools_folder "github.com/smart-me/mcp-server/tools/folder"
	tools_additionaldeviceinformation "github.com/smart-me/mcp-server/tools/additionaldeviceinformation"
	tools_registerforrealtimeapi "github.com/smart-me/mcp-server/tools/registerforrealtimeapi"
	tools_user "github.com/smart-me/mcp-server/tools/user"
	tools_virtualbillingmeters "github.com/smart-me/mcp-server/tools/virtualbillingmeters"
	tools_oauth "github.com/smart-me/mcp-server/tools/oauth"
	tools_picoenablefixcablelock "github.com/smart-me/mcp-server/tools/picoenablefixcablelock"
	tools_smartmedeviceconfiguration "github.com/smart-me/mcp-server/tools/smartmedeviceconfiguration"
	tools_health "github.com/smart-me/mcp-server/tools/health"
	tools_picoloadmanagementsetdynamiccurrent "github.com/smart-me/mcp-server/tools/picoloadmanagementsetdynamiccurrent"
	tools_pico "github.com/smart-me/mcp-server/tools/pico"
	tools_fastsenddevicevalues "github.com/smart-me/mcp-server/tools/fastsenddevicevalues"
	tools_devicesbysubtype "github.com/smart-me/mcp-server/tools/devicesbysubtype"
	tools_picocharginghistory "github.com/smart-me/mcp-server/tools/picocharginghistory"
	tools_account "github.com/smart-me/mcp-server/tools/account"
	tools_virtualbillingmeterdeactivate "github.com/smart-me/mcp-server/tools/virtualbillingmeterdeactivate"
	tools_devicebyserial "github.com/smart-me/mcp-server/tools/devicebyserial"
	tools_folderassign "github.com/smart-me/mcp-server/tools/folderassign"
	tools_valuesinpastmultiple "github.com/smart-me/mcp-server/tools/valuesinpastmultiple"
	tools_virtualtariffsforproperty "github.com/smart-me/mcp-server/tools/virtualtariffsforproperty"
	tools_values "github.com/smart-me/mcp-server/tools/values"
	tools_subuser "github.com/smart-me/mcp-server/tools/subuser"
	tools_foldermenu "github.com/smart-me/mcp-server/tools/foldermenu"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_picosettings.CreatePicosettings_getTool(cfg),
		tools_picocharging.CreatePicocharging_getTool(cfg),
		tools_devicesbyenergy.CreateDevicesbyenergy_getTool(cfg),
		tools_virtualmetercalculateformula.CreateVirtualmetercalculateformula_getTool(cfg),
		tools_virtualtariffsstatusforproperty.CreateVirtualtariffsstatusforproperty_getTool(cfg),
		tools_virtualtariff.CreateVirtualtariff_getTool(cfg),
		tools_customdevice.CreateCustomdevice_getTool(cfg),
		tools_customdevice.CreateCustomdevice_postTool(cfg),
		tools_devices.CreateGet_api_devices_idTool(cfg),
		tools_devices.CreateDevices_putTool(cfg),
		tools_valuesinpast.CreateValuesinpast_getTool(cfg),
		tools_mbus.CreateMbus_postTool(cfg),
		tools_accesstoken.CreateAccesstoken_putTool(cfg),
		tools_meterfolderinformation.CreateMeterfolderinformation_postTool(cfg),
		tools_foldersettings.CreateFoldersettings_postTool(cfg),
		tools_foldersettings.CreateFoldersettings_deleteTool(cfg),
		tools_foldersettings.CreateFoldersettings_getTool(cfg),
		tools_usertofolderassign.CreateUsertofolderassign_deleteTool(cfg),
		tools_usertofolderassign.CreateUsertofolderassign_postTool(cfg),
		tools_metervalues.CreateMetervalues_getTool(cfg),
		tools_meterfolderinformation.CreateMeterfolderinformation_getTool(cfg),
		tools_virtualtariffconsumption.CreateVirtualtariffconsumption_getTool(cfg),
		tools_actions.CreateActions_postTool(cfg),
		tools_virtualbillingmeteractive.CreateVirtualbillingmeteractive_getTool(cfg),
		tools_virtualbillingmeteractive.CreateVirtualbillingmeteractive_postTool(cfg),
		tools_picoloadmanagementgroup.CreatePicoloadmanagementgroup_getTool(cfg),
		tools_folder.CreateFolder_getTool(cfg),
		tools_additionaldeviceinformation.CreateAdditionaldeviceinformation_getTool(cfg),
		tools_registerforrealtimeapi.CreateRegisterforrealtimeapi_deleteTool(cfg),
		tools_actions.CreateActions_getTool(cfg),
		tools_virtualtariff.CreateGet_api_virtualtariff_idTool(cfg),
		tools_user.CreateUser_deleteTool(cfg),
		tools_user.CreateUser_getTool(cfg),
		tools_virtualbillingmeters.CreateVirtualbillingmeters_getTool(cfg),
		tools_registerforrealtimeapi.CreateRegisterforrealtimeapi_getTool(cfg),
		tools_registerforrealtimeapi.CreateRegisterforrealtimeapi_postTool(cfg),
		tools_oauth.CreatePost_api_oauth_authorizeTool(cfg),
		tools_oauth.CreateOauth_authorizeTool(cfg),
		tools_picoenablefixcablelock.CreatePicoenablefixcablelock_postTool(cfg),
		tools_smartmedeviceconfiguration.CreateSmartmedeviceconfiguration_postTool(cfg),
		tools_devices.CreateDevices_getTool(cfg),
		tools_devices.CreateDevices_postTool(cfg),
		tools_health.CreateHealth_getTool(cfg),
		tools_picoloadmanagementsetdynamiccurrent.CreatePicoloadmanagementsetdynamiccurrent_postTool(cfg),
		tools_picoloadmanagementgroup.CreateGet_api_pico_loadmanagementgroupTool(cfg),
		tools_pico.CreatePico_getTool(cfg),
		tools_fastsenddevicevalues.CreateFastsenddevicevalues_getTool(cfg),
		tools_devicesbysubtype.CreateDevicesbysubtype_getTool(cfg),
		tools_picocharginghistory.CreatePicocharginghistory_getTool(cfg),
		tools_account.CreatePost_api_account_loginTool(cfg),
		tools_account.CreateAccount_loginTool(cfg),
		tools_smartmedeviceconfiguration.CreateSmartmedeviceconfiguration_getTool(cfg),
		tools_virtualbillingmeterdeactivate.CreateVirtualbillingmeterdeactivate_postTool(cfg),
		tools_devicebyserial.CreateDevicebyserial_getTool(cfg),
		tools_folderassign.CreateFolderassign_postTool(cfg),
		tools_customdevice.CreateGet_api_customdevice_idTool(cfg),
		tools_valuesinpastmultiple.CreateValuesinpastmultiple_getTool(cfg),
		tools_virtualtariffsforproperty.CreateVirtualtariffsforproperty_getTool(cfg),
		tools_values.CreateValues_getTool(cfg),
		tools_subuser.CreateSubuser_postTool(cfg),
		tools_foldermenu.CreateFoldermenu_postTool(cfg),
		tools_foldermenu.CreateFoldermenu_getTool(cfg),
		tools_subuser.CreateSubuser_deleteTool(cfg),
		tools_subuser.CreateSubuser_getTool(cfg),
	}
}
