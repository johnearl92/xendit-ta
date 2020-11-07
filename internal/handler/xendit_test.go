package handler

import (
	"github.com/johnearl92/xendit-ta.git/internal/model"
	"github.com/johnearl92/xendit-ta.git/internal/store"
	"github.com/johnearl92/xendit-ta.git/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type XenditTestSuite struct {
	suite.Suite
}

type xenditServiceMock struct {
	mock.Mock
}

func (m *xenditServiceMock) FindCommentsByOrg(orgName string, opts store.ListOpts) (*store.CommentList, error) {
	log.Infoln("Mock getting Comments by Organization name")
	args := m.Called(orgName, opts)
	ret := args.Get(0)
	return ret.(*store.CommentList), nil
}

// had to implement other methods of the interface:
//
// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) CreateAccount(account *model.Account, opts store.GetOpts) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implementedv
func (m *xenditServiceMock) UpdateAccount(account *model.Account) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) GetAccount(id string, opts store.GetOpts) (*model.Account, error) {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) FindAccountsByOrg(orgName string, opts store.ListOpts) (*store.AccountList, error) {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) CreateOrganization(organization *model.Organization, opts store.GetOpts) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) UpdateOrganization(organization *model.Organization) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) GetOrganization(id string, opts store.GetOpts) (*model.Organization, error) {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) FindByOrgName(name string, opts store.GetOpts) (*model.Organization, error) {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) CreateComment(comment *model.Comment, opts store.GetOpts) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) UpdateComment(comment *model.Comment) error {
	panic("implement me")
}

// to be implemented when mock testing for this service is implemented
func (m *xenditServiceMock) GetComment(id string, opts store.GetOpts) (*model.Comment, error) {
	panic("implement me")
}

func (suite *XenditTestSuite) TestXenditHealthHandler() {
	//arrange
	r, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	var h *XenditHandler

	h.GetHeartBeat(w, r)

	resp := w.Result()
	assert.Equal(suite.T(), 200, resp.StatusCode, "expected 200 status code")

}

func (suite *XenditTestSuite) TestXenditGetComments() {
	log.Infoln("Testing Get Comments")
	r, _ := http.NewRequest("GET", "/orgs/xendit/comments", nil)
	w := httptest.NewRecorder()

	xenditService := new(xenditServiceMock)

	item := model.Comment{
		Message:        "test",
		IsDeleted:      false,
		OrganizationID: "1",
	}
	var items []*model.Comment
	items = append(items, &item)

	list := store.NewCommentList()
	list.SetItems(items)

	// there's some issue with mux in getting the vars["org"]
	xenditService.On("FindCommentsByOrg", "", nil).Return(list)

	handler := XenditHandler{xenditService}
	handler.GetComments(w, r)
	resp := w.Result()
	assert.Equal(suite.T(), 200, resp.StatusCode, "expected 200 status code")

	// reading response
	defer resp.Body.Close()
	httpResponseBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err, "Expected no error in converting")

	var response model.CommentResponse
	utils.FromJSON(httpResponseBody, &response)
	assert.Equal(suite.T(), "test", response.Comments[0])

	xenditService.AssertExpectations(suite.T())
}

func TestXenditTestSuite(t *testing.T) {
	suite.Run(t, new(XenditTestSuite))
}
