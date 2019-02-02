package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Haupt-Command der Anwendung. Da die Anwendung ausschließlich mit Subcommands
// arbeitet, geschieht beim Aufruf dieses Commands keine weitere Verarbeitung.
var rootCmd = &cobra.Command{
	Use: `budgetbook <Command>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to BudgetBook`)
	},
}

// Erstellt die Subcommands. Die Subcommands werden beim Erstellen bzw. Laden
// der Anwendung an den Haupt-Command gehängt. Dies geschieht nicht unmittelbar
// in dieser Datei, weil der Haupt-Command der Anwendung flexibel ist und nicht
// unbedingt der hier hinterlegte RootCmd verwendet werden muss.
func commands() []*cobra.Command {
	return nil
}
