package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

func main() {
	db, err := sqlx.Open("sqlite3", "file:test?mode=memory")
	checkErr(err)
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS user (user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT);`
	_, err = db.Exec(sqlStmt)
	checkErr(err, sqlStmt)

	r := newRepo(db, trmsqlx.DefaultCtxGetter)
	ctx := context.Background()
	trManager := manager.Must(trmsqlx.NewDefaultFactory(db))
	u := &user{Username: "username"}

	err = trManager.Do(ctx, func(ctx context.Context) error {
		checkErr(r.Save(ctx, u))

		// example of nested transactions
		return trManager.Do(ctx, func(ctx context.Context) error {
			u.Username = "new_username"
			return r.Save(ctx, u)
		})
	})
	checkErr(err)

	userFromDB, err := r.GetByID(ctx, u.ID)
	checkErr(err)

	fmt.Println(userFromDB)
}

func checkErr(err error, args ...interface{}) {
	if err != nil {
		panic(fmt.Sprint(append([]interface{}{err}, args...)...))
	}
}

type repo struct {
	db     *sqlx.DB
	getter *trmsqlx.CtxGetter
}

func newRepo(db *sqlx.DB, c *trmsqlx.CtxGetter) *repo {
	return &repo{db: db, getter: c}
}

type user struct {
	ID       int64  `db:"user_id"`
	Username string `db:"username"`
}

func (r *repo) GetByID(ctx context.Context, id int64) (*user, error) {
	query := "SELECT * FROM user WHERE user_id = ?;"
	u := user{}

	return &u, r.getter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &u, r.db.Rebind(query), id)
}

func (r *repo) Save(ctx context.Context, u *user) error {
	query := `UPDATE user SET username = :username WHERE user_id = :user_id;`
	if u.ID == 0 {
		query = `INSERT INTO user (username) VALUES (:username);`
	}

	res, err := sqlx.NamedExecContext(ctx, r.getter.DefaultTrOrDB(ctx, r.db), r.db.Rebind(query), u)
	if err != nil {
		return err
	} else if u.ID != 0 {
		return nil
	} else if u.ID, err = res.LastInsertId(); err != nil {
		return err
	}

	return err
}
