package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PicoChargingData represents the PicoChargingData schema from the OpenAPI specification
type PicoChargingData struct {
	Loadsheddingstate string `json:"LoadSheddingState,omitempty"` // Max. dynamic current (e.g. set over API or Modbus TCP) of this station or the loadmanagement group of the station if the station is in a group. in A
	Lastwarningorerrortime string `json:"LastWarningOrErrorTime,omitempty"` // The time when the LastWarningOrError happend
	Connectionmode string `json:"ConnectionMode,omitempty"` // The mode how this station is connected to the cloud
	Maxstationcurrent int `json:"MaxStationCurrent,omitempty"` // Max. current of the station in A
	Maxdynamiccurrent int `json:"MaxDynamicCurrent,omitempty"` // Max. dynamic current (e.g. set over API or Modbus TCP) of this station or the loadmanagement group of the station if the station is in a group. in A
	State string `json:"State,omitempty"` // The state of the charging station
	Lastwarningorerrormessage string `json:"LastWarningOrErrorMessage,omitempty"` // The message of the last warning or error of the station.
	Minstationcurrent int `json:"MinStationCurrent,omitempty"` // Min. current of the station in A
	Maxloadmanagementgroupcurrent int `json:"MaxLoadmanagementGroupCurrent,omitempty"` // Max. current of the loadmanagement group of this station (if there is any) in A
	Rssi int `json:"RSSI,omitempty"` // Received Signal Strength Indicator for the connection mode (wifi or mobile). -127 (min) to 0 (Max)
	Valuedate string `json:"ValueDate,omitempty"` // The date of this values
	Duration int `json:"Duration,omitempty"` // The duration of this charging in seconds
	Maxallowedchargingcurrent int `json:"MaxAllowedChargingCurrent,omitempty"` // Max allowed charging current in A
	Lastwarningorerror string `json:"LastWarningOrError,omitempty"` // The last warning or error of the station. This message is only shown if the warning or error happend in the last 5 minutes.
	Loadmanagementgroupname string `json:"LoadmanagementGroupName,omitempty"` // The name of the loadmanagement group. Or string.empty if the station is not in a group
	Activechargingenergy float64 `json:"ActiveChargingEnergy,omitempty"` // The energy used by this active charging (in kWh)
	Activechargingpower float64 `json:"ActiveChargingPower,omitempty"` // The power of the active charging (in kW)
}

// VMeterToDeactivate represents the VMeterToDeactivate schema from the OpenAPI specification
type VMeterToDeactivate struct {
	Id string `json:"ID,omitempty"` // The ID of the Virtual meter to deactivate
}

// VMeterToActivate represents the VMeterToActivate schema from the OpenAPI specification
type VMeterToActivate struct {
	Serialnumber string `json:"SerialNumber,omitempty"` // The Serialnumber of the Meter to activate.
}

// MeterFolderInformationToPost represents the MeterFolderInformationToPost schema from the OpenAPI specification
type MeterFolderInformationToPost struct {
	Id string `json:"Id,omitempty"` // The ID of the device or folder
	Name string `json:"Name,omitempty"` // Name of the Meter or Folder
}

// SwitchConfigurationContainer represents the SwitchConfigurationContainer schema from the OpenAPI specification
type SwitchConfigurationContainer struct {
	Canswitchoff bool `json:"CanSwitchOff,omitempty"` // Flag if the switch can be turned off or is always on.
	Number int `json:"Number,omitempty"` // The number of the phase. (e.g. 1 for Phase L1)
}

