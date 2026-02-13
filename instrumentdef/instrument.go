package instrumentdef

type InstrumentType string

const (
	Analyzer InstrumentType = "ANALYZER"
	Sorter   InstrumentType = "SORTER"
)

type InstrumentStatus string

const (
	InstrumentOffline InstrumentStatus = "OFFLINE"
	InstrumentReady   InstrumentStatus = "READY"
	InstrumentOnline  InstrumentStatus = "ONLINE"
)

type ConnectionMode string

const (
	TCPClientMode ConnectionMode = "TCP_CLIENT_ONLY"
	TCPServerMode ConnectionMode = "TCP_SERVER_ONLY"
	FileServer    ConnectionMode = "FILE_SERVER"
	HTTP          ConnectionMode = "HTTP"
	TCPMixed      ConnectionMode = "TCP_MIXED"
)

type Ability string

const (
	CanAcceptResultsAbility       Ability = "CAN_ACCEPT_RESULTS"
	CanReplyToQueryAbility        Ability = "CAN_REPLY_TO_QUERY"
	CanCaptureDiagnosticsAbility  Ability = "CAN_CAPTURE_DIAGNOSTICS"
	CanUseFtpConnectionAbility    Ability = "CAN_USE_FTP_CONNECTION"
	CanUseSftpConnectionAbility   Ability = "CAN_USE_SFTP_CONNECTION"
	CanUseWebdavConnectionAbility Ability = "CAN_USE_WEBDAV_CONNECTION"
)

type ProtocolSettingType string

const (
	String   ProtocolSettingType = "string"
	Int      ProtocolSettingType = "int"
	Bool     ProtocolSettingType = "bool"
	Password ProtocolSettingType = "password"
)

type MessageType string

const (
	MessageTypeCreate MessageType = "CREATE"
	MessageTypeUpdate MessageType = "UPDATE"
	MessageTypeDelete MessageType = "DELETE"
)

type ReprocessMessageType string

const (
	MessageTypeRetransmitResult      ReprocessMessageType = "RETRANSMIT_RESULT"
	MessageTypeReprocessBySampleCode ReprocessMessageType = "REPROCESS_BY_SAMPLE_CODE"
	MessageTypeReprocessByDEAIds     ReprocessMessageType = "REPROCESS_BY_DEA_IDS"
)

type FileServerType string

const (
	FTP    FileServerType = "FTP"
	SFTP   FileServerType = "SFTP"
	WEBDAV FileServerType = "WEBDAV"
)
