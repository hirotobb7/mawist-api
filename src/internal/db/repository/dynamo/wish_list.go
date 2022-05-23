package dynamo

import (
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"

	"github.com/hirotobb7/mawist/internal/db/repository"
	"github.com/hirotobb7/mawist/internal/model"
)

const tableName = "wish_lists"

type wishListRepository struct {
	table dynamo.Table
}

func NewWishListRepository(db *dynamo.DB) repository.WishListRepository {
	return &wishListRepository{table: db.Table(tableName)}
}

func (wR *wishListRepository) FindByUserId(userId string) ([]model.WishList, error) {

	var wishLists []model.WishList
	if err := wR.table.Get("user_id", userId).All(&wishLists); err != nil {
		return wishLists, errors.WithStack(err)
	}

	return wishLists, nil
}

func (wR *wishListRepository) Create(wishList model.WishList) error {
	if err := wR.table.Put(wishList).If("attribute_not_exists(id)").Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (wR *wishListRepository) Delete(wishList model.WishList) error {
	if err := wR.table.Delete("user_id", wishList.UserId).Range("id", wishList.Id).Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