// Object represents the Object schema from the OpenAPI specification
type Object struct {
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	Voltage float64 `json:"Voltage,omitempty"` // The Voltage (in V)
	Switchphasel3on bool `json:"SwitchPhaseL3On,omitempty"` // Flag if the Phase L3 is on on this device.
	Counterreadingt2 float64 `json:"CounterReadingT2,omitempty"` // The Meter Counter Reading Tariff 2
	Currentl2 float64 `json:"CurrentL2,omitempty"` // The Current Phase L2 (in A)
	Currentl1 float64 `json:"CurrentL1,omitempty"` // The Current Phase L1 (in A)
	Flowrate float64 `json:"FlowRate,omitempty"` // The current flow rate (e.g. m3/h)
	Digitaloutput2 bool `json:"DigitalOutput2,omitempty"` // The digital output number 2
	Serial int64 `json:"Serial,omitempty"` // The Serial number
	Voltagel1 float64 `json:"VoltageL1,omitempty"` // The Voltage Phase L1 (in V)
	Voltagel2 float64 `json:"VoltageL2,omitempty"` // The Voltage Phase L2 (in V)
	Metersubtype string `json:"MeterSubType,omitempty"` // The sub meter type (e.g. warmwater or coldwater)
	Counterreadingexport float64 `json:"CounterReadingExport,omitempty"` // The Meter Counter Reading only export
	Analogoutput1 int `json:"AnalogOutput1,omitempty"` // The analog output number 1 (PWM signal) (0 - 32183)
	Powerfactor float64 `json:"PowerFactor,omitempty"` // The Power Factor (cos phi). Range: 0 - 1
	Familytype string `json:"FamilyType,omitempty"` // The Family Type of the device.
	Switchon bool `json:"SwitchOn,omitempty"` // Flag if the Switch is on on this device.
	Chargingstationstate string `json:"ChargingStationState,omitempty"` // The state of a pico charging station. (Only available for pico charging stations)
	Counterreadingt1 float64 `json:"CounterReadingT1,omitempty"` // The Meter Counter Reading Tariff 1
	Counterreadingt4 float64 `json:"CounterReadingT4,omitempty"` // The Meter Counter Reading Tariff 4
	Activetariff int `json:"ActiveTariff,omitempty"` // The Number of the Active Tariff
	Currentl3 float64 `json:"CurrentL3,omitempty"` // The Current Phase L3 (in A)
	Digitalinput1 bool `json:"DigitalInput1,omitempty"` // The digital input number 1
	Digitaloutput1 bool `json:"DigitalOutput1,omitempty"` // The digital output number 1
	Digitalinput2 bool `json:"DigitalInput2,omitempty"` // The digital input number 2
	Activepowerl1 float64 `json:"ActivePowerL1,omitempty"` // The Actvie Power Phase L1
	Activepowerunit string `json:"ActivePowerUnit,omitempty"` // The Unit of the Active Power Value
	Counterreading float64 `json:"CounterReading,omitempty"` // The Meter Counter Reading (Total Energy used)
	Voltagel3 float64 `json:"VoltageL3,omitempty"` // The Voltage Phase L3 (in V)
	Switchphasel2on bool `json:"SwitchPhaseL2On,omitempty"` // Flag if the Phase L2 is on on this device.
	Counterreadingunit string `json:"CounterReadingUnit,omitempty"` // The Unit of the Counter Reading
	Temperature float64 `json:"Temperature,omitempty"` // The Temperature (in degree celsius)
	Powerfactorl2 float64 `json:"PowerFactorL2,omitempty"` // The Power Factor (cos phi) Phase L2. Range: 0 - 1
	Current float64 `json:"Current,omitempty"` // The Current (in A)
	Powerfactorl1 float64 `json:"PowerFactorL1,omitempty"` // The Power Factor (cos phi) Phase L1. Range: 0 - 1
	Deviceenergytype string `json:"DeviceEnergyType,omitempty"` // The Energy Type of this device
	Activepowerl3 float64 `json:"ActivePowerL3,omitempty"` // The Actvie Power Phase L3
	Additionalmeterserialnumber string `json:"AdditionalMeterSerialNumber,omitempty"` // An additional Meter serial number. e.g. the number of a meter a smart-me device is connected to.
	Switchphasel1on bool `json:"SwitchPhaseL1On,omitempty"` // Flag if the Phase L1 is on on this device.
	Valuedate string `json:"ValueDate,omitempty"` // Time of last successful connection the the smart-me Cloud.
	Counterreadingt3 float64 `json:"CounterReadingT3,omitempty"` // The Meter Counter Reading Tariff 3
	Id string `json:"Id,omitempty"` // The ID of the device
	Powerfactorl3 float64 `json:"PowerFactorL3,omitempty"` // The Power Factor (cos phi) Phase L3. Range: 0 - 1
	Activepower float64 `json:"ActivePower,omitempty"` // The Actvie Power or current flow rate
	Activepowerl2 float64 `json:"ActivePowerL2,omitempty"` // The Actvie Power Phase L2
	Analogoutput2 int `json:"AnalogOutput2,omitempty"` // The analog output number 2 (PWM signal) (0 - 32183)
	Counterreadingimport float64 `json:"CounterReadingImport,omitempty"` // The Meter Counter Reading only import
	Name string `json:"Name,omitempty"` // The Name of the Device
}

