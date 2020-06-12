package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"flocksock/log"
	"flocksock/app/reporting"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "flocksock invoke",
	Long: `Invokes the flocksock binary as a local process.
        flocksock invoke`,

	Run: func(cmd *cobra.Command, args []string) {
        // Configure flocksock with appropriate Twitch Application key
        // Configure flocksock to be enabled on stream(s).

		// Where is the message going?
		if optOutput == "discord" {
			//reporting.SendMessageDiscord(formattedOutput)
			} else if optOutput == "stdout" {
				//reporting.SendMessageStdout(formattedOutput)
				reporting.SendJSONStdout(formattedOutput)
		} else {
				log.Errorf("optOutput: [%s] not supported (discord, stdout)", optOutput)
			}
		},
}

func init() {
	rootCmd.AddCommand(invokeCmd)
	reportCmd.Flags().StringP("command", "o", "output", "Specify which output you would like to target.")
	reportCmd.Flags().StringP("command", "w", "waterypyro", "Performs a defying act of hydroarsonism.")
}
