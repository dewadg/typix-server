package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dewadg/typix-server/cmd"
	"github.com/dewadg/typix-server/pkg/mongow"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const name = "typix-server"
const version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "typix",
	Short: "Typix server executor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Specify command!")
	},
}

func init() {
	godotenv.Load()
}

func main() {
	mongoClient, err := mongow.ConnectMongo(
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_AUTH_SOURCE"),
		os.Getenv("MONGO_SSL"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.AddCommand(cmd.Migrate(mongoClient))
	rootCmd.AddCommand(cmd.Seed(mongoClient))
	rootCmd.AddCommand(cmd.ServeHTTP(name, version, mongoClient))

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
