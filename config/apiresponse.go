package config

import (
	"encoding/json"
)

type APIResponse struct {
	Code int
  Success bool
	Data json.RawMessage
	Message string
}
