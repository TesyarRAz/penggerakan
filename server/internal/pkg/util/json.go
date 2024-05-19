package util

import (
	"bytes"
	"encoding/json"
)

func TrimJsonRawMessage(raw json.RawMessage) (json.RawMessage, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, raw); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
