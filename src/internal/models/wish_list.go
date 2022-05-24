package models

type WishList struct {
	UserId     string `dynamo:"user_id" json:"userId"` // PK
	Id         string `dynamo:"id" json:"id"`          // SK
	Name       string `dynamo:"name" json:"name"`
	CreatedAt  string `dynamo:"created_at" json:"createdAt"`
	UpdatedAt  string `dynamo:"updated_at" json:"updatedAt"` // LSI
	IsDisabled bool   `dynamo:"is_disabled" json:"isDisabled"`
}
