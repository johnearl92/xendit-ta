// Package handler Official-Receipt API.
//
// Official-Receipt service API endpoints
//
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/johnearl92/xendit-ta.git/internal/service"
	"github.com/johnearl92/xendit-ta.git/internal/utils"
)

//XenditHandler main handler for the eor
type XenditHandler struct {
	xenditService service.XenditService
}

// NewORHandler provides XenditHandler definition
func NewXenditHandler(p *service.XenditServiceImpl) *XenditHandler {
	return &XenditHandler{
		xenditService: p,
	}
}

// Register registers the route
func (h *XenditHandler) Register(router *mux.Router) {

	// swagger:operation GET  /account/{id} Account
	// ---
	// summary: This will get the account with the given ID
	// parameters:
	// - name: id
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.Handle("/account/{id}", getAccount(h.xenditService)).Methods(http.MethodGet)
	log.Info("[GET] /account/{id} registered")

}

func getAccount(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke getAccount")
		vars := mux.Vars(req)

		responseModel, err := xenditService.Get(vars["productId"], nil)

		if err != nil {
			utils.WriteServerError(res, "/account/{id}", "Unable to get Account",
				fmt.Sprintf("Unable to get Account. Please contact the administrator. Error: %s", err.Error()))
			return
		}

		utils.WriteEntity(res, http.StatusOK, responseModel)
		log.Debugln("end getAccount")
	}
}
