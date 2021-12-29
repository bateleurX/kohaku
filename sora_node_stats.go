package kohaku

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	db "github.com/shiguredo/kohaku/db/sqlc"
)

func toNumeric(n uint64) pgtype.Numeric {
	var num pgtype.Numeric
	num.Set(n)
	return num
}

// TODO(v): sqlc 化
func (s *Server) collectorSoraNodeErlangVMStats(c *gin.Context, stats soraNodeErlangVMStats) error {
	if err := s.InsertSoraNode(c, stats); err != nil {
		return err
	}

	for _, v := range stats.Stats {
		erlangVMStats := new(ErlangVMStats)
		if err := json.Unmarshal(v, &erlangVMStats); err != nil {
			return err
		}

		// type をみて struct をさらに別途デコードする
		switch erlangVMStats.Type {
		case "erlang-vm-memory":
			e := new(ErlangVMMemoryStats)
			if err := json.Unmarshal(v, &e); err != nil {
				return err
			}

			if err := s.query.InsertErlangVMMemoryStats(c, db.InsertErlangVMMemoryStatsParams{
				Time:              stats.Timestamp,
				SoraVersion:       stats.Version,
				SoraLabel:         stats.Label,
				SoraNodeName:      stats.NodeName,
				StatsType:         e.Type,
				TypeTotal:         toNumeric(e.Total),
				TypeProcesses:     toNumeric(e.Processes),
				TypeProcessesUsed: toNumeric(e.ProcessesUsed),
				TypeSystem:        toNumeric(e.System),
				TypeAtom:          toNumeric(e.Atom),
				TypeAtomUsed:      toNumeric(e.AtomUsed),
				TypeBinary:        toNumeric(e.Binary),
				TypeCode:          toNumeric(e.Code),
				TypeEts:           toNumeric(e.ETS),
			}); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected erlangVMStats.Type: %s", erlangVMStats.Type)
		}
	}
	return nil
}

func (s *Server) InsertSoraNode(c *gin.Context, stats soraNodeErlangVMStats) error {
	if err := s.query.InsertSoraNode(c, db.InsertSoraNodeParams{
		Timestamp: stats.Timestamp,
		Label:     stats.Label,
		Version:   stats.Version,
		NodeName:  stats.NodeName,
	}); err != nil {
		return err
	}
	return nil
}
