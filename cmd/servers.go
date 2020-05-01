/*
Package cmd is where all commands will be.
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// serversCmd represents the servers command
var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("servers called")
	},
}

func init() {
	rootCmd.AddCommand(serversCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serversCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
