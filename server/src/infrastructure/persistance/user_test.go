package persistance

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

func Test_UserCreate(t *testing.T) {

	t.Run(
		"Createが成功する",
		func(t *testing.T) {
			// Arrange
			testuser := &entity.User{
				Name: "testName",
				Mail: "test@test.com",
			}
			db_test, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db_test.Close()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users")).
				WithArgs(testuser.Name, testuser.Mail).
				WillReturnResult(sqlmock.NewResult(1, 1))

			sqlx_db := sqlx.NewDb(db_test, "test")

			com := &database.Conn{DB: sqlx_db}
			r := &UserRepository{conn: com}
			ctx := context.Background()
			_, err = r.CreateUser(ctx, testuser)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)

}
