package messagetype

type MessageType string

const Query MessageType = "QUERY"
const Order MessageType = "ORDER"
const Result MessageType = "RESULT"
const Acknowledgement MessageType = "ACKNOWLEDGEMENT"
const Cancellation MessageType = "CANCELLATION"
const Reorder MessageType = "REORDER"
const Diagnostics MessageType = "DIAGNOSTICS"
const Unidentified MessageType = "UNIDENTIFIED"