// PicoLoadmanagementGroupDto represents the PicoLoadmanagementGroupDto schema from the OpenAPI specification
type PicoLoadmanagementGroupDto struct {
	Maxcurrent float64 `json:"MaxCurrent,omitempty"` // The max current of this loadmanagement group
	Name string `json:"Name,omitempty"` // The name of the station
	Numberofstations int `json:"NumberOfStations,omitempty"` // The number of assigned stations
	Id string `json:"Id,omitempty"` // The ID of the loadmanagement group
}

// VirtualTariffsOfFolder represents the VirtualTariffsOfFolder schema from the OpenAPI specification
type VirtualTariffsOfFolder struct {
	Virtualtariffs []VirtualTariff `json:"VirtualTariffs,omitempty"`
	Date string `json:"Date,omitempty"` // The DateTime (UTC) of this virtual tarfifs
	Folderid string `json:"FolderId,omitempty"`
	Name string `json:"Name,omitempty"` // The name of this folder
}

// OutputConfigurationContainer represents the OutputConfigurationContainer schema from the OpenAPI specification
type OutputConfigurationContainer struct {
	S0pulsevalue string `json:"S0PulseValue,omitempty"` // The S0 Pulse Value
	TypeField string `json:"Type,omitempty"` // The Type of the output
	Digitaloutputnoconnectionaction string `json:"DigitalOutputNoConnectionAction,omitempty"` // The Action when the device has lost the connection
	Name string `json:"Name,omitempty"` // The Name of the Output
	Number int `json:"Number,omitempty"` // The number of the Output. (1 for Output 1, 2 for Output 2)
}

// AccessTokenToPut represents the AccessTokenToPut schema from the OpenAPI specification
type AccessTokenToPut struct {
	Cardid int64 `json:"CardId,omitempty"` // The ID of the Card
	Userid int64 `json:"UserId,omitempty"` // The ID of the User. The credentials provided must have permission to edit the user. If no ID is provided, the user in the credentials is taken.
}

// ValuesData represents the ValuesData schema from the OpenAPI specification
type ValuesData struct {
	Date string `json:"Date,omitempty"` // The Date of the Value
	Deviceid string `json:"DeviceId,omitempty"` // The ID of the device
	Values []ValueData `json:"Values,omitempty"` // All values
}

// CustomDeviceToPost represents the CustomDeviceToPost schema from the OpenAPI specification
type CustomDeviceToPost struct {
	Id string `json:"Id,omitempty"` // The ID of the device
	Name string `json:"Name,omitempty"` // The Name of the Device
	Serial int64 `json:"Serial,omitempty"` // The Serial number
	Valuedate string `json:"ValueDate,omitempty"` // The Date of the Value (in UTC). If this is null the Server Time is used.
	Values []CustomDeviceValues `json:"Values,omitempty"` // The Values of the custom Device
}

// InputInformation represents the InputInformation schema from the OpenAPI specification
type InputInformation struct {
	Name string `json:"Name,omitempty"` // The Name of the Input
	Number int `json:"Number,omitempty"` // The Number of this Input. Use this as ID to switch it on or off.
}

// ActionToPost represents the ActionToPost schema from the OpenAPI specification
type ActionToPost struct {
	Actions []ActionToPostItem `json:"Actions,omitempty"` // List with all Actions for this device
	Deviceid string `json:"DeviceID,omitempty"` // The ID of the Device
}

