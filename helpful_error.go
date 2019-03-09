package projectcfg

type HelpfulError struct {
	error
	Help string
}

func AddHelpToError(err error, help string) HelpfulError {
	return HelpfulError{
		error: err,
		Help:  help,
	}
}

//func NewHelpfulError(msg, help string) HelpfulError {
//	return HelpfulError{
//		error: errors.New(msg),
//		Help:  help,
//	}
//}
