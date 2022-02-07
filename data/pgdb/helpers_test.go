package pgdb

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
)

func newTestDB(t *testing.T) *pgx.Conn {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5439/newnews_test?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close(context.Background())
	})

	return db
}
