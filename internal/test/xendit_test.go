package test

import (
	"github.com/johnearl92/xendit-ta.git/internal/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type XenditTestSuite struct {
	suite.Suite
	h *handler.XenditHandler
}

func (suite *XenditTestSuite) TestXenditHealthHandler() {
	//arrange
	r, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	suite.h.GetHeartBeat(w, r)

	resp := w.Result()
	assert.Equal(suite.T(), 200, resp.StatusCode, "expected 200 status code")

}

func TestXenditTestSuite(t *testing.T) {
	suite.Run(t, new(XenditTestSuite))
}
