package services

import (
	"discount/models"
	"discount/util"
	"fmt"
	"net/http"
)

func GetUserWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	url := fmt.Sprintf("/wallet?phone=%s", phoneNumber)
	return wallet, util.HttpSendRequest(url, http.MethodGet, , nil, wallet)
}
