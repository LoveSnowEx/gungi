package persist_test

import (
	"context"
	"testing"

	databasetool "github.com/LoveSnowEx/gotool/database"
	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
)

type UserSuite struct {
	suite.Suite
	db       *bun.DB
	userRepo user_repo.Repo
}

func (s *UserSuite) SetupSuite() {
	const connection = "test"
	config.Database.DefaultConnection = connection
	config.Database.Connections[connection] = &databasetool.Config{
		Driver:   "sqlite",
		Database: ":memory:",
	}
	bootstrap.SetupDB()
	s.db = database.Default()
	s.userRepo = persist.NewUserRepo(
		database.Default(),
	)
}

func (s *UserSuite) TearDownSuite() {
	s.db.Close()
}

func (s *UserSuite) AfterTest(suiteName, testName string) {
	_, err := s.db.NewTruncateTable().
		Model((*user_model.User)(nil)).
		Exec(context.Background())
	s.NoError(err)
}

func (s *UserSuite) TestCreate() {
	repo := persist.NewUserRepo(s.db)
	names := [...]string{"test", "test2", "test3"}

	for _, name := range names {
		user := user_model.NewUser(name)
		err := repo.Save(user)
		if !s.NoError(err) {
			return
		}
		if !s.NotEmpty(user.ID) {
			return
		}
	}
	users := []user_model.User{}
	err := s.db.NewSelect().Model(&users).Scan(context.Background())
	if !s.NoError(err) {
		return
	}
	s.Equal(len(names), len(users))
	for i, user := range users {
		s.Equal(names[i], user.Name)
	}
}

func (s *UserSuite) TestSave() {
	repo := persist.NewUserRepo(s.db)
	names := [...]string{"test", "test2", "test3"}

	user := user_model.NewUser("original")
	_ = repo.Save(user)
	originID := user.ID
	for _, name := range names {
		user.Name = name
		err := repo.Save(user)
		if !s.NoError(err) {
			return
		}
		if !s.Equal(originID, user.ID) {
			return
		}
	}
	users := []user_model.User{}
	err := s.db.NewSelect().Model(&users).Scan(context.Background())
	if !s.NoError(err) {
		return
	}
	s.Equal(1, len(users))
	s.Equal(names[len(names)-1], users[0].Name)
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
