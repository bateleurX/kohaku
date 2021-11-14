package kohaku

import (
	"net/http"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
)

// TODO: ログレベル、ログメッセージを変更する
func (s *Server) Collector(c *gin.Context) {
	// TODO(v): ヘッダーにて判定
	t := c.Request.Header.Get("x-sora-stats-exporter-type")
	switch t {
	case "connection.user-agent":
		// TODO(v): validator 処理
		stats := new(SoraConnectionStats)
		if err := c.Bind(stats); err != nil {
			zlog.Debug().Err(err).Msg("")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := CollectorUserAgentStats(s.pool, *stats); err != nil {
			zlog.Warn().Err(err).Msg("")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
		return
	default:
		zlog.Warn().Str("type", t).Msgf("UNEXPECTED-TYPE")
		c.Status(http.StatusBadRequest)
	}
}
