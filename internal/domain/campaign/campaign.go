package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name, content string, emails []string) Campaign {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}

	return Campaign{
		ID:        "2",
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
}
