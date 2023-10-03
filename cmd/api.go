package cmd

import (
	envconfig "github.com/kevinsantana/wex-coding-challenge/internal/config"
	"github.com/kevinsantana/wex-coding-challenge/internal/infra/database"
	"github.com/kevinsantana/wex-coding-challenge/internal/server"
	"github.com/kevinsantana/wex-coding-challenge/pkg/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the http server.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		log.WithField("project_version", version.PROJECT_VERSION)

		conf := envconfig.InitConfig(ctx)
		db := database.InitDb(ctx, conf)
		server.Run(ctx, server.HttpConfig{
			Cfg: conf,
			Db:  db,
		})
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
