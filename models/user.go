package models

type User struct {
	Id                  int    `json:"id"`
	UserId              string `json:"user_id"`
	Otp                 string `json:"otp"`
	RequestId           string `json:"request_id"`
	OtpGeneratedAtEpoch int64  `json:"otp_generated_at_epoch"`
	IsValidated         bool   `json:"is_validated"`
}

type OtpGenerationRequest struct {
	UserId    string `json:"user_id"`
	RequestId string `json:"request_id"`
}

type OtpGenerationResponse struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}

type OtpValidationRequest struct {
	UserId    string `json:"user_id"`
	RequestId string `json:"request_id"`
	Otp       string `json:"otp"`
}

type OtpValidationErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type OtpValidationSuccessResponse struct {
	UserId  string `json:"user_id"`
	Message string `json:"message"`
}
