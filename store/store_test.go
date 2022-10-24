package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndretrieval(t *testing.T) {
	initialLink := "https://github.com/szy0syz"
	shortURL := "J8XCb24w"

	// Persist data mapping
	SaveURLInRedis(shortURL, initialLink)

	// Retrieve initial URL
	retrievedURL := RetrieveInitialURLFromRedis(shortURL)
	assert.Equal(t, initialLink, retrievedURL)
}
