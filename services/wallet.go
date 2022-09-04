package services

import (
	"discount/models"
	"discount/util"
	"fmt"
	"net/http"
)

func ChargeUserWallet(userId int, amount float64) error {
	params := models.ChargeWalletParam{
		Amount: amount,
	}

	url := fmt.Sprintf("http://localhost:3000/user/%d/wallet/charge", userId)
	return util.HttpSendRequest(url, http.MethodPost, params, nil, nil)
}
