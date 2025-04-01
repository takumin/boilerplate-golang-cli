package fish

import (
	"fmt"
	"log/slog"

	"github.com/urfave/cli/v3"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "fish",
		Usage:    "fish completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			fish, err := ctx.App.ToFishCompletion()
			if err != nil {
				return err
			}
			if _, err := fmt.Fprint(ctx.App.Writer, fish); err != nil {
				slog.ErrorContext(ctx.Context, "failed to fish completion fprint error", slog.Any("error", err))
			}
			return nil
		},
	}
}
