package services

import (
	"github.com/hirotobb7/mawist/internal/db/repositories"
	"github.com/hirotobb7/mawist/internal/models"
)

type WishListService struct {
	repository repositories.WishListRepository
}

// NewWishListService DI Containerを検討
func NewWishListService(repository repositories.WishListRepository) *WishListService {
	return &WishListService{
		repository: repository,
	}
}

func (wishListService WishListService) FindByUserId(userId string) ([]models.WishList, error) {
	return wishListService.repository.FindByUserId(userId)
}

func (wishListService WishListService) Create(wishList *models.WishList) error {
	return wishListService.repository.Create(wishList)
}

func (wishListService WishListService) Delete(wishList *models.WishList) error {
	return wishListService.repository.Delete(wishList)
}
