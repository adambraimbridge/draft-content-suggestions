package suggestions

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Financial-Times/draft-content-suggestions/mocks"
	"github.com/Financial-Times/draft-content-suggestions/draft"
)

const apiKey = "12345"

func TestUmbrellaAPI_IsHealthySuccess(t *testing.T) {

	testServer := mocks.NewUmbrellaTestServer(true)
	defer testServer.Close()

	umbrellaAPI, err := NewUmbrellaAPI(testServer.URL+"/content/suggest", apiKey, http.DefaultClient)

	assert.NoError(t, err)

	_, err = umbrellaAPI.IsHealthy(context.Background())
	assert.NoError(t, err)
}
func TestUmbrellaAPI_IsHealthyFailure(t *testing.T) {
	testServer := mocks.NewUmbrellaTestServer(false)
	defer testServer.Close()

	umbrellaAPI, err := NewUmbrellaAPI(testServer.URL+"/content/suggest", apiKey, http.DefaultClient)

	assert.NoError(t, err)

	_, err = umbrellaAPI.IsHealthy(context.Background())
	assert.Error(t, err)
}
func TestUmbrellaAPI_FetchSuggestions(t *testing.T) {

	mockDraftContent := newMockDraftContent()

	testServer := mocks.NewUmbrellaTestServer(true)
	defer testServer.Close()

	umbrellaAPI, err := NewUmbrellaAPI(testServer.URL+"/content/suggest", apiKey, http.DefaultClient)
	assert.NoError(t, err)

	suggestions, err := umbrellaAPI.FetchSuggestions(context.Background(), mockDraftContent)

	assert.NoError(t, err)
	assert.True(t, suggestions != nil)
}
func TestUmbrellaAPI_FetchDraftContentFailure(t *testing.T) {

	testServer := mocks.NewUmbrellaTestServer(true)
	testServer.Close()

	contentAPI, err := NewUmbrellaAPI(testServer.URL+"/content/suggest", apiKey, http.DefaultClient)
	assert.NoError(t, err)

	suggestions, err := contentAPI.FetchSuggestions(context.Background(), newMockDraftContent())

	assert.Error(t, err)
	assert.True(t, suggestions == nil)
}

func newMockDraftContent() *draft.Content {
	mockDraftContent := &draft.Content{
		UUID:   "9d5e441e-0b02-11e8-8eb7-42f857ea9f0",
		Body:   "<body><content data-embedded=\"true\" id=\"c0cc4ca2-0b43-11e8-24ad-bec2279df517\" type=\"http://www.ft.com/ontology/content/ImageSet\"></content><p>US stocks see-sawed in early trading on Tuesday, as volatility on global markets intensified, breaking an extended period of calm for investors.xxxx</body>",
		Title:  "Wall Street stocks xxx",
		Byline: "Eric Platt in New York, Michael Hunter and Adam Samson in London",
	}
	return mockDraftContent
}
