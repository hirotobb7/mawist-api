package repository

import "github.com/hirotobb7/mawist/internal/model"

type WishListRepository interface {
	FindByUserId(userId string) ([]model.WishList, error)
	Create(wishList model.WishList) error
	Delete(wishList model.WishList) error
}
