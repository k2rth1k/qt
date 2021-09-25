package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/k2rth1k/qt/model"
	"github.com/k2rth1k/qt/utilities/resource"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *SQL) CreateUser(user *model.User) (*model.User, error) {
	_, err := s.GetUserWithEmail(user.Email)
	if err == ErrNoRecordFound {
		return s.createUser(user, nil)
	}
	return nil, ErrUserAlreadyExists

}
func newResourceID(prefix string) (string, error) {
	return resource.NewResourceID(prefix)
}

func (s *SQL) createUser(user *model.User, tx *sqlx.Tx) (*model.User, error) {
	userID, err := newResourceID("u")

	if err != nil {
		return nil, errors.Wrap(err, "generate user id")
	}
	createBuilder, args, err := sq.Insert(usersTable).Columns("user_id", "first_name", "last_name", "email", "password", "phone").Values(userID, user.FirstName, user.LastName, user.Email, user.Password, user.Phone).ToSql()
	if err != nil {
		msg := "failed to build insert stmt"
		return nil, errors.Wrap(err, msg)
	}

	createBuilder = s.db.Rebind(createBuilder)
	// Safe from Check tx == nil
	if tx == nil {
		_ = s.db.QueryRowx(createBuilder, args...)
	} else {
		_ = tx.QueryRowx(createBuilder, args...)
	}

	insertedUser, err := s.getUserWithEmail(user.Email, tx)
	if err != nil {
		if status.Convert(err).Message() == status.Convert(ErrNoRecordFound).Message() {
			return nil, errors.Wrap(err, "failed to create user")
		}
		return nil, errors.Wrap(err, "get inserted user")
	}

	return insertedUser, nil
}

func (s *SQL) getUserWithEmail(email string, tx *sqlx.Tx) (*model.User, error) {
	selectStmt, args, err := sq.Select("*").From(usersTable).Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "encountered error while building sql to fetch user")
	}

	selectStmt = s.db.Rebind(selectStmt)

	var user []*model.User
	if tx == nil {
		err = s.db.Select(&user, selectStmt, args...)
	} else {
		err = tx.Select(&user, selectStmt, args...)
	}
	if err != nil {
		return nil, errors.Wrap(err, "fail , find user by user id")
	}

	if len(user) == 0 {
		return nil, ErrNoRecordFound
	}

	return user[0], nil
}

func (s *SQL) GetUserWithEmail(email string) (*model.User, error) {
	return s.getUserWithEmail(email, nil)
}

func (s *SQL) DeleteUsers(email string) error {
	deleteStmt, args, err := sq.Delete(usersTable).Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return errors.New("encountered error while generating sql for user deletion from users table")
	}

	deleteStmt = s.db.Rebind(deleteStmt)

	_, err = s.db.Exec(deleteStmt, args...)

	if err != nil {
		return errors.New("encountered error while deleting user from users table")
	}

	return nil
}

func (s *SQL) DeleteAllUsers() error {
	deleteStmt, args, err := sq.Delete(usersTable).ToSql()
	if err != nil {
		return errors.New("encountered error while generating sql for user deletion from users table")
	}

	deleteStmt = s.db.Rebind(deleteStmt)

	_, err = s.db.Exec(deleteStmt, args...)

	if err != nil {
		return errors.New("encountered error while deleting user from users table")
	}

	return nil
}
