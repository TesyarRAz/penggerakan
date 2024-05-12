package util

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func StringOrDefault(str string, def string) string {
	if str == "" {
		return def
	}

	return str
}
