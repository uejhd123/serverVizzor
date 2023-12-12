package cli

import (
	"serverVizzor/internal/server"

	"github.com/spf13/cobra"
)

func StartCliClient() {
	rootCmd := &cobra.Command{
		Use:     "serverVizzor",
		Version: "1.0.0",
	}

	rootCmd.AddCommand(initServer())

	rootCmd.Execute()
}

func initServer() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "start",
		Short: "запуск сервера",
		Long:  "Запускает приложение на localhost:8080",
	}

	listCmd.Run = func(cmd *cobra.Command, args []string) {
		server.StartServer()
	}
	return listCmd
}
