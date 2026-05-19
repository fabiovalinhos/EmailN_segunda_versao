package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Test Campaign"
	content = "This is a test campaign."
	emails  = []string{"email@example.com", "another@example.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, emails)

	assert.Equal(name, campaign.Name, "Expected Name to match")
	assert.Equal(content, campaign.Content, "Expected Content to match")
	assert.Len(campaign.Contacts, len(emails), "Expected number of contacts to match")

	//// Sem usar o pacote testing/assert
	////
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

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID, "Expected ID to be not nil")
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}
