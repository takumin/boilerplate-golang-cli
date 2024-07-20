package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/command/completion"
	"github.com/takumin/boilerplate-golang-cli/internal/command/subcommand"
	"github.com/takumin/boilerplate-golang-cli/internal/config"
	"github.com/takumin/boilerplate-golang-cli/internal/metadata"
	"github.com/takumin/boilerplate-golang-cli/internal/version"
)

func main() {
	cfg := config.NewConfig()

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"l"},
			Usage:       "log level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
		},
	}

	cmds := []*cli.Command{
		completion.NewCommands(cfg, flags),
		subcommand.NewCommands(cfg, flags),
	}

	app := &cli.App{
		Name:                 metadata.AppName(),
		Usage:                metadata.AppDesc(),
		Version:              fmt.Sprintf("%s (%s)", version.Version(), version.Revision()),
		Authors:              []*cli.Author{{Name: metadata.AuthorName()}},
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
		Before:               before(cfg),
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("failed application", slog.Any("error", err))
		os.Exit(1)
	}
}

func before(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		opts := slog.HandlerOptions{}
		switch cfg.LogLevel {
		case "debug":
			opts.Level = slog.LevelDebug
		case "info":
			opts.Level = slog.LevelInfo
		case "warn":
			opts.Level = slog.LevelWarn
		case "error":
			opts.Level = slog.LevelError
		default:
			return fmt.Errorf("unknown log level: %s", cfg.LogLevel)
		}
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &opts)))
		return nil
	}
}
