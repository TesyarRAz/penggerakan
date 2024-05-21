package main

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/joho/godotenv"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	dotenv := config.NewDotEnv()

	appId, err := GenerateRandomStringURLSafe(64)
	if err != nil {
		panic(err)
	}

	jwtSecret, err := GenerateRandomStringURLSafe(64)
	if err != nil {
		panic(err)
	}

	jwtRefreshSecret, err := GenerateRandomStringURLSafe(64)
	if err != nil {
		panic(err)
	}

	dotenv["APP_ID"] = appId
	dotenv["JWT_SECRET"] = jwtSecret
	dotenv["JWT_REFRESH_SECRET"] = jwtRefreshSecret

	godotenv.Write(dotenv, ".env")
}
