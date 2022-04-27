package delete

import "github.com/spf13/cobra"

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user or note",
	Long:  "Root command to delete a user or note",
}
