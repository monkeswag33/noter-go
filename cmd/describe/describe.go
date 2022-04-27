package describe

import "github.com/spf13/cobra"

var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Information about a specific user or note",
	Long:  "This is the root command to get specific info about a user or note",
}
