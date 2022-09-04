package services

import (
	"discount/models"
	"discount/util"
	"fmt"
	"net/http"
	"net/url"
)

func GetUserWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error) {
	user := models.User{}
	url := fmt.Sprintf("http://localhost:3000/user?phone=%s", url.QueryEscape(phoneNumber))
	util.HttpSendRequest(url, http.MethodGet, nil, nil, &user)

	wallet := models.Wallet{}
	url = fmt.Sprintf("http://localhost:3000/user/%d/wallet", user.ID)
	err := util.HttpSendRequest(url, http.MethodGet, nil, nil, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
