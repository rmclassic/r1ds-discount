package handlers

import (
	"discount/database"
	"discount/models"
	"discount/services"
	"net/http"

	"github.com/martini-contrib/render"
)

func SubmitDiscount(req *http.Request, param models.SubmitDiscountParam, db database.IDatabase, r render.Render) {
	services.SubmitDiscount(db, param.DiscountCode, param.PhoneNumber)
}

func AddDiscount(req *http.Request, param models.AddDiscountParam, db database.IDatabase, r render.Render) {
	discount := models.Discount{
		Code:          param.Code,
		TotalQuantity: param.TotalQuantity,
		Amount:        param.Amount,
	}

	err := services.AddDiscount(db, &discount)
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	r.Status(http.StatusOK)
}
