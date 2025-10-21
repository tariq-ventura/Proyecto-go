package logs

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerInitialization(t *testing.T) {
	var buf bytes.Buffer
	Logger.SetOutput(&buf)

	Logger.Info("Mensaje de prueba")

	assert.Contains(t, buf.String(), "Mensaje de prueba")
}

func TestLoggerFileOutput(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_logger_*.log")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	multiWriter := io.MultiWriter(tempFile, os.Stdout)
	Logger.SetOutput(multiWriter)

	Logger.Info("Mensaje de prueba en archivo")

	tempFile.Seek(0, 0)
	content, err := io.ReadAll(tempFile)
	assert.NoError(t, err)

	assert.Contains(t, string(content), "Mensaje de prueba en archivo")
}
