package kohaku

import (
	"net/http"

	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"
)

// TODO: ログレベル、ログメッセージを変更する
func (s *Server) collector(c *gin.Context) {
	t := c.Request.Header.Get("x-sora-stats-exporter-type")
	switch t {
	case "connection.user-agent":
		// TODO(v): validator 処理
		stats := new(soraConnectionStats)
		if err := c.Bind(stats); err != nil {
			zlog.Debug().Str("type", t).Err(err).Msg("")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.collectorUserAgentStats(c, *stats); err != nil {
			zlog.Warn().Str("type", t).Err(err).Msg("")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
		return
	case "node.erlang-vm":
		stats := new(soraNodeErlangVMStats)
		if err := c.Bind(stats); err != nil {
			zlog.Debug().Str("type", t).Err(err).Msg("")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.collectorSoraNodeErlangVMStats(c, *stats); err != nil {
			zlog.Warn().Str("type", t).Err(err).Msg("")
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
