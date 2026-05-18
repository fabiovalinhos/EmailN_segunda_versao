package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Test Campaign"
	content := "This is a test campaign."
	emails := []string{"email@example.com", "another@example.com"}

	campaign := NewCampaign(name, content, emails)

	if campaign.ID != "1" {
		t.Errorf("Expected ID to be '1', got '%s'", campaign.ID)
	}

	if campaign.Name != name {
		t.Errorf("Expected Name to be '%s', got '%s'", name, campaign.Name)
	}

	if campaign.Content != content {
		t.Errorf("Expected Content to be '%s', got '%s'", content, campaign.Content)
	}

	if len(campaign.Contacts) != len(emails) {
		t.Errorf("Expected %d contacts, got %d", len(emails), len(campaign.Contacts))
	}
}
