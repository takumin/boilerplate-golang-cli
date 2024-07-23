package command

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/command/completion"
	"github.com/takumin/boilerplate-golang-cli/internal/command/subcommand"
	"github.com/takumin/boilerplate-golang-cli/internal/config"
	"github.com/takumin/boilerplate-golang-cli/internal/metadata"
	"github.com/takumin/boilerplate-golang-cli/internal/version"
)

const (
	ExitOK int = 0
	ExitNG int = 1
)

func Main(stdout io.Writer, stderr io.Writer, stdin io.Reader, args []string) int {
	cfg := config.NewConfig(
		config.LogLevel("info"),
		config.LogFormat("json"),
	)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"l"},
			Usage:       "log level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
			Action: func(*cli.Context, string) error {
				switch cfg.LogLevel {
				case "debug":
					slog.SetLogLoggerLevel(slog.LevelDebug)
				case "info":
					slog.SetLogLoggerLevel(slog.LevelInfo)
				case "warn":
					slog.SetLogLoggerLevel(slog.LevelWarn)
				case "error":
					slog.SetLogLoggerLevel(slog.LevelError)
				default:
					return fmt.Errorf("unknown log level: %s", cfg.LogLevel)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "log-format",
			Aliases:     []string{"f"},
			Usage:       "log format",
			EnvVars:     []string{"LOG_FORMAT"},
			Value:       cfg.LogFormat,
			Destination: &cfg.LogFormat,
			Action: func(ctx *cli.Context, _ string) error {
				switch cfg.LogFormat {
				case "text":
					slog.SetDefault(slog.New(slog.NewTextHandler(ctx.App.Writer, nil)))
				case "json":
					slog.SetDefault(slog.New(slog.NewJSONHandler(ctx.App.Writer, nil)))
				default:
					return fmt.Errorf("unknown log format: %s", cfg.LogFormat)
				}
				return nil
			},
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
		Reader:               stdin,
		Writer:               stdout,
		ErrWriter:            stderr,
		ExitErrHandler:       func(ctx *cli.Context, err error) {},
	}

	ctx := context.Background()
	if err := app.RunContext(ctx, args); err != nil {
		slog.ErrorContext(ctx, "failed application", slog.Any("error", err))
		return ExitNG
	}

	return ExitOK
}
