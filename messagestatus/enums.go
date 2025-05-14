package messagestatus

type MessageStatus string

const Stored MessageStatus = "STORED"
const Processed MessageStatus = "PROCESSED"
const Error MessageStatus = "ERROR"
const Sent MessageStatus = "SENT"