// MeterFolderInformation represents the MeterFolderInformation schema from the OpenAPI specification
type MeterFolderInformation struct {
	Communicationmodulehardwareversion int `json:"CommunicationModuleHardwareVersion,omitempty"` // The Version of the Communication Module (if exists)
	Firmwareversion int `json:"FirmwareVersion,omitempty"` // The Firmware Version of a Meter
	Hardwareversion int `json:"HardwareVersion,omitempty"` // The Hardware Version of a Meter.
	Inputinformations []InputInformation `json:"InputInformations,omitempty"` // Informations about the available Inputs
	Isfolder bool `json:"IsFolder,omitempty"` // Flag if it's a Folder or a Meter
	Name string `json:"Name,omitempty"` // Name of the Meter or Folder
	Outputinformations []OutputInformation `json:"OutputInformations,omitempty"` // Informations about the available Outputs
	Communicationmodulefirmwareversion int `json:"CommunicationModuleFirmwareVersion,omitempty"` // The Version of the Communication Module (if exists)
}

// OutputInformation represents the OutputInformation schema from the OpenAPI specification
type OutputInformation struct {
	Actiontype string `json:"ActionType,omitempty"` // The type of the Output
	Name string `json:"Name,omitempty"` // The Name of the Output
	Number int `json:"Number,omitempty"` // The Number of this Output. Use this as ID to switch it on or off.
	Obiscode string `json:"ObisCode,omitempty"` // The Obis Code for this Output
}

// PicoChargingHistoryData represents the PicoChargingHistoryData schema from the OpenAPI specification
type PicoChargingHistoryData struct {
	Energyused float64 `json:"EnergyUsed,omitempty"` // The energy used (in kWh)
	Starttime string `json:"StartTime,omitempty"` // The starttime of the charging (in UTC)
	Transactionstopreason string `json:"TransactionStopReason,omitempty"`
	Duration int `json:"Duration,omitempty"` // The duration in seconds
}

// VirtualTariffConsumptionData represents the VirtualTariffConsumptionData schema from the OpenAPI specification
type VirtualTariffConsumptionData struct {
	Tarifftype string `json:"TariffType,omitempty"` // The type of the virtual tariff (e.g. solar)
	Consumption float64 `json:"Consumption,omitempty"` // The consumption (e.g. kWh) of this tariff
	Currency string `json:"Currency,omitempty"` // The currency of the price
	Name string `json:"Name,omitempty"` // The Name of this virtual tariff
	Price float64 `json:"Price,omitempty"` // The price of the energy in this timerange
}

// FolderSettings represents the FolderSettings schema from the OpenAPI specification
type FolderSettings struct {
	Valuecorrectionparentfolder float64 `json:"ValueCorrectionParentFolder,omitempty"` // The value correction on all parent folders. but not on the meter itself
	Enable bool `json:"Enable,omitempty"` // Flag if the meter is enabled (folder not supported yet)
	Name string `json:"Name,omitempty"` // The Name of the folder or meter
	Parentfolderid string `json:"ParentFolderId,omitempty"` // The parent folder ID of this item
	Visualizationname string `json:"VisualizationName,omitempty"` // The name of the visualization of the folder
	Description string `json:"Description,omitempty"` // The Description of the folder or meter
	Foldertype string `json:"FolderType,omitempty"` // The Type of the folder
	Serialnumber int64 `json:"SerialNumber,omitempty"` // The serial number (meter only)
	Useableforvirtualbillingmeters bool `json:"UseableForVirtualBillingMeters,omitempty"` // Flag if the meter is usable for virtual billing meters (e.g. washroom)
	Valuecorrection float64 `json:"ValueCorrection,omitempty"` // The value correction on this meter
}

// SubUserData represents the SubUserData schema from the OpenAPI specification
type SubUserData struct {
	Accesstimestartdate string `json:"AccessTimeStartDate,omitempty"` // The start date. From this date the user has access
	Email string `json:"Email,omitempty"` // The Email adress
	Id string `json:"Id,omitempty"` // The ID of the user
	Newpassword string `json:"NewPassword,omitempty"` // If set this is used a new password
	Permissionlevel string `json:"PermissionLevel,omitempty"` // The permission level of the user
	Username string `json:"Username,omitempty"` // The username
	Accessenddate string `json:"AccessEndDate,omitempty"` // The end date. until this date the user has access
}

