package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"shorty/internal/dto"
	"shorty/internal/service"
	"shorty/internal/service/mocks"
	"testing"
)

func TestHandler_redirect(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockLink, urlParam string)
	testTable := []struct {
		name                 string
		urlParam             string
		mockBehavior         mockBehavior
		expectErr            bool
		expectedStatus       int
		expectedResponseBody string
	}{
		{
			name:     "OK",
			urlParam: "4dfk5ga",
			mockBehavior: func(s *mock_service.MockLink, url string) {
				s.EXPECT().ProcessURLClick(url).Return(dto.Link{
					OriginURL:     "https://google.com",
					ShortURL:      "4dfk5ga",
					RedirectCount: 2,
				}, nil)
			},
			expectErr:      false,
			expectedStatus: 307,
			expectedResponseBody: `<a href="https://google.com">Temporary Redirect</a>.

`,
		},
		{
			name:                 "Missing url parameter",
			urlParam:             "",
			mockBehavior:         nil,
			expectErr:            true,
			expectedStatus:       404,
			expectedResponseBody: "",
		},
		{
			name:     "Service error",
			urlParam: "4dfk5ga",
			mockBehavior: func(s *mock_service.MockLink, url string) {
				s.EXPECT().ProcessURLClick(url).Return(dto.Link{
					OriginURL:     "https://google.com",
					ShortURL:      "4dfk5ga",
					RedirectCount: 2,
				}, errors.New("some error"))
			},
			expectErr:            true,
			expectedStatus:       500,
			expectedResponseBody: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockLink := mock_service.NewMockLink(c)

			if testCase.mockBehavior != nil {
				testCase.mockBehavior(mockLink, testCase.urlParam)
			}

			services := &service.Service{Link: mockLink}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/:url", handler.redirect)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/%s", testCase.urlParam), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatus, w.Code)
			if !testCase.expectErr {
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}
}
