package repositories

import (
	"github.com/hirotobb7/mawist/internal/models"
)

type WishListRepository interface {
	FindByUserId(userId string) ([]models.WishList, error)
	Create(wishList models.WishList) error
	Delete(wishList models.WishList) error
}
