package seeders

import (
	"github.com/hiroto7/mawist/internal/dynamo"
)

var wishLists = [3]dynamo.WishList{
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
		if err := wishList.Create(); err != nil {
			return err
		}
	}

	return nil
}

func DeleteWishLists() error {
	for _, wishList := range wishLists {
		if err := wishList.Delete(); err != nil {
			return err
		}
	}

	return nil
}
