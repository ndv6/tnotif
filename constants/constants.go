package constants

// Response status
const (
	Success = "SUCCESS"
	Failed  = "FAILED"
)

// Error messages
const (
	RequestHasInvalidFields = "Mandatory field(s) is missing from request, or is invalid. Please ensure all fields are properly filled."
	CannotReadRequest       = "Cannot read request body"
	CannotParseRequest      = "Unable to parse request"
	UnauthorizedRequest     = "Failed to recognize deposit source. Please use authorized banks to deposit."
	CannotEncodeResponse    = "Failed to encode response."
	InsertFailed            = "Failed to insert to database."
	InitLogFailed           = "Failed to insert log of this transaction."
	CannotParseURLParams    = "Failed to parse URL Params"
	FailedConnectDatabase   = "Failed to connect to database"
	FailedParseTemplate     = "Failed to parse email template"
)

// Headers
const (
	ContentType = "Content-Type"
)

// Header types
const (
	JSON = "application/json"
)

// Response Messages
const (
	Welcome         = "Welcome to Tnotif"
	SendMailSuccess = "Send Mail success"
	SendMailFailed  = "Send Mail failed"
)
