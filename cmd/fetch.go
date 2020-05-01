/*
Package cmd is where all commands will be.
*/
package cmd

import (
	"log"

	"github.com/Eldius/speedtest/geolocation"
	"github.com/Eldius/speedtest/persistence"
	"github.com/Eldius/speedtest/speedtest"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("fetch called")
		cs := speedtest.OoklaClient{}
		cg := geolocation.Client{}
		l, err := cg.FetchIPLocation()
		if err != nil {
			panic(err.Error())
		}
		if servers, err := cs.FindServers(); err == nil {
			persistence.SaveAllServers(servers)
			nearestServers := cs.FindNearestServers(l.GetLocation(), 15)
			for i, s := range nearestServers {
				log.Printf("%d) %s [%v]", i, s.Name, s.ToSelectedServer(*l.GetLocation()))
			}
			log.Printf("nearest servers: %v\n", nearestServers)
		} else {
			panic(err.Error())
		}
	},
}

func init() {
	serversCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
