package kohaku

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
)

// TODO: ログレベル、ログメッセージを変更する
func (s *Server) health(c *gin.Context) {
	if err := s.pool.Ping(context.Background()); err != nil {
		zlog.Error().Err(err).Msg("")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