// RegisterRealtimeApiData represents the RegisterRealtimeApiData schema from the OpenAPI specification
type RegisterRealtimeApiData struct {
	Basicauthusername string `json:"BasicAuthUsername,omitempty"` // The Username (basic auth) of your endpoint. Leave empty of none.
	Id string `json:"Id,omitempty"` // The ID of the registration
	Meterid string `json:"MeterId,omitempty"` // The ID of the Meter. Just used if the RegistrationType is "SingleMeterRegistration".
	Registrationtype string `json:"RegistrationType,omitempty"` // The Type of this registration (per meter, per user, ...)
	Serialnumber string `json:"SerialNumber,omitempty"` // The serial number of the Meter. Just used if the RegistrationType is "SingleMeterRegistration" and the MeterId is null. Example: 1 SME 01 63000000 or 6300000
	Apiurl string `json:"ApiUrl,omitempty"` // The URL of your endpoint. To this endpoint all the values are send to.
	Basicauthpassword string `json:"BasicAuthPassword,omitempty"` // The Password (basic auth) of your endpoint. Leave empty of none.
}

// FolderMenuConfiguration represents the FolderMenuConfiguration schema from the OpenAPI specification
type FolderMenuConfiguration struct {
	Items []FolderMenuItem `json:"Items,omitempty"`
	Browsertimezonename string `json:"BrowserTimeZoneName,omitempty"` // The time zone name taken from the browser
	Browserutctime string `json:"BrowserUtcTime,omitempty"` // The UTC time taken from the browser
}

// CustomDeviceValues represents the CustomDeviceValues schema from the OpenAPI specification
type CustomDeviceValues struct {
	Name string `json:"Name,omitempty"` // The Name of the Value.
	Value float64 `json:"Value,omitempty"` // The Value
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Childusers []User `json:"ChildUsers,omitempty"` // The Users created by this users.
	Email string `json:"Email,omitempty"` // The EMail Address of the User
	Id int64 `json:"Id,omitempty"` // The ID of the User
	Idasstring string `json:"IdAsString,omitempty"` // The ID of the user as string
	Isadmin bool `json:"IsAdmin,omitempty"` // Flag if this User is an Admin User
	Permissions []string `json:"Permissions,omitempty"` // Additional Permissions
	Username string `json:"Username,omitempty"` // The Username of the User
}

// PicoSettingsDto represents the PicoSettingsDto schema from the OpenAPI specification
type PicoSettingsDto struct {
	Cariddetection bool `json:"CarIdDetection,omitempty"` // Flag if the car id detection is enabled
	Internalip string `json:"InternalIp,omitempty"` // The internal IP address
	Modbustcp bool `json:"ModbusTcp,omitempty"` // Flag if ModbusTcp is enabled
	Idleimageurl string `json:"IdleImageUrl,omitempty"` // The url of the idle image
	Mincurrent int `json:"MinCurrent,omitempty"` // The max current of this station (in A)
	Name string `json:"Name,omitempty"` // The name of the station
	Displaybrightness string `json:"DisplayBrightness,omitempty"` // The Brightness of the LCD Matrix display. 0 = minimum, 255 = maximum
	Fixcablelockenable bool `json:"FixCableLockEnable,omitempty"` // Enable the fix lock of the charging cable
	Loadmanagementgroupid string `json:"LoadmanagementGroupId,omitempty"` // The ID of the loadmanagement group
	Maxcurrent int `json:"MaxCurrent,omitempty"` // The max current of this station (in A)
	Authenticationtype string `json:"AuthenticationType,omitempty"` // The authentication type
	Dnsname string `json:"DnsName,omitempty"` // The DNS name of the pico's internal ip
	Serialnumber string `json:"SerialNumber,omitempty"` // The Serial number of the station
}

// VirtualTariff represents the VirtualTariff schema from the OpenAPI specification
type VirtualTariff struct {
	Factor float64 `json:"Factor,omitempty"` // Says how many of the active power is used in this tariff. This is calculated from the last meter values.
	Id string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	TypeField string `json:"Type,omitempty"`
	Unit string `json:"Unit,omitempty"`
	Value float64 `json:"Value,omitempty"` // The Counter Value of this tariff
}

