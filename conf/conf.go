package conf

// These constants allow the strings to be accessed easily from other packages.
// However, all messages will be provided by a dedicated config library soon.
const (
	// Command-associated error messages. They are mainly used by the CLI types
	// which convert and use the interchangeable commands (see cli/cli.go).
	CmdTypeAssertionFailed = "assertion to CLI library's command type failed"
	CmdNotFoundInCtrSet    = "command not found in CLI library's container set"
	CmdParamAlreadyExists  = "specified param already exists for the command"
	CmdOptionAlreadyExists = "specified option already exists for the command"
	// Database-related error messages.
	DbNotOpened      = "database implementation is not opened or nil"
	TableNotExisting = "specified table does not exist"
	// Database-related config values.
	DbName        = "budgetbook"
	CatTable      = "cats"
	TxTable       = "txs"
	BoltDBTimeout = 1200
)
