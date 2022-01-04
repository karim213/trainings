package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "name of application",
	Short: "Short description",
	Long: `Longer description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var CfgFile string
var Verbose bool


func init() {
    RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/dagobah/config.yaml)")
    RootCmd.PersistentFlags().String("mongodb_uri", "mongodb://localhost:27017/", "Uri to connect to mongoDB")
}
