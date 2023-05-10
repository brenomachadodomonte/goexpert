package database

import (
	"github.com/brenomachadodomonte/goexpert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUser_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Breno", "breno@email.com", "123456")

	userDb := NewUser(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
	assert.NotNil(t, user.Password, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("Breno", "breno@email.com", "123456")

	userDb := NewUser(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
	assert.NotNil(t, user.Password, userFound.Password)
}
