package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortenerService_Shorten(t *testing.T) {
	minShortLength := 10
	testTable := []struct {
		given             string
		minExpectedLength int
	}{
		{given: "https://google.com", minExpectedLength: minShortLength},
		{given: "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley", minExpectedLength: minShortLength},
		{given: "https://www.youtube.com/watch?v=mrPHeRvmjqM&ab_channel=TheVoice%3Alaplusbellevoix", minExpectedLength: minShortLength},
		{given: "https://www.youtube.com/watch?v=IFr6f33oWd4&list=PLeCgg1XaybKuDbgfWhS6-fwRUPSnHT0ed&index=6&ab_channel=%D0%A1%D1%82%D0%B0%D1%81%D0%9A%D0%B8%D1%81%D0%B5%D0%BB%D0%B5%D0%B2", minExpectedLength: minShortLength},
	}

	for _, testCase := range testTable {
		service := NewShortenerService()
		actual := service.Shorten(testCase.given)

		assert.GreaterOrEqual(t, len(actual), testCase.minExpectedLength)
	}
}
