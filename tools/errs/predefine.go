package errs

const (
	ServerInternalError =  500
	ArgsError = 1001
)

var (
	ErrArgs = NewCodeError(ArgsError,"ArgsError")
)
