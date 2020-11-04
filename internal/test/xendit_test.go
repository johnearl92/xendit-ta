package test

import (
	"github.com/johnearl92/xendit-ta.git/internal/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXenditHandler(t *testing.T) {
	//arrange
	r, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	var h *handler.XenditHandler
	h.GetHeartBeat(w, r)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")

}
