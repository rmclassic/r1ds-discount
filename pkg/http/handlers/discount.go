package handlers

import (
	"discount/database"
	"discount/models"
	"discount/services"
	"net/http"

	"github.com/martini-contrib/render"
)

func SubmitDiscount(req *http.Request, param models.SubmitDiscountParam, db database.IDatabase, r render.Render) {
	services.SubmitDiscount(db, 0, 0)
}
