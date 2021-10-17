package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/k2rth1k/qt/pkg/config"
	"github.com/k2rth1k/qt/utilities/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SQL struct {
	Store
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

const (
	ErrDBConnectionMessage = "error connecting to DB"
	usersTable             = "accounts.users"
)

var (
	ErrNoRecordFound     = errors.New("No Record found")
	ErrUserAlreadyExists = errors.New("User already exists")
)

func NewSQL(c config.SQLConfig) (*SQL, error) {
	s := &SQL{logger: log.InitZapLog()}
	conn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.DBName, c.User, c.Pass, c.SSLMode,
	)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		s.logger.Error("Failed to connect with database due to following error ", err)
		return nil, errors.Wrap(err, ErrDBConnectionMessage)
	}
	s.logger.Info("Successfully connected to DB")
	s.db = db
	return s, nil
}

// Close invokes db.Close()
func (s *SQL) Close() error {
	return s.db.Close()
}
