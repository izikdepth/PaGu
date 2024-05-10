package main

import (
	"os"
	"os/signal"
	"syscall"

	pCmd "github.com/pagu-project/Pagu/cmd"
	"github.com/pagu-project/Pagu/config"
	"github.com/pagu-project/Pagu/discord"
	"github.com/pagu-project/Pagu/engine"
	"github.com/pagu-project/Pagu/log"
	"github.com/spf13/cobra"
)

func runCommand(parentCmd *cobra.Command) {
	run := &cobra.Command{
		Use:   "run",
		Short: "Runs a mainnet instance of RoboPac",
	}

	parentCmd.AddCommand(run)

	run.Run = func(cmd *cobra.Command, _ []string) {
		// load configuration.
		config, err := config.Load()
		pCmd.ExitOnError(cmd, err)

		// Initialize global logger.
		log.InitGlobalLogger(config.Logger)
		// starting botEngine.
		botEngine, err := engine.NewBotEngine(config)
		pCmd.ExitOnError(cmd, err)

		botEngine.RegisterAllCommands()
		botEngine.Start()

		discordBot, err := discord.NewDiscordBot(botEngine, config.DiscordBot.Token,
			config.DiscordBot)
		pCmd.ExitOnError(cmd, err)

		err = discordBot.Start()
		pCmd.ExitOnError(cmd, err)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sigChan

		if err := discordBot.Stop(); err != nil {
			pCmd.ExitOnError(cmd, err)
		}

		botEngine.Stop()
	}
}
