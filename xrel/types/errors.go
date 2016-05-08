package types

const ERROR_TYPE_CLIENT = "client"
const ERROR_TYPE_OAUTH = "oauth2"
const ERROR_TYPE_API = "api"

/*
Error is a struct which implements the normal `error` type, all methods may return this struct.
Additional to the `Error()` function this struct contains some variables.

Errors with type `ERROR_TYPE_API` are xREL.to errors (https://www.xrel.to/wiki/6435/api-errors.html), for all other
types and error codes see the `NewError` function.

To use this in your code try to cast it:
err := xrel.SomeMethod()
if eErr, ok := err.(*types.Error); ok {
	// Is of type types.Error, you can use the variables
} else {
	// Is normal error
}
*/
type Error struct {
	Type        string `json:"error_type"`
	Code        string `json:"error"`
	Extra       string
	Description string `json:"error_description"`
}

/*
Error returns the error description.
*/
func (e *Error) Error() string {
	return e.Description
}

/*
NewError creates a new error. If you don't specify a description it will be set by the error code.
*/
func NewError(errorType, errorCode, errorExtra, errorDesc string) *Error {
	err := &Error{}
	err.Type = errorType
	err.Code = errorCode
	err.Extra = errorExtra
	if errorDesc == "" {
		switch errorCode {
		// Client errors
		case "parsing_failed":
			errorDesc = "Failed to parse xREL response. Please report to: https://github.com/hashworks/go-xREL-API/issues"

		case "function_not_found":
			errorDesc = "API function not found. Please report to: https://github.com/hashworks/go-xREL-API/issues"

		case "argument_missing":
			if err.Extra != "" && len(err.Extra) < 30 {
				errorDesc = "Required parameter " + err.Extra + " missing."
			} else {
				errorDesc = "Required parameter missing."
			}
		case "invalid_argument":
			if err.Extra != "" && len(err.Extra) < 30 {
				errorDesc = "The argument for " + err.Extra + " is invalid."
			} else {
				errorDesc = "An argument is invalid."
			}
		case "file_not_found":
			if err.Extra != "" && len(err.Extra) < 60 {
				errorDesc = "File '" + err.Extra + " not found."
			} else {
				errorDesc = "File not found."
			}
		case "not_authenticated":
			errorDesc = "You're not authenticated with xREL."

		// API errors. Also see https://www.xrel.to/wiki/6435/api-errors.html
		case "ratelimit_reached":
			errorDesc = "Ratelimit reached."

		case "internal_error":
			errorDesc = "Internal error on the xREL server. If this keeps occuring, please report to a developer: https://www.xrel.to/wiki/213/Team.html"

		// Other
		default:
			errorDesc = "Unknown error code '" + err.Code + "'. Please report to: https://github.com/hashworks/go-xREL-API/issues"
		}
	}
	err.Description = errorDesc
	return err
}
