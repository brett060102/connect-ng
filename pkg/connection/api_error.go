package connection

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SUSE/connect-ng/internal/util"
)

// ApiError contains all the information for any given API error response. Don't
// build it directly, but use `ErrorFromResponse` instead.
type ApiError struct {
	Code             int
	Message          string `json:"error"`
	LocalizedMessage string `json:"localized_error"`
}

func (ae *ApiError) Error() string { 
        out:=fmt.Sprintf("api_error.go\n\n")
        util.LogStuff(out)

	if ae.LocalizedMessage != "" {
		return fmt.Sprintf("Error: Registration server returned '%v' (%d)", ae.LocalizedMessage, ae.Code)
	}
        out=fmt.Sprintf("api_error.go:message %+v\n\n",ae)
        util.LogStuff(out)
	return fmt.Sprintf("Error: Registration server returned '%v' (%d)", ae.Message, ae.Code)
}

// Returns a new ApiError from the given response if it contained an API error
// response. Otherwise it just returns nil.
func ErrorFromResponse(resp *http.Response) *ApiError {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	ae := &ApiError{Code: resp.StatusCode}
	if err := json.NewDecoder(resp.Body).Decode(ae); err != nil {
		// In some servers the response is actually not a JSON message, but
		// rather some NGinx default page. In that case, just set the HTML
		// status string as the message.
		ae.Message = resp.Status
		ae.LocalizedMessage = resp.Status
	}
	return ae
}
