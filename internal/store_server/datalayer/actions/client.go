package actions

import "gorm.io/gorm"

type Client struct {
	AuthEmailPasswordAction AuthEmailPasswordActionInterface
	ProfileAction           ProfileActionInterface
}

func New(
	db *gorm.DB,
) *Client {
	return &Client{
		AuthEmailPasswordAction: NewAuthEmailPasswordAction(db),
		ProfileAction:           NewProfileAction(db),
	}
}
