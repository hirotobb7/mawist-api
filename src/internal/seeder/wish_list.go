package seeder

import (
	"github.com/hirotobb7/mawist/internal/db/repository/dynamo"
	"github.com/hirotobb7/mawist/internal/db/service"
	"github.com/hirotobb7/mawist/internal/model"
)

var db = dynamo.GetDb()
var wishListService = service.NewWishListService(dynamo.NewWishListRepository(db))

var wishLists = [3]model.WishList{
	{
		UserId:     "test-user-id-1",
		Id:         "test-id-1",
		Name:       "マイリスト",
		CreatedAt:  "2022-05-08T17:00:00Z",
		UpdatedAt:  "2022-05-08T17:00:00Z",
		IsDisabled: false,
	},
	{
		UserId:     "test-user-id-1",
		Id:         "test-id-2",
		Name:       "プレゼントリスト",
		CreatedAt:  "2022-05-10T09:00:00Z",
		UpdatedAt:  "2022-05-10T09:00:00Z",
		IsDisabled: false,
	},
	{
		UserId:     "test-user-id-2",
		Id:         "test-id-10",
		Name:       "マイリスト",
		CreatedAt:  "2022-05-16T09:00:00Z",
		UpdatedAt:  "2022-05-16T09:00:00Z",
		IsDisabled: false,
	},
}

func CreateWishLists() error {
	for _, wishList := range wishLists {
		if err := wishListService.Create(wishList); err != nil {
			return err
		}
	}

	return nil
}

func DeleteWishLists() error {
	for _, wishList := range wishLists {
		if err := wishListService.Delete(wishList); err != nil {
			return err
		}
	}

	return nil
}
