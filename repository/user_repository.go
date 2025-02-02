package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/addione/web-service/model"
	"github.com/addione/web-service/model/request"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	userCollection *mongo.Collection
	userTable      *sql.DB
}

func newUserRepository(rdi *RepositoryDIContainer) *UserRepo {
	return &UserRepo{
		userCollection: rdi.DependenciesDI.GetMongoCollection(DBName, model.UserCollectionName),
		userTable:      rdi.DependenciesDI.GetMysql(MySQLDBName),
	}
}

func (u *UserRepo) CreateNewUser(user *request.CreateUserParams) (int64, error) {
	return u.insertIntoMysqlTable(user)
}

func (u *UserRepo) insertIntoMysqlTable(user *request.CreateUserParams) (int64, error) {
	query := "INSERT into users(email, phone_number, password, status) VALUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE phone_number=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelfunc()

	stmt, err := u.userTable.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	res, err := stmt.ExecContext(ctx, user.Email, user.PhoneNumber, user.Password, model.STATUS_NEW, user.PhoneNumber)
	if err != nil {
		return 0, err
	}
	stmt.Close()
	return res.LastInsertId()
}
