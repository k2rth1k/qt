package api

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/k2rth1k/qt/model"
	"github.com/k2rth1k/qt/pkg/db"
	qt "github.com/k2rth1k/qt/pkg/proto"
	auth "github.com/k2rth1k/qt/utilities/authentication"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

var (
	errInvalidEmail       = "invalid email"
	errInvalidPhoneNumber = "invalid phone number"
	errInvalidName        = "invalid name"
	errInternalError      = "internal error"
	errEmptyPassword      = "empty password"
	errInvalidArgument    = "invalid argument"
	errUserAlreadyExists  = status.Convert(db.ErrUserAlreadyExists).Message()
	errUserNotFound       = "user not found"
)

func (s *QuickTradeService) HelloWorld(ctx context.Context, req *qt.EmptyMessage) (*qt.Message, error) {
	s.logger.Info("HelloWorld API has been called")
	return &qt.Message{Message: "Hello World"}, nil
}

func (s *QuickTradeService) CreateUser(ctx context.Context, req *qt.CreateUserRequest) (*qt.User, error) {
	s.logger.Info("CreateUser API has been called")
	err := validateCreateUserRequest(req)
	if err != nil {
		s.logger.Error("invalid create user request", "error", err)
		return nil, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Errorw("failed to encrypt the password", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	user := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(password),
	}
	createdUser, err := s.store.CreateUser(user)
	if err != nil {
		if err == db.ErrUserAlreadyExists {
			s.logger.Errorw("user already exits", "error", err)
			return nil, status.Error(codes.AlreadyExists, errUserAlreadyExists)
		}
		s.logger.Errorw("failed to create user", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	return &qt.User{
		Email:     createdUser.Email,
		Phone:     createdUser.Phone,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		UserId:    createdUser.UserId,
	}, nil
}

func validateCreateUserRequest(req *qt.CreateUserRequest) error {
	if !govalidator.IsEmail(req.Email) || req.Email == "" {
		return status.Error(codes.InvalidArgument, errInvalidEmail)
	}
	if !govalidator.IsNumeric(req.Phone) || req.Phone == "" {
		return status.Error(codes.InvalidArgument, errInvalidPhoneNumber)
	}
	if !govalidator.IsAlpha(req.LastName) || !govalidator.IsAlpha(req.FirstName) || req.LastName == "" || req.FirstName == "" {
		return status.Error(codes.InvalidArgument, errInvalidName)
	}
	if req.Password == "" {
		return status.Error(codes.InvalidArgument, errEmptyPassword)
	}
	return nil
}

func (s *QuickTradeService) Login(ctx context.Context, req *qt.LoginRequest) (*qt.LoginResponse, error) {
	s.logger.Info("Login API has been called")
	if req.Email == "" || req.Password == "" {
		if req.Email == "" {
			s.logger.Errorw("empty email has been sent", "error", errInvalidEmail)
		} else {
			s.logger.Errorw("empty password has been sent", "error", errEmptyPassword)
		}
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	_, err := s.store.GetUserWithEmail(req.Email)
	if err != nil {
		if err == db.ErrNoRecordFound {
			s.logger.Errorw("No user record find", "error", err)
			return nil, status.Error(codes.NotFound, errUserNotFound)
		}
		s.logger.Errorw("failed to get user with email", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	ts, err := auth.CreateToken(req.Email)
	if err != nil {
		s.logger.Errorw("failed to create authentication tokens", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	saveErr := auth.CreateAuth(req.Email, ts)
	if saveErr != nil {
		s.logger.Errorw("failed to cache authentication tokens")
	}
	return &qt.LoginResponse{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}, nil
}

func (s *QuickTradeService) Refresh(ctx context.Context, req *qt.RefreshRequest) (*qt.RefreshResponse, error) {
	s.logger.Info("Refresh API has been called")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		s.logger.Errorw("Failed to parse refresh token because refresh_token expired", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		s.logger.Errorw("Token is not valid")
		return nil, status.Error(codes.Internal, errInternalError)
	}
	var response *qt.RefreshResponse
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			s.logger.Errorw("failed to convert interface to string")
			return nil, status.Error(codes.Internal, errInternalError)
		}
		userId := claims["user_id"].(string)

		////Delete the previous Access Token
		//deleted, delErr := auth.DeleteAuth(claims["access_uuid"].(string))
		//if delErr != nil || deleted == 0 { //if any goes wrong
		//	s.logger.Errorw("Failed to delete previous access Token", "error", delErr)
		//	return nil, status.Error(codes.Internal, errInternalError)
		//}

		//Delete the previous Refresh Token
		deleted, delErr := auth.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			s.logger.Errorw("Failed to delete previous Refresh Token", "error", delErr)
			return nil, status.Error(codes.Internal, errInternalError)
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := auth.CreateToken(userId)
		if createErr != nil {
			s.logger.Errorw("failed to create new token", "error", createErr)
			return nil, status.Error(codes.Internal, errInternalError)
		}
		//save the tokens metadata to redis
		saveErr := auth.CreateAuth(userId, ts)
		if saveErr != nil {
			s.logger.Errorw("failed to save tokens metadata to redis", "error", err)
			return nil, status.Error(codes.Internal, errInternalError)
		}
		response = &qt.RefreshResponse{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		return response, nil
	} else {
		s.logger.Errorw("refresh token is not valid")
		return nil, status.Error(codes.Internal, errInternalError)
	}
}

func (s *QuickTradeService) Logout(ctx context.Context, req *qt.LogoutRequest) (*qt.EmptyMessage, error) {
	au, err := auth.ExtractTokenMetadata(ctx)
	if err != nil {
		s.logger.Errorw("failed to extract token metadata", "error", err)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		s.logger.Errorw("failed to delete auth token", "error", delErr)
		return nil, status.Error(codes.Internal, errInternalError)
	}
	return &qt.EmptyMessage{}, nil
}
