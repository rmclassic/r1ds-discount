package services

import (
	"discount/models"
	"discount/util"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func GetUserWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error) {
	user := models.User{}
	url := fmt.Sprintf("%s/user?phone=%s", os.Getenv("WALLET_ENDPOINT"), url.QueryEscape(phoneNumber))
	util.HttpSendRequest(url, http.MethodGet, nil, nil, &user)

	wallet := models.Wallet{}
	url = fmt.Sprintf("%s/user/%d/wallet", os.Getenv("WALLET_ENDPOINT"), user.ID)
	err := util.HttpSendRequest(url, http.MethodGet, nil, nil, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
