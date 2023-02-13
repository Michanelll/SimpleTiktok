package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	Msg        string `json:"status_msg,omitempty"`
}
