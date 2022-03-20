package main

import (
	"log"
	"passport-date/commands"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	MainCommand()
}

func MainCommand() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var rootCmd = &cobra.Command{Use: "passport-date"}
	rootCmd.AddCommand(commands.DateCmd)
	rootCmd.Execute()
}
