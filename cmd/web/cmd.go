package web

import (
	"github.com/spf13/cobra"
	"proj/internal/server"
	"proj/internal/server/webserver"
	"proj/public/config"
	"proj/public/context"
)

var (
	ServerCmd = &cobra.Command{
		Use:   "webserver",
		Short: "web server",
		Long:  "start a web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.NewSignalContext()
			defer cancel()
			cfg := config.GetConfig()
			app := webserver.NewApp()
			return server.Run(ctx, app, cfg.Http)
		},
	}
)