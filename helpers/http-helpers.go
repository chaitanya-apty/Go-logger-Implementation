package helpers

import (
	"go-logging-proposal/errors"
	"net/http"
)

type ResponseMessage struct {
	Success   bool        `json:"success"`
	RequestID string      `json:"reqID"`
	Payload   interface{} `json:"payload"`
}

//GetResponsePayload  - Common Helper
func GetResponsePayload(data interface{}, err *errors.Error, reqID string) (int, ResponseMessage) {
	if err == nil {
		return http.StatusOK, buildResponse(true, reqID, data)
	}
	switch err.ErrorType() { // Future - Common Function
	case errors.Unexpected:
		const code = http.StatusInternalServerError
		return code, buildResponse(false, reqID, err.Error())
	case errors.UnAuthorizedError:
		const code = http.StatusUnauthorized
		return code, buildResponse(false, reqID, "Request Not Authenticated")
	default:
		const code = http.StatusServiceUnavailable
		return code, buildResponse(false, reqID, err.Error())
	}
}

func buildResponse(success bool, reqID string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Success:   success,
		RequestID: reqID,
		Payload:   data,
	}
}