// FolderMenuItem represents the FolderMenuItem schema from the OpenAPI specification
type FolderMenuItem struct {
	Description string `json:"Description,omitempty"` // The Description of the folder
	Name string `json:"Name,omitempty"` // The Name of the item
	Autoexportsettings AutoExportSettings `json:"AutoExportSettings,omitempty"` // Settings for the auto export functionality of a meter
	Foldertype string `json:"FolderType,omitempty"` // The folder type of the item
	Children []FolderMenuItem `json:"Children,omitempty"` // Children folder menu items (sub folder menu items)
	Icon string `json:"Icon,omitempty"` // The path to the Icon of this folder
	Meterserialnumber string `json:"MeterSerialNumber,omitempty"` // The serial number if the folder menu item is a meter. Serial number is handled as a string, as javascript on client side cannot handle long integers properly.
	Id string `json:"Id,omitempty"` // The id of the folder menu item
	Userid string `json:"UserId,omitempty"` // The ID of the user of this folder (only for foldertype user)
}

// ValueData represents the ValueData schema from the OpenAPI specification
type ValueData struct {
	Value float64 `json:"Value,omitempty"` // The Value
	Obis string `json:"Obis,omitempty"` // The Obis code of this value. A description you can find here: http://wiki.smart-me.com/index.php/Obis_codes
}

// FolderData represents the FolderData schema from the OpenAPI specification
type FolderData struct {
	Heatpower float64 `json:"HeatPower,omitempty"` // The Power for Heat (kW)
	Watercountervalue float64 `json:"WaterCounterValue,omitempty"` // The Counter values for Water (m3)
	Waterflowrate float64 `json:"WaterFlowRate,omitempty"` // The Flow Rate for Water (m3/h)
	Electricitycountervalue float64 `json:"ElectricityCounterValue,omitempty"` // The Counter values for electricity (kWh)
	Electricitypower float64 `json:"ElectricityPower,omitempty"` // The Power for electricity (kW)
	Gascountervalue float64 `json:"GasCounterValue,omitempty"` // The Counter values for Gas (m3)
	Gasflowrate float64 `json:"GasFlowRate,omitempty"` // The Flow Rate for Gas (m3/h)
	Heatcountervalue float64 `json:"HeatCounterValue,omitempty"` // The Counter values for Heat (kWh)
}

// MBusData represents the MBusData schema from the OpenAPI specification
type MBusData struct {
	Date string `json:"Date,omitempty"` // The Date of the M-BUS Telegram Readout (in UTC). If this is null the Server Time is used.
	Telegram string `json:"Telegram,omitempty"` // The M-BUS Telegram as Hex string. Example: 68 1F 1F 68 08 02 72 78 56 34 12 24 40 01 07 55 00 00 00 03 13 15 31 00 DA 02 3B 13 01 8B 60 04 37 18 02 18 16
}

// SmartMeDeviceConfigurationContainer represents the SmartMeDeviceConfigurationContainer schema from the OpenAPI specification
type SmartMeDeviceConfigurationContainer struct {
	Devicepincode string `json:"DevicePinCode,omitempty"` // PIN code to enter on a external meter (e.g. for the FNN meters)
	Dnsupdatestate string `json:"DnsUpdateState,omitempty"` // Configuration of the dynamic DNS service. More information: http://wiki.smart-me.com/index.php/Dynamisches_DNS
	Id string `json:"Id,omitempty"` // The ID of the device
	Outputconfiguration []OutputConfigurationContainer `json:"OutputConfiguration,omitempty"` // The configuration for the external outputs
	Showreactiveenergy bool `json:"ShowReactiveEnergy,omitempty"` // Shows the reactive energy values (if the meter supports it).
	Enablemodbustcp bool `json:"EnableModbusTcp,omitempty"` // Enables or disables Modbus TCP (if the meter supports it).
	Switchconfiguration []SwitchConfigurationContainer `json:"SwitchConfiguration,omitempty"` // The configuration for the phase switches
	Inputconfiguration []InputConfigurationContainer `json:"InputConfiguration,omitempty"` // The configuration for the intput outputs
	Uploadinterval string `json:"UploadInterval,omitempty"` // Number of seconds the device will upload the data. For smaller values maybe a professional license is needed.
	Deviceencryptionkey string `json:"DeviceEncryptionKey,omitempty"` // The encryption key used to decrypt messages received from an external meter (used only for the smart-me modules)
}

