package app

import (
	"budgetBook/persistence"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Repräsentiert eine ausführbare Instanz der Anwendung. Neben der Datenbank
// und der Konfiguration wird auch ein Haupt-Command geladen, der beim Start
// ausgeführt wird. An diesen Haupt-Command werden mit bindCommands() die
// verschiedenen Subcommands angehängt.
type app struct {
	DB       persistence.Database // Implementierung von persistence.Database.
	RootCmd  *cobra.Command       // Haupt-Command, an dem die Subcommands hängen.
	Commands []*cobra.Command     // Die Subcommands der Anwendung.
	Config   *viper.Viper         // Library zum Zugriff auf Config-Werte.
}

// Erzeugt eine neue Instanz der Anwendung. Hierbei wird eine leere Instanz einer
// beliebigen persistence.Database-Implementierung angelegt. Als Haupt-Command
// wird der im Paket hinterlegte RootCmd herangezogen.
func New() *app {
	bolt := &persistence.Bolt{}
	if err := bolt.Load(); err != nil {
		return nil
	}
	return &app{
		DB:       bolt,
		RootCmd:  rootCmd,
		Commands: commands(),
	}
}

// Führt die Instanz der Anwendung aus. Sollten zu diesem Zeitpunkt noch keine
// Subcommands an den Haupt-Command gehängt worden sein, wird dies nachgeholt.
// Anschließend wird der Haupt-Command der App ausgeführt.
func (a *app) Run() {
	if !a.RootCmd.HasSubCommands() {
		a.bindCommands(a.Commands)
	}
	a.RootCmd.Execute()
}

// Hängt ein Slice von Commands an den Haupt-Command der App. Diese Funktion
// kann auch genutzt werden, um auch nachträglich beliebig viele Commands
// an den Haupt-Command zu binden.
func (a *app) bindCommands(cmds []*cobra.Command) {
	for i, _ := range cmds {
		a.RootCmd.AddCommand(cmds[i])
	}
}
