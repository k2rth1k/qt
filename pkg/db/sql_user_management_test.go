package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/k2rth1k/qt/model"
	"github.com/k2rth1k/qt/pkg/config"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCreateUser(t *testing.T) {
	t.Run("happy_path", func(t *testing.T) {
		sql := newSetup(t)
		defer func() {
			err := sql.Close()
			if err != nil {
				t.Fatal("failed to close sql connection")
			}
		}()
		user := getDummyUserModelWithoutUserId()
		createdUser, err := createUser(t, user)
		if err != nil {
			t.Fatal("failed to create user due to following error: ", err)
		}
		assert.Nil(t, err)
		assert.Equal(t, user.FirstName, createdUser.FirstName)
		assert.Equal(t, user.LastName, createdUser.LastName)
		assert.Equal(t, user.Email, createdUser.Email)
		assert.Equal(t, user.Phone, createdUser.Phone)
		assert.Equal(t, user.Password, createdUser.Password)

		err = sql.DeleteUsers(user.Email)
		if err != nil {
			t.Fatal("failed to delete user")
		}
		assert.Nil(t, err)
	})
	t.Run("user was already created with a unique email", func(t *testing.T) {
		sql := newSetup(t)
		defer func() {
			err := sql.Close()
			if err != nil {
				t.Fatal("failed to close sql connection")
			}
		}()
		user := getDummyUserModelWithoutUserId()
		createdUser, err := createUser(t, user)
		assert.Nil(t, err)
		assert.Equal(t, user.FirstName, createdUser.FirstName)
		assert.Equal(t, user.LastName, createdUser.LastName)
		assert.Equal(t, user.Email, createdUser.Email)
		assert.Equal(t, user.Phone, createdUser.Phone)
		assert.Equal(t, user.Password, createdUser.Password)
		assert.Nil(t, err)

		newlyCreatedUser, err := createUser(t, user)
		assert.Nil(t, newlyCreatedUser)
		assert.Equal(t, err, ErrUserAlreadyExists)

		err = sql.DeleteUsers(user.Email)
		if err != nil {
			t.Fatal("failed to delete user")
		}

	})
}

func newSetup(t *testing.T) *SQL {
	cfg := config.GetServiceConfig()
	store, err := NewSQL(cfg.DBConfig)
	if err != nil {
		t.Fatal("Error while configuring DB.", err)
	}
	return store
}

func getDummyUserModelWithoutUserId() *model.User {
	user := &model.User{
		FirstName: fmt.Sprintf("first-name-%s", uuid.New().String())[0:21],
		LastName:  fmt.Sprintf("last-name-%s", uuid.New().String())[0:21],
		Email:     fmt.Sprintf("email-%s@email.com", uuid.New().String()),
		Phone:     "8" + fmt.Sprintf("%08d", rand.Intn(100000000)),
		Password:  uuid.New().String()[0:15],
	}
	return user
}

func createUser(t *testing.T, user *model.User) (*model.User, error) {
	sql := newSetup(t)
	defer func() {
		err := sql.Close()
		if err != nil {
			t.Fatal("failed to close sql connection")
		}
	}()
	if user != nil {
		return sql.CreateUser(user)
	} else {
		return sql.CreateUser(getDummyUserModelWithoutUserId())
	}

}
