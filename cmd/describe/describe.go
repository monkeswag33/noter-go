package describe

import (
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/spf13/cobra"
)

var database *db.DB = global.DB

var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Information about a specific user or note",
	Long:  "This is the root command to get specific info about a user or note",
}
