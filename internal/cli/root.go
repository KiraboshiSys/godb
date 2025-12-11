package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/KiraboshiSys/godb/internal/cli/create"
	"github.com/KiraboshiSys/godb/internal/cli/drop"
	"github.com/KiraboshiSys/godb/internal/cli/generate"
	initPkg "github.com/KiraboshiSys/godb/internal/cli/init"
	"github.com/KiraboshiSys/godb/internal/cli/migrate"
	"github.com/KiraboshiSys/godb/internal/cli/reset"
	"github.com/KiraboshiSys/godb/internal/cli/rollback"
	"github.com/KiraboshiSys/godb/internal/cli/version"
)

var cmd = &cobra.Command{
	Use:   "godb",
	Short: "A lightweight DB management CLI for Go projects",
	Long: `godb is a database management tool designed for Go projects.

It helps you handle database lifecycle tasks such as creation, migration, rollback, seeding, and more — all from a single, developer-friendly CLI.`,
}

func init() {
	cmd.AddCommand(create.Cmd)
	cmd.AddCommand(drop.Cmd)
	cmd.AddCommand(generate.Cmd)
	cmd.AddCommand(initPkg.Cmd)
	cmd.AddCommand(migrate.Cmd)
	cmd.AddCommand(reset.Cmd)
	cmd.AddCommand(rollback.Cmd)
	cmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
