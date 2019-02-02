package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Haupt-Command der Anwendung. Da die Anwendung ausschließlich mit Subcommands
// arbeitet, geschieht beim Aufruf dieses Commands keine weitere Verarbeitung.
var rootCmd = &cobra.Command{
	Use: "budgetbook",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to BudgetBook`)
	},
}

// Gibt die Subcommands zurück. Die Subcommands werden beim Erstellen bzw. Laden
// der Anwendung an den Haupt-Command gehängt. Dies geschieht nicht unmittelbar
// in dieser Datei, weil der Haupt-Command der Anwendung flexibel ist und nicht
// unbedingt der hier hinterlegte RootCmd verwendet werden muss.
func commands() []*cobra.Command {
	// Erstellt eine neue Ein- oder Ausgabekategorie für Finanztransaktionen.
	addCategory := &cobra.Command{
		Use: "add-cat",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add Category")
		},
	}
	// Erstellt eine neue Finanztransaktion.
	newTransaction := &cobra.Command{
		Use: "new-tx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("New Transaction")
		},
	}
	// Gibt die Bilanz aller Transaktionen für einen gegebene Zeitraum aus.
	// Wird kein Zeitraum angegeben, werden alle vorhandenen Tranaktionen
	// ausgewertet.
	balanceSheet := &cobra.Command{
		Use: "balance",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Balance Sheet")
		},
	}
	return []*cobra.Command{
		addCategory, newTransaction, balanceSheet,
	}
}
