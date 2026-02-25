package instrument

type Type string

const (
	TypeAnalyzer Type = "ANALYZER"
	TypeSorter   Type = "SORTER"
)

type ConnectionStatus string

const (
	ConnectionStatusOffline ConnectionStatus = "OFFLINE"
	ConnectionStatusReady   ConnectionStatus = "READY"
	ConnectionStatusOnline  ConnectionStatus = "ONLINE"
)

type ConnectionMode string

const (
	ConnectionModeTCPClient  ConnectionMode = "TCP_CLIENT_ONLY"
	ConnectionModeTCPServer  ConnectionMode = "TCP_SERVER_ONLY"
	ConnectionModeFileServer ConnectionMode = "FILE_SERVER"
	ConnectionModeHTTP       ConnectionMode = "HTTP"
	ConnectionModeTCPMixed   ConnectionMode = "TCP_MIXED"
)

type Ability string

const (
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
