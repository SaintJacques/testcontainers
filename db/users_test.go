package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbInstance *DB

func initDB() {
	ctx := context.Background()
	connStr, err := NewTestDB(ctx, "testdb")
	if err != nil {
		panic(err)
	}

	if err := Migrate(ctx, connStr); err != nil {
		panic(err)
	}

	dbInstance, err = New(ctx, connStr)
	if err != nil {
		panic(err)
	}
}

func TestInsertUser(t *testing.T) {
	initDB()
	assert.NotNil(t, dbInstance)

	t.Run("SuccessInsertUser", func(t *testing.T) {
		u, err := dbInstance.InsertUser(context.Background(), "misha")
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, int32(1), u.ID)
	})
}

func TestSelectUser(t *testing.T) {
	initDB()
	assert.NotNil(t, dbInstance)

	u, err := dbInstance.InsertUser(context.Background(), "misha")
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, int32(1), u.ID)

	t.Run("SuccessSelectUser", func(t *testing.T) {
		u, err := dbInstance.SelectUser(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, int32(1), u.ID)
		assert.Equal(t, "misha", u.Name)
	})
}
