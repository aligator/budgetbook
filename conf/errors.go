package conf

const (
	CmdTypeAssertionFailed = "assertion to CLI library's command type failed"
	CmdNotFoundInCtrSet    = "command not found in CLI library's container set"
	CmdParamAlreadyExists  = "specified param already exists for the command"
	CmdOptionAlreadyExists = "specified option already exists for the command"

	DbNotOpened      = "database implementation is not opened or nil"
	TableNotExisting = "specified table does not exist"
)