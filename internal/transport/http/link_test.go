package http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"shorty/internal/dto"
	"shorty/internal/service"
	"shorty/internal/service/mocks"
	"testing"
)

func TestHandler_createShortURL(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockLink, link dto.Link)

	testTable := []struct {
		name                 string
		inputBody            string
		inputLink            dto.Link
		mockBehavior         mockBehavior
		expectErr            bool
		expectedStatus       int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"originURL": "https://google.com"}`,
			inputLink: dto.Link{OriginURL: "https://google.com"},
			mockBehavior: func(s *mock_service.MockLink, link dto.Link) {
				s.EXPECT().CreateShortLink(link.OriginURL).Return(dto.Link{
					OriginURL: link.OriginURL,
					ShortURL:  "4dfk5ga",
				}, nil)
			},
			expectErr:            false,
			expectedStatus:       201,
			expectedResponseBody: `{"originUrl":"https://google.com","shortUrl":"4dfk5ga","redirectCount":0}`,
		},
		{
			name:      "Service error",
			inputBody: `{"originURL": "https:google.com"}`,
			inputLink: dto.Link{OriginURL: "https:google.com"},
			mockBehavior: func(s *mock_service.MockLink, link dto.Link) {
				s.EXPECT().CreateShortLink(link.OriginURL).Return(dto.Link{
					OriginURL: link.OriginURL,
					ShortURL:  "4dfk5ga",
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
			testCase.mockBehavior(mockLink, testCase.inputLink)

			services := &service.Service{Link: mockLink}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/api/links", handler.createShortURL)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/links", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatus, w.Code)
			if !testCase.expectErr {
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}
}

func TestHandler_getLinkInfo(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockLink, url string)
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
				s.EXPECT().GetByShortURL(url).Return(dto.Link{
					OriginURL:     "https://google.com",
					ShortURL:      "4dfk5ga",
					RedirectCount: 2,
				}, nil)
			},
			expectErr:            false,
			expectedStatus:       200,
			expectedResponseBody: `{"originUrl":"https://google.com","shortUrl":"4dfk5ga","redirectCount":2}`,
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
				s.EXPECT().GetByShortURL(url).Return(dto.Link{
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
			r.GET("/api/links/:url", handler.getLinkInfo)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/links/%s", testCase.urlParam), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatus, w.Code)
			if !testCase.expectErr {
				assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}
}
