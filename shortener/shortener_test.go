package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink := "https://github.com/szy0syz"
	userId := "jwQqr3aIf9dM5x4c"
	shortLink := GenerateShortURL(initialLink, userId)

	assert.Equal(t, shortLink, "J8XCb24w")
}
