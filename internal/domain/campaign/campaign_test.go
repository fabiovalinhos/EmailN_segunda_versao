package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	name := "Test Campaign"
	content := "This is a test campaign."
	emails := []string{"email@example.com", "another@example.com"}
	assert := assert.New(t)

	campaign := NewCampaign(name, content, emails)

	assert.Equal("1", campaign.ID, "Expected ID to be '1'")
	assert.Equal(name, campaign.Name, "Expected Name to match")
	assert.Equal(content, campaign.Content, "Expected Content to match")
	assert.Len(campaign.Contacts, len(emails), "Expected number of contacts to match")

	// if campaign.ID != "1" {
	// 	t.Errorf("Expected ID to be '1', got '%s'", campaign.ID)
	// }

	// if campaign.Name != name {
	// 	t.Errorf("Expected Name to be '%s', got '%s'", name, campaign.Name)
	// }

	// if campaign.Content != content {
	// 	t.Errorf("Expected Content to be '%s', got '%s'", content, campaign.Content)
	// }

	// if len(campaign.Contacts) != len(emails) {
	// 	t.Errorf("Expected %d contacts, got %d", len(emails), len(campaign.Contacts))
	// }
}
