package errs

import (
	"runtime"
	"strconv"

	"go.uber.org/zap"
)

type Message struct {
	Message  string `json:"message"`
	File     string `json:"file"`
	Line     string `json:"line"`
	Function string `json:"function"`
}

func ErrorMessageWithLog(err error, zlog *zap.Logger) (*Message, error) {
	if err == nil {
		return nil, nil
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return nil, err
	}

	fn := runtime.FuncForPC(pc)
	funcName := ""
	if fn != nil {
		funcName = fn.Name()
	}

	zlog.Error("error", zap.String("message", err.Error()), zap.String("file", file), zap.Int("line", line), zap.String("func", funcName))

	return &Message{
		Message:  err.Error(),
		File:     file,
		Line:     strconv.Itoa(line),
		Function: funcName,
	}, nil
}

func ErrorMessageWithoutLog(err error) (*Message, error) {
	if err == nil {
		return nil, nil
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return nil, err
	}

	fn := runtime.FuncForPC(pc)
	funcName := ""
	if fn != nil {
		funcName = fn.Name()
	}

	return &Message{
		Message:  err.Error(),
		File:     file,
		Line:     strconv.Itoa(line),
		Function: funcName,
	}, nil
}

func ErrorWithLog(err error, zlog *zap.Logger) error {
	if err == nil {
		return nil
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}

	fn := runtime.FuncForPC(pc)
	funcName := ""
	if fn != nil {
		funcName = fn.Name()
	}

	zlog.Error("error", zap.String("message", err.Error()), zap.String("file", file), zap.Int("line", line), zap.String("func", funcName))

	return nil
}
