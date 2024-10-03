package persist_test

import (
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
		Driver: "sqlite",
		Host:   ":memory:",
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

func (s *UserSuite) TestCreate() {
	repo := persist.NewUserRepo(s.db)
	user := user_model.NewUser("test")
	err := repo.Create(user)
	if !s.NoError(err) {
		return
	}
	if !s.NotEmpty(user.ID) {
		return
	}
	s.Equal(user.Name, "test")
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
