package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitMockDB(t *testing.T) {
	db := InitMockDB()
	assert.NotNil(t, db)

	// Check if the User table is created
	assert.True(t, db.Migrator().HasTable("users"))

	// Clean up
	err := db.Migrator().DropTable("users")
	assert.NoError(t, err)

	// Close the database connection
	dbSQL, err := db.DB()
	assert.NoError(t, err)
	dbSQL.Close()
}
