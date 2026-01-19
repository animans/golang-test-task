package storage

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostgresRepo_Add(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectExec(`INSERT INTO numbers\(value\) VALUES \(\$1\)`).
		WithArgs(int64(5)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewPostgresRepo(db)
	if err := repo.Add(context.Background(), 5); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestPostgresRepo_ListSorted(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"value"}).
		AddRow(int64(1)).
		AddRow(int64(2)).
		AddRow(int64(3))

	mock.ExpectQuery(`SELECT value FROM numbers ORDER BY value ASC`).
		WillReturnRows(rows)

	repo := NewPostgresRepo(db)
	got, err := repo.ListSorted(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := []int64{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("unexpected len: %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
