package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/festivio/festivio-backend/domain"
	mock_service "github.com/festivio/festivio-backend/internal/service/mocks"
	"github.com/festivio/festivio-backend/pkg/converter"
)

func TestSignupUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockServiceInterface(ctrl)
	h := handler{srv: mockService}

	router := gin.New()
	router.POST("/sign-up", h.SignUpUser())

	t.Run("Successful Signup", func(t *testing.T) {
		input := &domain.SignUpInput{
			Email:           "test@example.com",
			Name:            "Alex",
			Password:        "password123",
			PasswordConfirm: "password123",
			Role:            "Manager",
		}
		expectedStatusCode := http.StatusCreated

		mockService.EXPECT().SignUpUser(input).Return(nil)

		buf, err := converter.AnyToBytesBuffer(input)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sign-up", strings.NewReader(buf.String()))
		req.SetBasicAuth(input.Email, input.Password)
		router.ServeHTTP(w, req)

		assert.Equal(t, expectedStatusCode, w.Code)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		input := &domain.SignUpInput{}
		expectedStatusCode := http.StatusBadRequest
		expectedError := &domain.ErrorStruct{
			Error: struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusBadRequest,
				Message: "Invalid request body.",
			},
		}

		buf, err := converter.AnyToBytesBuffer(input)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sign-up", strings.NewReader(buf.String()))
		router.ServeHTTP(w, req)

		assert.Equal(t, expectedStatusCode, w.Code)

		var errorStruct domain.ErrorStruct
		err = json.Unmarshal(w.Body.Bytes(), &errorStruct)
		assert.Nil(t, err)
		assert.Equal(t, expectedError, &errorStruct)
	})

	// TODO: Add more test cases for other scenarios
}
