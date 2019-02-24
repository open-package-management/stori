package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/open-package-management/stori/core"
	storihttp "github.com/open-package-management/stori/http"

	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the registry server.",
	Run:   startServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func startServer(cmd *cobra.Command, args []string) {
	reg := core.Registry{}

	handler := storihttp.Handler(reg)

	ln, _ := net.Listen("tcp", ":5000")

	srv := http.Server{Handler: handler}
	go srv.Serve(ln)

	// Wait for a signal to stop the server.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		select {
		case sig := <-sigs:
			if sig == syscall.SIGINT {
				colorstring.Printf("\n[light_yellow][bold]SIGINT received: Initiating graceful shutdown.\n")
				srv.Shutdown(context.Background())
				colorstring.Println("[light_green][bold]Shudown complete.")
				os.Exit(0)
			}
		}
	}
}
