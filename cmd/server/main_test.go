package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tariq-ventura/Proyecto-go/internal/db"
	"github.com/tariq-ventura/Proyecto-go/internal/mocks"
)

func TestRunApp(t *testing.T) {
	t.Run("Success - Application runs successfully", func(t *testing.T) {
		mockDB := &mocks.MockDatabase{
			TaskMigrationFunc: func(ctx context.Context) error {
				return nil
			},
		}

		originalNewDatabase := db.NewDatabase
		db.NewDatabase = func(ctx context.Context) (db.Database, error) {
			return mockDB, nil
		}
		defer func() { db.NewDatabase = originalNewDatabase }()

		err := RunApp(context.Background())

		assert.NoError(t, err)
	})

	t.Run("Error - Database connection fails", func(t *testing.T) {
		originalNewDatabase := db.NewDatabase
		db.NewDatabase = func(ctx context.Context) (db.Database, error) {
			return nil, errors.New("failed to connect to database")
		}
		defer func() { db.NewDatabase = originalNewDatabase }()

		err := RunApp(context.Background())

		assert.Error(t, err)
		assert.Equal(t, "failed to connect to database", err.Error())
	})

	t.Run("Error - Database migration fails", func(t *testing.T) {
		mockDB := &mocks.MockDatabase{
			TaskMigrationFunc: func(ctx context.Context) error {
				return errors.New("migration error")
			},
		}

		originalNewDatabase := db.NewDatabase
		db.NewDatabase = func(ctx context.Context) (db.Database, error) {
			return mockDB, nil
		}
		defer func() { db.NewDatabase = originalNewDatabase }()

		err := RunApp(context.Background())

		assert.Error(t, err)
		assert.Equal(t, "migration error", err.Error())
	})
}
