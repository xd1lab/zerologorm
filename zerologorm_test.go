package zerologorm

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
	gormlogger "gorm.io/gorm/logger"
)

func TestLogger_Info(t *testing.T) {
	// Initialize logger with a mock zerolog.Logger and log level
	logger := NewLogger(&zerolog.Logger{}, gormlogger.Info)

	// Create a context
	ctx := context.Background()

	// Call the Info method
	logger.Info(ctx, "test message")

	// Assert the expected output using t.Run() or t.Errorf()
	// Example: t.Run("TestLogger_Info", func(t *testing.T) {
	//              // Assertion code
	//          })
}

func TestLogger_Warn(t *testing.T) {
	// Initialize logger with a mock zerolog.Logger and log level
	logger := NewLogger(&zerolog.Logger{}, gormlogger.Warn)

	// Create a context
	ctx := context.Background()

	// Call the Warn method
	logger.Warn(ctx, "test message")

	// Assert the expected output using t.Run() or t.Errorf()
}

func TestLogger_Error(t *testing.T) {
	// Initialize logger with a mock zerolog.Logger and log level
	logger := NewLogger(&zerolog.Logger{}, gormlogger.Error)

	// Create a context
	ctx := context.Background()

	// Call the Error method
	logger.Error(ctx, "test message")

	// Assert the expected output using t.Run() or t.Errorf()
}

func TestLogger_Trace(t *testing.T) {
	// Initialize logger with a mock zerolog.Logger and log level
	logger := NewLogger(&zerolog.Logger{}, gormlogger.Info)
	// Set the SlowThreshold to a non-zero value if necessary

	// Create a context
	ctx := context.Background()

	// Call the Trace method with test arguments
	logger.Trace(ctx, time.Now(), func() (string, int64) {
		return "SELECT * FROM table", 10
	}, nil)

	// Assert the expected output using t.Run() or t.Errorf()
}
