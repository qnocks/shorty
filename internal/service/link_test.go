package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"shorty/internal/dto"
	"shorty/internal/repository/mocks"
	"shorty/internal/service/mocks"
	"testing"
)

func TestLinkService_CreateShortLink(t *testing.T) {
	type mockBehavior = func(s *mock_repository.MockLink, link dto.Link)
	type shortenerServiceMockBehavior = func(s *mock_service.MockShortener, link dto.Link)

	testTable := []struct {
		name                         string
		link                         dto.Link
		mockBehavior                 mockBehavior
		shortenerServiceMockBehavior shortenerServiceMockBehavior
		expectErr                    bool
		errorMessage                 string
	}{
		{
			name: "OK",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 0,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().Save(link).Return(link, nil)
			},
			shortenerServiceMockBehavior: func(s *mock_service.MockShortener, link dto.Link) {
				s.EXPECT().Shorten(link.OriginURL).Return(link.ShortURL)
			},
			expectErr: false,
		},
		{
			name: "Repository error",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 0,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().Save(link).Return(link, errors.New("some error"))
			},
			shortenerServiceMockBehavior: func(s *mock_service.MockShortener, link dto.Link) {
				s.EXPECT().Shorten(link.OriginURL).Return(link.ShortURL)
			},
			expectErr:    true,
			errorMessage: "some error",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)

			mockLink := mock_repository.NewMockLink(c)
			mockShortener := mock_service.NewMockShortener(c)
			testCase.mockBehavior(mockLink, testCase.link)
			testCase.shortenerServiceMockBehavior(mockShortener, testCase.link)

			link := NewLinkService(mockLink, mockShortener)
			actual, err := link.CreateShortLink(testCase.link.OriginURL)

			if testCase.expectErr {
				assert.Equal(t, err.Error(), testCase.errorMessage)
			} else {
				assert.Equal(t, err, nil)
				assert.Equal(t, actual.ShortURL, testCase.link.ShortURL)
			}
		})
	}
}

func TestLinkService_GetByShortURL(t *testing.T) {
	type mockBehavior = func(s *mock_repository.MockLink, link dto.Link)

	testTable := []struct {
		name         string
		link         dto.Link
		mockBehavior mockBehavior
		expectErr    bool
		errorMessage string
	}{
		{
			name: "OK",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 2,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().FindByShortUrl(link.ShortURL).Return(link, nil)
			},
			expectErr: false,
		},
		{
			name: "Repository error",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 2,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().FindByShortUrl(link.ShortURL).Return(link, errors.New("some error"))
			},
			expectErr:    true,
			errorMessage: "some error",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockLink := mock_repository.NewMockLink(c)
			testCase.mockBehavior(mockLink, testCase.link)

			link := NewLinkService(mockLink, nil)
			actual, err := link.GetByShortURL(testCase.link.ShortURL)

			if testCase.expectErr {
				assert.Equal(t, err.Error(), testCase.errorMessage)
			} else {
				assert.Equal(t, err, nil)
				assert.Equal(t, actual, testCase.link)
			}
		})
	}
}

func TestLinkService_ProcessURLClick(t *testing.T) {
	type mockBehavior = func(s *mock_repository.MockLink, link dto.Link)

	testTable := []struct {
		name         string
		link         dto.Link
		mockBehavior mockBehavior
		expectErr    bool
		errorMessage string
	}{
		{
			name: "OK",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 2,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().FindByShortUrl(link.ShortURL).Return(link, nil)
				s.EXPECT().Save(dto.Link{
					OriginURL:     link.OriginURL,
					ShortURL:      link.ShortURL,
					RedirectCount: link.RedirectCount + 1,
				}).Return(link, nil)
			},
			expectErr: false,
		},
		{
			name: "Repository error (cannot find link)",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 2,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().FindByShortUrl(link.ShortURL).Return(link, errors.New("some error"))
			},
			expectErr:    true,
			errorMessage: "some error",
		},
		{
			name: "Repository error (cannot update link info)",
			link: dto.Link{
				OriginURL:     "https://google.com",
				ShortURL:      "4abcde",
				RedirectCount: 2,
			},
			mockBehavior: func(s *mock_repository.MockLink, link dto.Link) {
				s.EXPECT().FindByShortUrl(link.ShortURL).Return(link, nil)
				s.EXPECT().Save(dto.Link{
					OriginURL:     link.OriginURL,
					ShortURL:      link.ShortURL,
					RedirectCount: link.RedirectCount + 1,
				}).Return(link, errors.New("some error"))
			},
			expectErr:    true,
			errorMessage: "some error",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockLink := mock_repository.NewMockLink(c)
			testCase.mockBehavior(mockLink, testCase.link)

			link := NewLinkService(mockLink, nil)
			actual, err := link.ProcessURLClick(testCase.link.ShortURL)

			if testCase.expectErr {
				assert.Equal(t, err.Error(), testCase.errorMessage)
			} else {
				assert.Equal(t, err, nil)
				assert.Equal(t, actual.OriginURL, testCase.link.OriginURL)
			}
		})
	}
}
