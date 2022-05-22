package dynamo

import (
	"github.com/pkg/errors"
)

var table = db.Table("wish_lists")

type WishList struct {
	UserId     string `dynamo:"user_id" json:"userId"` // PK
	Id         string `dynamo:"id" json:"id"`          // SK
	Name       string `dynamo:"name" json:"name"`
	CreatedAt  string `dynamo:"created_at" json:"createdAt"`
	UpdatedAt  string `dynamo:"updated_at" json:"updatedAt"` // LSI
	IsDisabled bool   `dynamo:"is_disabled" json:"isDisabled"`
}

func FindWishListsByUserId(userId string) ([]WishList, error) {
	var wishLists []WishList
	if err := table.Get("user_id", userId).All(&wishLists); err != nil {
		return wishLists, errors.WithStack(err)
	}

	return wishLists, nil
}

func (wishList WishList) Create() error {
	if err := table.Put(wishList).If("attribute_not_exists(id)").Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (wishList WishList) Delete() error {
	if err := table.Delete("user_id", wishList.UserId).Range("id", wishList.Id).Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
