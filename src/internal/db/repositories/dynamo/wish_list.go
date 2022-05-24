package dynamo

import (
	"github.com/guregu/dynamo"
	"github.com/hirotobb7/mawist/internal/db/repositories"
	"github.com/hirotobb7/mawist/internal/models"
	"github.com/pkg/errors"
)

const tableName = "wish_lists"

type wishListRepository struct {
	table dynamo.Table
}

func NewWishListRepository(db *dynamo.DB) repositories.WishListRepository {
	return &wishListRepository{table: db.Table(tableName)}
}

func (wR *wishListRepository) FindByUserId(userId string) ([]models.WishList, error) {

	var wishLists []models.WishList
	if err := wR.table.Get("user_id", userId).All(&wishLists); err != nil {
		return wishLists, errors.WithStack(err)
	}

	return wishLists, nil
}

func (wR *wishListRepository) Create(wishList models.WishList) error {
	if err := wR.table.Put(wishList).If("attribute_not_exists(id)").Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (wR *wishListRepository) Delete(wishList models.WishList) error {
	if err := wR.table.Delete("user_id", wishList.UserId).Range("id", wishList.Id).Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
