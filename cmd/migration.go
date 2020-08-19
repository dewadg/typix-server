package cmd

import (
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

func Migrate(client *mongo.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "db:migrate",
		Short: "Run DB migrations",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
