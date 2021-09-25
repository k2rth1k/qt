package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/k2rth1k/qt/model"
	"github.com/k2rth1k/qt/pkg/config"
	"github.com/k2rth1k/qt/pkg/db"
	qt "github.com/k2rth1k/qt/pkg/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"testing"
)

func TestDeleteUsers(t *testing.T) {
	sql := NewSetup(t)
	defer func() {
		err := sql.Close()
		if err != nil {
			t.Fatal("failed to close sql connection")
		}
	}()
	sql.DeleteAllUsers()
}

func TestQuickTradeService_CreateUser(t *testing.T) {
	ctx := context.Background()
	t.Run("happy path", func(t *testing.T) {
		svc, err := NewQuickTradeService()
		if err != nil {
			t.Fatal("failed to create quick trade service", err)
		}
		sql := NewSetup(t)
		defer func() {
			err := sql.Close()
			if err != nil {
				t.Fatal("failed to close sql connection")
			}
		}()
		req := getCustomerUserRequest()
		user, err := svc.CreateUser(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, req.FirstName, user.FirstName)
		assert.Equal(t, req.LastName, user.LastName)
		assert.Equal(t, req.Email, user.Email)
		assert.Equal(t, req.Phone, user.Phone)

		err = sql.DeleteAllUsers()
		if err != nil {
			t.Fatal("failed to delete user")
		}
		assert.Nil(t, err)
	})
	t.Run("Db failure", func(t *testing.T) {
		mockSql := &MockSql{}
		mockSql.On("CreateUser", mock.Anything, mock.Anything).Return(nil, status.Error(codes.Internal, errInternalError))

		svc, err := NewQuickTradeService()
		if err != nil {
			t.Fatal("failed to create quick trade service", err)
		}
		svc.store = mockSql
		user, err := svc.CreateUser(ctx, getCustomerUserRequest())
		assert.Nil(t, user)
		assert.NotNil(t, errInternalError, status.Convert(err).Message())
		assert.NotNil(t, codes.Internal, status.Convert(err).Code())
	})
	t.Run("invalid arguments", func(t *testing.T) {
		testCases := map[string]createUserTests{}
		emailTestCase := getCustomerUserRequest()
		emailTestCase.Email = ""
		emailError := status.Error(codes.InvalidArgument, errInvalidEmail)
		testCases["invalid email"] = createUserTests{
			req: emailTestCase,
			err: emailError,
		}

		phoneTestCase := getCustomerUserRequest()
		phoneTestCase.Phone = ""
		phoneError := status.Error(codes.InvalidArgument, errInvalidPhoneNumber)
		testCases["invalid phone"] = createUserTests{
			req: phoneTestCase,
			err: phoneError,
		}

		passwordTestCase := getCustomerUserRequest()
		passwordTestCase.Password = ""
		passwordError := status.Error(codes.InvalidArgument, errEmptyPassword)
		testCases["invalid password"] = createUserTests{
			req: passwordTestCase,
			err: passwordError,
		}

		firstNameTestCase := getCustomerUserRequest()
		firstNameTestCase.FirstName = "dawha7w"
		firstNameError := status.Error(codes.InvalidArgument, errInvalidName)
		testCases["invalid first name"] = createUserTests{
			req: firstNameTestCase,
			err: firstNameError,
		}

		lastNameTestCase := getCustomerUserRequest()
		lastNameTestCase.LastName = "asjba62"
		lastNameError := status.Error(codes.InvalidArgument, errInvalidName)
		testCases["invalid last names"] = createUserTests{
			req: lastNameTestCase,
			err: lastNameError,
		}
		svc, err := NewQuickTradeService()
		if err != nil {
			t.Fatal("failed to create quick trade service", err)
		}
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				res, err := svc.CreateUser(ctx, testCase.req)
				assert.Nil(t, res)
				assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
				assert.Equal(t, status.Convert(testCase.err).Code(), status.Convert(err).Code())
			})
		}
	})
	t.Run("if user already exists", func(t *testing.T) {
		svc, err := NewQuickTradeService()
		if err != nil {
			t.Fatal("failed to create quick trade service", err)
		}
		sql := NewSetup(t)
		defer func() {
			err := sql.Close()
			if err != nil {
				t.Fatal("failed to close sql connection")
			}
		}()
		req := getCustomerUserRequest()
		user, err := svc.CreateUser(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, req.FirstName, user.FirstName)
		assert.Equal(t, req.LastName, user.LastName)
		assert.Equal(t, req.Email, user.Email)
		assert.Equal(t, req.Phone, user.Phone)

		res, err := svc.CreateUser(ctx, req)
		assert.Nil(t, res)
		assert.Equal(t, codes.AlreadyExists, status.Convert(err).Code())
		assert.Equal(t, errUserAlreadyExists, status.Convert(err).Message())

		err = sql.DeleteAllUsers()
		if err != nil {
			t.Fatal("failed to delete user")
		}
		assert.Nil(t, err)
	})
}

type createUserTests struct {
	req *qt.CreateUserRequest
	err error
}

func NewSetup(t *testing.T) *db.SQL {
	cfg := config.GetServiceConfig()
	store, err := db.NewSQL(cfg.DBConfig)
	if err != nil {
		t.Fatal("Error while configuring DB.", err)
	}
	return store
}

func getCustomerUserRequest() *qt.CreateUserRequest {
	return &qt.CreateUserRequest{
		FirstName: fmt.Sprintf("firstName"),
		LastName:  fmt.Sprintf("lastName"),
		Email:     fmt.Sprintf("email-%s@email.com", uuid.New().String()),
		Phone:     "8" + fmt.Sprintf("%08d", rand.Intn(100000000)),
		Password:  uuid.New().String()[0:15],
	}
}

type MockSql struct {
	*db.SQL
	mock.Mock
}

func (c *MockSql) CreateUser(user *model.User) (*model.User, error) {
	args := c.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}
