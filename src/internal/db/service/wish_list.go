package service

import (
	"github.com/hirotobb7/mawist/internal/db/repository"
	"github.com/hirotobb7/mawist/internal/model"
)

type WishListService struct {
	repository repository.WishListRepository
}

// NewWishListService DI Containerを検討
func NewWishListService(repository repository.WishListRepository) *WishListService {
	return &WishListService{
		repository: repository,
	}
}

func (wishListService WishListService) FindByUserId(userId string) ([]model.WishList, error) {
	return wishListService.repository.FindByUserId(userId)
}

func (wishListService WishListService) Create(wishList model.WishList) error {
	return wishListService.repository.Create(wishList)
}

func (wishListService WishListService) Delete(wishList model.WishList) error {
	return wishListService.repository.Delete(wishList)
}
