package kohaku

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// TODO: ログレベル、ログメッセージを変更する
func (s *Server) Collector(c *gin.Context) {
	if c.Request.Proto != "HTTP/2.0" {
		err := fmt.Errorf("UNSUPPORTED-HTTP-VERSION: %s", c.Request.Proto)
		log.Warn().Msg(err.Error())
		// TODO: 505 を返すかの検討
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO(v): validator 処理
	exporter := new(SoraStatsExporter)
	if err := c.Bind(exporter); err != nil {
		log.Warn().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// exporter.Type の conection.remote にのみ対応する
	// TODO(v): 将来的には conneciton.sora やそれ以外にも対応していく
	if exporter.Type == "connection.remote" {
		if err := CollectorRemoteStats(s.pool, *exporter); err != nil {
			log.Warn().Msg(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
		return
	}

	log.Warn().Msgf("UNEXPECTED-TYPE: %s", exporter.Type)

	c.Status(http.StatusBadRequest)
}