// InputConfigurationContainer represents the InputConfigurationContainer schema from the OpenAPI specification
type InputConfigurationContainer struct {
	Ontext string `json:"OnText,omitempty"` // The visualization text for an ON action
	TypeField string `json:"Type,omitempty"` // The Type of the output
	Name string `json:"Name,omitempty"` // The Name of the Input
	Number int `json:"Number,omitempty"` // The number of the Input. (1 for Input 1)
	Offtext string `json:"OffText,omitempty"` // The visualization text for an OFF action
}

// AutoExportSettings represents the AutoExportSettings schema from the OpenAPI specification
type AutoExportSettings struct {
	Meterpointid string `json:"MeterPointId,omitempty"` // The meter point id set by the user
	Uploadtype string `json:"UploadType,omitempty"` // The upload type
	Exportformat string `json:"ExportFormat,omitempty"` // The export format
	Exportinterval string `json:"ExportInterval,omitempty"` // The export interval of the auto export
}

// ActionToPostItem represents the ActionToPostItem schema from the OpenAPI specification
type ActionToPostItem struct {
	Obiscode string `json:"ObisCode,omitempty"` // The ObisCode (ID) of the Action
	Value float64 `json:"Value,omitempty"` // The Value to Post
}

// ActionInformation represents the ActionInformation schema from the OpenAPI specification
type ActionInformation struct {
	Name string `json:"Name,omitempty"` // The Name of this action
	Obiscode string `json:"ObisCode,omitempty"` // The Obis Code of this action. This is used as ID.
	Actiontype string `json:"ActionType,omitempty"` // The Type of this action.
	Maxvalue float64 `json:"MaxValue,omitempty"` // The maximum value of this action (e.g. for an AnalogAction)
	Minvalue float64 `json:"MinValue,omitempty"` // The minimum value of this action (e.g. for an AnalogAction)
}

// AdditionalDeviceInformation represents the AdditionalDeviceInformation schema from the OpenAPI specification
type AdditionalDeviceInformation struct {
	Additionalmeterserialnumber string `json:"AdditionalMeterSerialNumber,omitempty"` // An additional Meter serial number. e.g. the number of a meter a smart-me device is connected to.
	Firmwareversion int `json:"FirmwareVersion,omitempty"` // The Firmware Version of a Meter
	Hardwareversion int `json:"HardwareVersion,omitempty"` // The Hardware Version of a Meter.
	Id string `json:"ID,omitempty"` // The ID of the device
	Networkconnection string `json:"NetworkConnection,omitempty"` // The mode how the device is connected to the network. Valid values are: wifi, gprs, ltecatm1, ltenbiot, meshmobile, meshwifi
	Networkconnectionrssi int `json:"NetworkConnectionRSSI,omitempty"` // The connection RSSI value (0 is bad, 100 is best)
}

// DeviceInPast represents the DeviceInPast schema from the OpenAPI specification
type DeviceInPast struct {
	Counterreadingexportt1 float64 `json:"CounterReadingExportT1,omitempty"` // The Meter Counter Reading Export Tariff 1
	Serial int64 `json:"Serial,omitempty"` // The Serial number
	Counterreadingexportt2 float64 `json:"CounterReadingExportT2,omitempty"` // The Meter Counter Reading Export Tariff 2
	Counterreadingexportt4 float64 `json:"CounterReadingExportT4,omitempty"` // The Meter Counter Reading Export Tariff 4
	Counterreadingimport float64 `json:"CounterReadingImport,omitempty"` // The Meter Counter Reading Import
	Counterreadingimportt1 float64 `json:"CounterReadingImportT1,omitempty"` // The Meter Counter Reading Import Tariff 1
	Counterreadingexport float64 `json:"CounterReadingExport,omitempty"` // The Meter Counter Reading Export
	Counterreadingimportt3 float64 `json:"CounterReadingImportT3,omitempty"` // The Meter Counter Reading Import Tariff 3
	Counterreadingt3 float64 `json:"CounterReadingT3,omitempty"` // The Meter Counter Reading Tariff 3
	Counterreadingt4 float64 `json:"CounterReadingT4,omitempty"` // The Meter Counter Reading Tariff 4
	Counterreading float64 `json:"CounterReading,omitempty"` // The Meter Counter Reading (Total Energy used)
	Counterreadingexportt3 float64 `json:"CounterReadingExportT3,omitempty"` // The Meter Counter Reading Export Tariff 3
	Counterreadingimportt2 float64 `json:"CounterReadingImportT2,omitempty"` // The Meter Counter Reading Import Tariff 2
	Counterreadingimportt4 float64 `json:"CounterReadingImportT4,omitempty"` // The Meter Counter Reading Import Tariff 4
	Counterreadingunit string `json:"CounterReadingUnit,omitempty"` // The Unit of the Counter Reading
	Id string `json:"Id,omitempty"` // The ID of the device
	Counterreadingt2 float64 `json:"CounterReadingT2,omitempty"` // The Meter Counter Reading Tariff 2
	Counterreadingt1 float64 `json:"CounterReadingT1,omitempty"` // The Meter Counter Reading Tariff 1
	Date string `json:"Date,omitempty"` // The Date of the Values
}

