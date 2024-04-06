package main

import (
	robopac "github.com/robopac-project/RoboPac"
	"github.com/robopac-project/RoboPac/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "robopac-grpc",
		Version: robopac.StringVersion(),
	}

	runCommand(rootCmd)

	err := rootCmd.Execute()
	cmd.ExitOnError(rootCmd, err)
}
