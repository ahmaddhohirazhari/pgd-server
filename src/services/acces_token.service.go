package services

import (
	"pgd-server.com/config"
	"pgd-server.com/src/entities"
)

type AccessTokenService struct{}

func (ats *AccessTokenService) StoreAccessToken(accessToken *entities.AccessToken) error {
	if err := config.DB.Create(accessToken).Error; err != nil {
		return err
	}
	return nil
}

func (ats *AccessTokenService) GetTokenByAccessToken(accessToken string) (entities.AccessToken, error) {
	var at entities.AccessToken
	result := config.DB.Where("access_token = ?", accessToken).First(&at)
	return at, result.Error
}
