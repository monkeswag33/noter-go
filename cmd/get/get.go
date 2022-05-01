package get

import (
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/spf13/cobra"
)

var database *db.DB = global.DB

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource",
	Long: `This is the root command to get a resource.
You usually are going to be doing something like "noter get users", or "noter get notes"`,
}
