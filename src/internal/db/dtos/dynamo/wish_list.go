package dynamo

import "github.com/hirotobb7/mawist/internal/models"

type WishList struct {
	UserId     string `dynamo:"user_id"` // PK
	Id         string `dynamo:"id"`      // SK
	Name       string `dynamo:"name"`
	CreatedAt  string `dynamo:"created_at"`
	UpdatedAt  string `dynamo:"updated_at"` // LSI
	IsDisabled bool   `dynamo:"is_disabled"`
}

func ConvertWishListDto(wishList *models.WishList) *WishList {
	return &WishList{
		UserId:     wishList.UserId,
		Id:         wishList.Id,
		Name:       wishList.Name,
		CreatedAt:  wishList.CreatedAt,
		UpdatedAt:  wishList.UpdatedAt,
		IsDisabled: wishList.IsDisabled,
	}

}

func ConvertWishListModel(wishList *WishList) *models.WishList {
	return &models.WishList{
		UserId:     wishList.UserId,
		Id:         wishList.Id,
		Name:       wishList.Name,
		CreatedAt:  wishList.CreatedAt,
		UpdatedAt:  wishList.UpdatedAt,
		IsDisabled: wishList.IsDisabled,
	}
}
