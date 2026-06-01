package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
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

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Campaign",
		Content: "This is a test campaign",
		Emails:  []string{"teste1@example.com"},
	}

	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {

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

	service.Repository = repositoryMock
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {

	assert := assert.New(t)
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on databasße"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.True(errors.Is(err, internalerrors.ErrInternal))

}
