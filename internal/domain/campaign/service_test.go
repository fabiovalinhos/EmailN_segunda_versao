package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaignDto{
		Name:    "Test Campaign",
		Content: "This is a test campaign",
		Emails:  []string{"teste1@example.com"},
	}

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotNil(id)
}
func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaignDto{
		Name:    "Test Campaign",
		Content: "This is a test campaign",
		Emails:  []string{"teste1@example.com"},
	}
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content {
			return false
		}

		if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
