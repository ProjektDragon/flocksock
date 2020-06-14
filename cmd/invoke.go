package cmd

import (
	"github.com/spf13/cobra"

	"flocksock/log"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "flocksock invoke",
	Long: `Invokes the flocksock binary as a local process.
        flocksock invoke`,

	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve variables passed into the CLI
		optOutput, _ := cmd.Flags().GetString("output")
		//optEasterEgg, _ := cmd.Flags().GetString("easterEgg")

        // Configure flocksock with appropriate Twitch Application key
        // Configure flocksock to be enabled on stream(s).

		// Where is the message going?
		if optOutput == "discord" {
			//reporting.SendMessageDiscord(formattedOutput)
			} else if optOutput == "stdout" {
				//reporting.SendMessageStdout(formattedOutput)
				//reporting.SendJSONStdout(formattedOutput)
		} else {
				log.Errorf("optOutput: [%s] not supported (discord, stdout)", optOutput)
			}
		},
}

func init() {
	rootCmd.AddCommand(invokeCmd)
	invokeCmd.Flags().StringP("output", "o", "output", "Specify which output you would like to target.")
	//invokeCmd.Flags().StringP("easterEgg", "w", "waterypyro", "Performs a defying act of hydroarsonism.")
}
