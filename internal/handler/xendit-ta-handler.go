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
	"github.com/johnearl92/xendit-ta.git/internal/model"
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
func NewXenditHandler(p service.XenditService) *XenditHandler {
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

	// swagger:operation POST /orgs/{org-name}/comments orgs comment
	// Add comment a given organization
	// ---
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.Handle("/orgs/{org}/comments", addComment(h.xenditService)).Methods(http.MethodPost)
	log.Info("[POST] /e-or/generate registered")

}

func addComment(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke addComment")
		vars := mux.Vars(req)
		var data model.CommentReq
		log.Infoln("Reading Request")
		if err := utils.ReadEntity(req, &data); err != nil {
			log.Errorln(err)
			utils.WriteError(res, http.StatusBadRequest, err)
			return
		}

		log.Infoln("Validating Request")
		if err := data.ValidateComment(); err != nil {
			log.Errorln(err)
			utils.WriteError(res, http.StatusBadRequest, err)
			return
		}

		log.Infoln("Getting Organization")
		if org, err := xenditService.GetOrganization(vars["org"], nil); err != nil {
			log.Error(err.Error())
			utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to Organization",
				fmt.Sprintf("Failed to get Organization, it might not exist. Please contact the administrator. Error: %s", err.Error()))
			return
		} else {
			log.Debugf("Successfully retreived organization: " + org.Name)
			log.Infoln("Saving Comment...")
			var comment = &model.Comment{
				OrganizationID: org.ID,
				Message:        data.Comment,
				IsDeleted:      false,
			}

			if err := xenditService.CreateComment(comment, nil); err != nil {
				log.Error(err.Error())
				utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to save comment",
					fmt.Sprintf("Failed to save comment. Please contact the administrator. Error: %s", err.Error()))
				return
			}

		}

		resp := &model.GenericResponse{
			Success: true,
		}

		log.Debugln("end addComment")
		utils.WriteEntity(res, http.StatusOK, resp)

	}
}

func getAccount(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke getAccount")
		vars := mux.Vars(req)

		responseModel, err := xenditService.GetAccount(vars["productId"], nil)

		if err != nil {
			utils.WriteServerError(res, "/account/{id}", "Unable to get Account",
				fmt.Sprintf("Unable to get Account. Please contact the administrator. Error: %s", err.Error()))
			return
		}

		utils.WriteEntity(res, http.StatusOK, responseModel)
		log.Debugln("end getAccount")
	}
}
