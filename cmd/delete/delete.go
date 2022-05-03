package delete

import (
	"github.com/monkeswag33/noter-go/db"
	"github.com/spf13/cobra"
)

var database *db.DB

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user or note",
	Long:  "Root command to delete a user or note",
}
