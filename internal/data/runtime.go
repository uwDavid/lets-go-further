package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// implement MarshalJSON method on Runtime type => satisfy json.Marshaler interface
// return JSON-encoded value for the movie runtime, a string "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// wrap jsonValue in double quotes
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
