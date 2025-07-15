package log

import "fmt"

// type LogHandler Struct {
// 	ErrorHandler(err error, info string) bool
// 	WarnHandler(err error, info string) bool
// 	InfoHandler(msg string)
// }

// func NewHandler(l Logger) LogHandler {
// 	return &logHandler{
// 		logger: l,
// 	}
// }

type LogHandler struct {
	logger Logger
}

func (lh *LogHandler) SetLogger(l Logger) {
	lh.logger = l
}

func (lh *LogHandler) ErrorHandler(err error, info string) bool {
	hasError := err != nil
	if lh != nil && lh.logger != nil && hasError {
		lh.logger.ErrorPkg(fmt.Errorf("%s: %w", info, err))
	}
	return hasError
}

func (lh *LogHandler) WarnHandler(err error, info string) bool {
	hasError := err != nil
	if lh != nil && lh.logger != nil && hasError {
		lh.logger.WarnPkg(fmt.Errorf("%s: %w", info, err))
	}
	return hasError
}

func (lh *LogHandler) InfoHandler(msg string) {
	if lh != nil && lh.logger != nil {
		lh.logger.Info(msg)
	}
}
