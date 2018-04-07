package main

import (
	"flag"
	"os"

	"chaingers-abci/code"
	"chaingers-abci/server"

	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

func main() {
	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	// abciPtr := flag.String("abci", "socket", "ABCI server: socket | grpc")
	flag.Parse()
	app := storage.NewStorageApplication()

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Start the listener
	srv := server.NewSocketServer(*addrPtr, app)
	srv.SetLogger(logger.With("module", "abci-server"))
	if _, err := srv.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})

}
