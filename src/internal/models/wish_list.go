package models

type WishList struct {
	UserId     string `json:"userId"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	IsDisabled bool   `json:"isDisabled"`
}
