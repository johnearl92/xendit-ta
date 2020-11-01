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
	"github.com/johnearl92/xendit-ta.git/internal/store"
	"net/http"
	"strings"

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

	// swagger:operation GET  /orgs/{org}/comments CommentResponse
	// ---
	// summary: This will get all the comments in an organization
	// parameters:
	// - name: org
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/CommentResponse"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.Handle("/orgs/{org}/comments", getComments(h.xenditService)).Methods(http.MethodGet)
	log.Info("[GET] /orgs/{org}/comments registered")

	// swagger:operation GET  /orgs/{org}/members AccountResponse
	// ---
	// summary: This will get all the members in an organization
	// parameters:
	// - name: org
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/AccountResponse"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.Handle("/orgs/{org}/members", getMembers(h.xenditService)).Methods(http.MethodGet)
	log.Info("[GET] /orgs/{org}/comments registered")

	// swagger:operation DELETE /orgs/{org}/comments Comment
	// ---
	// summary: This will delete all the comments in an organization
	// parameters:
	// - name: org
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
	router.Handle("/orgs/{org}/comments", deleteComments(h.xenditService)).Methods(http.MethodDelete)
	log.Info("[DELETE] /orgs/{org}/comments registered")

	// swagger:operation POST /orgs/{org}/comments org CommentReq
	// Add comment to an organization
	// ---
	// parameters:
	// - name: org
	//   in: path
	//   required: true
	//   schema:
	//     type: string
	// - name: CommentReq
	//   in: body
	//   required: true
	//   schema:
	//     $ref: "#/definitions/CommentReq"
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

func getComments(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke addComment")
		vars := mux.Vars(req)
		log.Infof("Getting Comments of Organization: %s", vars["org"])
		// TODO return list of comments

		commentList, err := xenditService.FindCommentsByOrg(strings.ToLower(vars["org"]), nil)

		if err != nil {
			log.Error(err.Error())
			utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to get Comments",
				fmt.Sprintf("Failed to get Comments. Please check organization name or contact the administrator. Error: %s", err.Error()))
			return
		}

		var comments []string

		for _, v := range commentList.Items() {
			comments = append(comments, v.Message)
		}

		commentResponse := &model.CommentResponse{
			Comments: comments,
		}

		log.Debugln("end getComments")
		utils.WriteEntity(res, http.StatusOK, commentResponse)

	}
}
func getMembers(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke getMembers")
		vars := mux.Vars(req)
		log.Infof("Getting Members of Organization: %s", vars["org"])
		var sorts []*store.Sort
		sort := store.NewSort("followers_num", store.SortOrderDesc)
		sorts = append(sorts, sort)
		listOpts := store.NewListOpts()
		listOpts.SetSort(sorts)

		accountList, err := xenditService.FindAccountsByOrg(strings.ToLower(vars["org"]), listOpts)

		if err != nil {
			log.Error(err.Error())
			utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to get Accounts",
				fmt.Sprintf("Failed to get Accounts. Please check organization name or contact the administrator. Error: %s", err.Error()))
			return
		}

		var members []model.Account

		for _, v := range accountList.Items() {
			members = append(members, *v)
		}

		accountResponse := &model.AccountResponse{
			Accounts: members,
		}

		log.Debugln("end getMembers")
		utils.WriteEntity(res, http.StatusOK, accountResponse)

	}
}

func deleteComments(xenditService service.XenditService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Debugln("invoke deleteComments")
		vars := mux.Vars(req)
		log.Infof("Getting Comments of Organization: %s", vars["org"])
		// TODO return list of comments

		commentList, err := xenditService.FindCommentsByOrg(strings.ToLower(vars["org"]), nil)

		if err != nil {
			log.Error(err.Error())
			utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to get Comments",
				fmt.Sprintf("Failed to get Comments. Please check organization name or contact the administrator. Error: %s", err.Error()))
			return
		}

		log.Debugf("Deleting comments: %d", commentList.Total())

		comments := commentList.Items()

		for k := range comments {
			comment := comments[k]
			comment.Delete()
			if err := xenditService.UpdateComment(comment); err != nil {
				log.Error(err.Error())
				utils.WriteServerError(res, "/orgs/{org}/comments", "Failed to Delete Comments",
					fmt.Sprintf("Failed to Delete Comments. Please check organization name or contact the administrator. Error: %s", err.Error()))
				return
			}
		}

		resp := &model.GenericResponse{
			Success: true,
		}

		log.Debugln("end getComments")
		utils.WriteEntity(res, http.StatusOK, resp)

	}
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
		if org, err := xenditService.FindByOrgName(strings.ToLower(vars["org"]), nil); err != nil {
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
