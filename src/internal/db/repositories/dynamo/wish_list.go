package dynamo

import (
	"github.com/guregu/dynamo"

	dtos "github.com/hirotobb7/mawist/internal/db/dtos/dynamo"
	"github.com/hirotobb7/mawist/internal/db/repositories"
	"github.com/hirotobb7/mawist/internal/models"
	"github.com/pkg/errors"
)

const tableName string = "wish_lists"

type wishListRepository struct {
	table dynamo.Table
}

func NewWishListRepository(db *dynamo.DB) repositories.WishListRepository {
	return &wishListRepository{table: db.Table(tableName)}
}

func (wR *wishListRepository) FindByUserId(userId string) ([]models.WishList, error) {
	_wishLists := make([]dtos.WishList, 0)

	if err := wR.table.Get("user_id", userId).All(&_wishLists); err != nil {
		return nil, errors.WithStack(err)
	}

	wishLists := make([]models.WishList, len(_wishLists), len(_wishLists))

	for i, _wishList := range _wishLists {
		wishLists[i] = *dtos.ConvertWishListModel(&_wishList)
	}

	return wishLists, nil
}

func (wR *wishListRepository) Create(wishList *models.WishList) error {
	if err := wR.table.Put(dtos.ConvertWishListDto(wishList)).If("attribute_not_exists(id)").Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (wR *wishListRepository) Delete(wishList *models.WishList) error {
	if err := wR.table.Delete("user_id", wishList.UserId).Range("id", wishList.Id).Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
