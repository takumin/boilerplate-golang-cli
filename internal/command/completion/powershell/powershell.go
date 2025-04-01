package powershell

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

const powershellCompletion = `
$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
	param($commandName, $wordToComplete, $cursorPosition)
	$other = "$wordToComplete --generate-bash-completion"
	Invoke-Expression $other | ForEach-Object {
		[System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
	}
}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "powershell",
		Usage:    "powershell completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			if _, err := fmt.Fprint(ctx.App.Writer, strings.TrimSpace(powershellCompletion)+"\n"); err != nil {
				slog.ErrorContext(ctx.Context, "failed to powershell completion fprint error", slog.Any("error", err))
			}
			return nil
		},
	}
}
