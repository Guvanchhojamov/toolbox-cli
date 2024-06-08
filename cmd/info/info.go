package info

import (
	"github.com/spf13/cobra"
	"os"
)

var InfoCmd = &cobra.Command{
	Use: "info",

	Short:   "Info operations command ",
	Long:    "For info operations command run",
	Example: "toolbox info ",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := InfoCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	InfoCmd.Flags().BoolP("help", "h", true, "show help")
	InfoCmd.AddCommand(systemCmd)
	fmt.Println("this line printed from NeoVim editor.")
}

