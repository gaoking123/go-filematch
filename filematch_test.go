package filematch

import "testing"
import "github.com/stretchr/testify/assert"

func TestMatchPath(t *testing.T) {
	testCases := []struct {
		Path     string
		Patterns []string
		Match    bool
		Error    bool
	}{
		{".git/objects/nope.txt", []string{".git/"}, true, false},
	}

	for _, tc := range testCases {
		match, err := MatchPath(tc.Path, tc.Patterns)
		if !tc.Error {
			assert.Nil(t, err, "Should be nil")
			assert.Equal(t, match, tc.Match, "Should be equal")
		}

	}
}