// DeviceToPost represents the DeviceToPost schema from the OpenAPI specification
type DeviceToPost struct {
	Counterreadingt1 float64 `json:"CounterReadingT1,omitempty"` // The Meter Counter Reading Tariff 1 in kWh or m3.
	Name string `json:"Name,omitempty"` // The Name of the Device
	Temperature float64 `json:"Temperature,omitempty"` // The Temperature (in degree celsius)
	Counterreadingexportt1 float64 `json:"CounterReadingExportT1,omitempty"` // The Meter Counter Reading only export (Tariff 1)
	Activepower float64 `json:"ActivePower,omitempty"` // The Active Power or current flow rate. In kW or m3/h
	Powerfactor float64 `json:"PowerFactor,omitempty"` // The Power Factor (cos phi). Range: 0 - 1
	Serial int64 `json:"Serial,omitempty"` // The Serial number
	Voltage float64 `json:"Voltage,omitempty"` // The Voltage (in V)
	Currentl3 float64 `json:"CurrentL3,omitempty"` // The Current Phase L3 (in A)
	Current float64 `json:"Current,omitempty"` // The Current (in A)
	Currentl1 float64 `json:"CurrentL1,omitempty"` // The Current Phase L1 (in A)
	Currentl2 float64 `json:"CurrentL2,omitempty"` // The Current Phase L2 (in A)
	Powerfactorl2 float64 `json:"PowerFactorL2,omitempty"` // The Power Factor (cos phi) Phase L2. Range: 0 - 1
	Valuedate string `json:"ValueDate,omitempty"` // The Date of the Value (in UTC). If this is null the Server Time is used.
	Powerfactorl1 float64 `json:"PowerFactorL1,omitempty"` // The Power Factor (cos phi) Phase L1. Range: 0 - 1
	Counterreading float64 `json:"CounterReading,omitempty"` // The Meter Counter Reading (Total Energy used) in kWh or m3.
	Counterreadingt2 float64 `json:"CounterReadingT2,omitempty"` // The Meter Counter Reading Tariff 2 in kWh or m3.
	Voltagel2 float64 `json:"VoltageL2,omitempty"` // The Voltage Phase L2 (in V)
	Counterreadingexportt2 float64 `json:"CounterReadingExportT2,omitempty"` // The Meter Counter Reading only export (Tariff 2)
	Deviceenergytype string `json:"DeviceEnergyType,omitempty"` // The Energy Type of this device
	Digitalinput1 bool `json:"DigitalInput1,omitempty"` // The digital input number 1
	Metersubtype string `json:"MeterSubType,omitempty"` // The Sub Type of this Meter.
	Id string `json:"Id,omitempty"` // The ID of the device
	Voltagel3 float64 `json:"VoltageL3,omitempty"` // The Voltage Phase L3 (in V)
	Powerfactorl3 float64 `json:"PowerFactorL3,omitempty"` // The Power Factor (cos phi) Phase L3. Range: 0 - 1
	Voltagel1 float64 `json:"VoltageL1,omitempty"` // The Voltage Phase L1 (in V)
	Counterreadingexport float64 `json:"CounterReadingExport,omitempty"` // The Meter Counter Reading only export
}
