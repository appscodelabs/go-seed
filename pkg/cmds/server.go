package cmds

import (
	"github.com/appscode/analytics/pkg/analytics"
	"github.com/appscode/analytics/pkg/server"
	"github.com/spf13/cobra"
)

func NewCmdServer(version string) *cobra.Command {
	srv := hostfacts.Server{
		APIAddress:      ":9844",
		OpsAddress:      ":56790",
		EnableAnalytics: true,
	}
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run server",
		DisableAutoGenTag: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			if srv.EnableAnalytics {
				analytics.Enable()
			}
			analytics.SendEvent("analytics", "started", version)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			analytics.SendEvent("analytics", "stopped", version)
		},
		Run: func(cmd *cobra.Command, args []string) {
			srv.ListenAndServe()
		},
	}

	cmd.Flags().StringVar(&srv.APIAddress, "address", srv.APIAddress, "Http server address")
	cmd.Flags().StringVar(&srv.CACertFile, "caCertFile", srv.CACertFile, "File containing CA certificate")
	cmd.Flags().StringVar(&srv.CertFile, "certFile", srv.CertFile, "File container server TLS certificate")
	cmd.Flags().StringVar(&srv.KeyFile, "keyFile", srv.KeyFile, "File containing server TLS private key")

	cmd.Flags().StringVar(&srv.OpsAddress, "web-addr", srv.OpsAddress, "Address to listen on for web interface and telemetry.")
	cmd.Flags().BoolVar(&srv.EnableAnalytics, "analytics", srv.EnableAnalytics, "Send analytical events to Google Analytics")
	return cmd
}