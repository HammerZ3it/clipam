package config

type AuthAPIResponse struct {
	// The HTTP result code.
	Token string `json:"token"`
  TTL string `json:"expires"`
}
