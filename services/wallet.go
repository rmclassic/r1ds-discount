package services

import (
	"discount/models"
	"discount/util"
	"fmt"
	"net/http"
	"os"
)

func ChargeUserWallet(userId int, amount float64) error {
	params := models.ChargeWalletParam{
		Amount: amount,
	}

	url := fmt.Sprintf("%s/user/%d/wallet/charge", os.Getenv("WALLET_ENDPOINT"), userId)
	return util.HttpSendRequest(url, http.MethodPost, params, nil, nil)
}
