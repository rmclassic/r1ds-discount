package handlers

import (
	"discount/database"
	"discount/models"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetDiscoutUsageReport(req *http.Request, db database.IDatabase, params martini.Params, r render.Render) {
	discountId, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	report := &[]models.DiscountUsage{}
	if err = db.DiscountUsageGet(models.DiscountUsage{DiscountID: discountId}, report); err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	r.JSON(http.StatusOK, report)
}
