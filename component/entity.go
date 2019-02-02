package component

// Alle Typen aus dem Package components stellen eine Komposition dar, die auf
// Entity basiert. Folglich kann Entity in anderen Packages - beispielsweise in
// database - stellvertretend für alle anderen Components verwendet werden.
type Entity struct {
	ID string // Die ID als String für beispielsweise transaction.Tx
}
