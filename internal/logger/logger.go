// internal/logger/logger.go
package logger

import (
	"io"
	"log"
	"os"
)

var (
	Logger *log.Logger
	file   *os.File
)

func Init(logFile string) error {
	if logFile == "" {
		// Если не указан лог-файл — логируем только в stdout
		Logger = log.New(os.Stdout, "", log.LstdFlags)
		return nil
	}

	var err error
	file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Логируем и в файл, и в консоль
	multi := io.MultiWriter(os.Stdout, file)
	Logger = log.New(multi, "", log.LstdFlags)
	return nil
}

func Close() {
	if file != nil {
		_ = file.Close()
	}
}
