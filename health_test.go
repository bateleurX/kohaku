package kohaku

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/health", strings.NewReader(""))
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.health(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())
}
