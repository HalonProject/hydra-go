package hydra

import "fmt"

const (
	HYDRA_IPFS_READ_RECORD_ERROR    = iota + 1
	HYDRA_RECORD_UNMARSHAL_ERROR    = iota
	HYDRA_RESPONSE_TIMEOUT_ERROR    = iota
	HYDRA_ZERO_BYTES_RECIEVED_ERROR = iota
)

// Error represent a general error returned from the reading or writing of a message
// to IPFS pubsub.
type Error struct {
	code   int
	reason string
}

// newError returns a new error with an error code provided and a blank reason.
func newError(code int, reason string) Error {
	return Error{
		code:   code,
		reason: reason,
	}
}

// String returns the reason for the error
func (e Error) String() string {
	return e.reason
}

// Error return a formatted error string so that we can pass as type `error`
func (e Error) Error() string {
	return fmt.Sprintf("error code %d: %s", e.code, e.reason)
}

// Code returns the int value of the error
func (e Error) Code() int {
	return e.code
}
