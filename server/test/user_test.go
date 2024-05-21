package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	requestBody, response, responseBody := GetAdmin(t)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Email, responseBody.Email)
	assert.NotNil(t, responseBody.AccessToken)
	assert.NotNil(t, responseBody.CreatedAt)
}

func TestMe(t *testing.T) {
	token := GetAdminToken(t)

	request := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := fiber.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	var responseBody user_model.LoginUserResponse
	err = json.Unmarshal(bytes, &responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, responseBody.ID)
	assert.NotNil(t, responseBody.Email)
	assert.NotNil(t, responseBody.CreatedAt)
}
