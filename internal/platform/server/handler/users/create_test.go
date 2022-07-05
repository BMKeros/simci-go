package users

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"simci-go/internal/platform/storage/storagemocks"
	"testing"
)

func TestCreate(t *testing.T) {
	userRepository := new(storagemocks.UserRepository)
	userRepository.On("Save", mock.Anything, mock.Anything).Return(nil)

	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.POST("/users", Create(userRepository))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		newUserRequest := createUserRequest{
			ID:   "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name: "Testing",
		}

		data, err := json.Marshal(newUserRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(data))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given an valid request it returns 201", func(t *testing.T) {
		newUserRequest := createUserRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name:     "Testing",
			Email:    "testing@test.com",
			Password: "testing",
		}

		data, err := json.Marshal(newUserRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(data))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
