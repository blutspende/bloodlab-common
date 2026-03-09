package instrumentenum

type Type string

const (
	TypeAnalyzer Type = "ANALYZER"
	TypeSorter   Type = "SORTER"
)

type ConnectionMode string

const (
	ConnectionModeTCPClient  ConnectionMode = "TCP_CLIENT_ONLY"
	ConnectionModeTCPServer  ConnectionMode = "TCP_SERVER_ONLY"
	ConnectionModeFileServer ConnectionMode = "FILE_SERVER"
	ConnectionModeHTTP       ConnectionMode = "HTTP"
	ConnectionModeTCPMixed   ConnectionMode = "TCP_MIXED"
)

type ConnectionStatus string

const (
	ConnectionStatusOffline ConnectionStatus = "OFFLINE"
	ConnectionStatusReady   ConnectionStatus = "READY"
	ConnectionStatusOnline  ConnectionStatus = "ONLINE"
)

type Ability string

const (
	AbilityAllowResending         Ability = "ALLOW_RESENDING"
	AbilityCanAcceptResults       Ability = "CAN_ACCEPT_RESULTS"
	AbilityCanReplyToQuery        Ability = "CAN_REPLY_TO_QUERY"
	AbilityCanCaptureDiagnostics  Ability = "CAN_CAPTURE_DIAGNOSTICS"
	AbilityCanUseFtpConnection    Ability = "CAN_USE_FTP_CONNECTION"
	AbilityCanUseSftpConnection   Ability = "CAN_USE_SFTP_CONNECTION"
	AbilityCanUseWebdavConnection Ability = "CAN_USE_WEBDAV_CONNECTION"
)

type ProtocolSettingType string

const (
	ProtocolSettingTypeString   ProtocolSettingType = "string"
	ProtocolSettingTypeInt      ProtocolSettingType = "int"
	ProtocolSettingTypeBool     ProtocolSettingType = "bool"
	ProtocolSettingTypePassword ProtocolSettingType = "password"
)

type FileServerType string

const (
	FileServerTypeFTP    FileServerType = "FTP"
	FileServerTypeSFTP   FileServerType = "SFTP"
	FileServerTypeWEBDAV FileServerType = "WEBDAV"
)

type ConfigurationMessageType string

const (
	ConfigurationMessageTypeCreate ConfigurationMessageType = "CREATE"
	ConfigurationMessageTypeUpdate ConfigurationMessageType = "UPDATE"
	ConfigurationMessageTypeDelete ConfigurationMessageType = "DELETE"
)

type ReprocessMessageType string

const (
	ReprocessMessageTypeRetransmitResult      ReprocessMessageType = "RETRANSMIT_RESULT"
	ReprocessMessageTypeReprocessBySampleCode ReprocessMessageType = "REPROCESS_BY_SAMPLE_CODE"
	ReprocessMessageTypeReprocessByDEAIds     ReprocessMessageType = "REPROCESS_BY_DEA_IDS"
)

type MessageStatus string

const MessageStatusStored MessageStatus = "STORED"
const MessageStatusProcessed MessageStatus = "PROCESSED"
const MessageStatusError MessageStatus = "ERROR"
const MessageStatusSent MessageStatus = "SENT"

type MessageType string

const MessageTypeQuery MessageType = "QUERY"
const MessageTypeOrder MessageType = "ORDER"
const MessageTypeResult MessageType = "RESULT"
const MessageTypeAcknowledgement MessageType = "ACKNOWLEDGEMENT"
const MessageTypeCancellation MessageType = "CANCELLATION"
const MessageTypeReorder MessageType = "REORDER"
const MessageTypeDiagnostics MessageType = "DIAGNOSTICS"
const MessageTypeUnidentified MessageType = "UNIDENTIFIED"

type ResultMode string

const (
	// Simulated results will not be transmitted to Cerberus and stay within the driver
	ResultModeSimulation ResultMode = "SIMULATION"
	// Qualification Results are transmitted to cerberus but not returned to any EIA interface
	ResultModeQualification ResultMode = "QUALIFICATION"
	// Production allows the results to be returned via EIA
	ResultModeProduction ResultMode = "PRODUCTION"
)

type ResultType string

const (
	ResultTypeInt            ResultType = "int"
	ResultTypeDecimal        ResultType = "decimal"
	ResultTypeBoundedDecimal ResultType = "boundedDecimal"
	ResultTypeString         ResultType = "string"
	ResultTypePein           ResultType = "pein"
	ResultTypeReact          ResultType = "react"
	ResultTypeInValid        ResultType = "invalid"
	ResultTypeEnum           ResultType = "enum"
)

type ReagentType string

const (
	ReagentTypeStandard ReagentType = "reagent"
	ReagentTypeDiluent  ReagentType = "diluent"
)
