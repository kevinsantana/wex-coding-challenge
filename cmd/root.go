package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wallet",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("Error to execute commands")
	}
}
