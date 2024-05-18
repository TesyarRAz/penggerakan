package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	token := GetAdminToken(t)

	requestBody := model.CreateCourseRequest{
		Name:  "Course 1",
		Image: "Image 1",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/courses", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CourseResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.Image, responseBody.Data.Image)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	// assert.NotNil(t, responseBody.Data.UpdatedAt)

	created, err := db.MustExec("SELECT COUNT(*) FROM courses WHERE id = $1", responseBody.Data.ID).RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), created)
}

func TestList(t *testing.T) {
	token := GetAdminToken(t)

	request := httptest.NewRequest(http.MethodGet, "/courses", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.PageResponse[model.CourseResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
