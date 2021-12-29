package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shiguredo/kohaku"

	zlog "github.com/rs/zerolog/log"
)

// curl -v --http2-prior-knowledge http://localhost:8080

// TODO(v): 特定の IP アドレスからしか受け付けないようにする
// TODO(v): 性能測定
// TODO(v): YAML 化

func NewDB(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}

func init() {
	// 設定の読み込みと、Logger の準備
	flag.Parse()
	// load config
	if err := kohaku.LoadConfigFromFlags(kohaku.ConfigFilePath); err != nil {
		log.Fatalf("config file read error: %s", err)
	}
	err := kohaku.InitLogger(kohaku.Config.LogDir, kohaku.Config.LogName, kohaku.Config.LogDebug, kohaku.Config.LogStdout)
	if err != nil {
		log.Fatalf("logger building failed. %s", err)
	}
	zlog.Info().Msg("FinishInitProcess")
}

func main() {
	connStr := kohaku.Config.TimescaleURL
	timescaleSSLMode := kohaku.Config.TimescaleSSLMode
	timescaleRootcertFile := kohaku.Config.TimescaleRootcertFile
	if (timescaleSSLMode != "") && (timescaleRootcertFile != "") {
		params := url.Values{
			"sslrootcert": {timescaleRootcertFile},
			"sslmode":     {timescaleSSLMode},
		}
		connStr = connStr + "?" + params.Encode()
	}
	pool, err := NewDB(context.Background(), connStr)
	if err != nil {
		// TODO: エラーメッセージを修正する
		// TODO(v): zlog を利用する
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	s := kohaku.NewServer(kohaku.Config, pool)

	if err := s.Start(kohaku.Config); err != nil {
		zlog.Fatal().Err(err).Msg("")
	}
}
