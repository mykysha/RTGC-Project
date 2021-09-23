package v1

type Response struct {
	ID      string `json:"id"`
	Error   bool   `json:"err"`
	ErrText string `json:"errtext,omitempty"`
}
