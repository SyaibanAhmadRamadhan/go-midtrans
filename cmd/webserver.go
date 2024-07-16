package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// webServerRestfullApi represents the hello command
var webServerRestfullApi = &cobra.Command{
	Use:   "restfull-api",
	Short: "webserver restfull api",
	Long:  `This subcommand run webserver restfull api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver starting")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//webServerRestfullApi.Flags().IntVar(&listenPort, "port", conf.Config.WebConf.RestfullApiPort, "Listen on given port")
	rootCmd.AddCommand(webServerRestfullApi)
}
