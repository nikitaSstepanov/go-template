package user

import (
	"app/internal/entity"
	"app/internal/entity/types"
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/brianvoe/gofakeit"
	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	gopg "github.com/gosuit/pg"
	"github.com/gosuit/rs"
	"github.com/gosuit/sl"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	repo, mock := getUser()
	defer mock.Close()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName string
		User     *entity.User
		IsError  bool
		Error    e.Error
	}{
		{
			TestName: "Success",
			User: &entity.User{
				Id:       uint64(1),
				Email:    gofakeit.Email(),
				Name:     gofakeit.Email(),
				Password: gofakeit.Letter(),
				Role:     types.USER,
				Age:      gofakeit.Number(10, 100),
				Verified: true,
			},
			IsError: false,
		},
		{
			TestName: "Not found",
			IsError:  true,
			User: &entity.User{
				Id: uint64(2),
			},
			Error: e.New("Not found", e.NotFound),
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if !tc.IsError {
				rows := mock.NewRows([]string{"id", "email", "name", "password", "age", "role", "verified"}).
					AddRow(tc.User.Id, tc.User.Email, tc.User.Name, tc.User.Password, tc.User.Age, tc.User.Role, tc.User.Verified)

				mock.ExpectQuery("SELECT *").WithArgs(tc.User.Id).WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT * ").WithArgs(tc.User.Id).WillReturnError(gopg.ErrNoRows)
			}

			user, err := repo.GetById(ctx, tc.User.Id)
			if err != nil {
				if tc.IsError {
					assert.Equal(t, tc.Error.GetCode(), err.GetCode())
				} else {
					t.Errorf("Test failing: %v", err)
				}
			} else {
				if tc.IsError {
					t.Error("Test failing: expected not nil error")
				} else {
					assert.Equal(t, tc.User, user)
				}
			}
		})
	}
}

func TestGetEmail(t *testing.T) {
	repo, mock := getUser()
	defer mock.Close()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName string
		User     *entity.User
		IsError  bool
		Error    e.Error
	}{
		{
			TestName: "Success",
			User: &entity.User{
				Id:       uint64(1),
				Email:    gofakeit.Email(),
				Name:     gofakeit.Email(),
				Password: gofakeit.Letter(),
				Role:     types.USER,
				Age:      gofakeit.Number(10, 100),
				Verified: true,
			},
			IsError: false,
		},
		{
			TestName: "Not found",
			IsError:  true,
			User: &entity.User{
				Id: uint64(2),
			},
			Error: e.New("Not found", e.NotFound),
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if !tc.IsError {
				rows := mock.NewRows([]string{"id", "email", "name", "password", "age", "role", "verified"}).
					AddRow(tc.User.Id, tc.User.Email, tc.User.Name, tc.User.Password, tc.User.Age, tc.User.Role, tc.User.Verified)

				mock.ExpectQuery("SELECT *").WithArgs(tc.User.Email).WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT *").WithArgs(tc.User.Email).WillReturnError(gopg.ErrNoRows)
			}

			user, err := repo.GetByEmail(ctx, tc.User.Email)
			if err != nil {
				if tc.IsError {
					assert.Equal(t, tc.Error.GetCode(), err.GetCode())
				} else {
					t.Errorf("Test failing: %v", err)
				}
			} else {
				if tc.IsError {
					t.Error("Test failing: expected not nil error")
				} else {
					assert.Equal(t, tc.User, user)
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	repo, mock := getUser()
	defer mock.Close()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName string
		User     *entity.User
		IsError  bool
		Error    e.Error
	}{
		{
			TestName: "Success",
			User: &entity.User{
				Id:       uint64(1),
				Email:    gofakeit.Email(),
				Name:     gofakeit.Email(),
				Password: gofakeit.Letter(),
				Role:     types.USER,
				Age:      gofakeit.Number(10, 100),
				Verified: true,
			},
			IsError: false,
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if !tc.IsError {
				mock.ExpectBegin()

				query := "INSERT INTO users "
				mock.ExpectQuery(query).WithArgs(tc.User.Email, tc.User.Name, tc.User.Password, tc.User.Age).
					WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(tc.User.Id))

				mock.ExpectCommit()
			}

			toCreate := &entity.User{
				Email:    tc.User.Email,
				Name:     tc.User.Name,
				Password: tc.User.Password,
				Age:      tc.User.Age,
			}

			err := repo.Create(ctx, toCreate)
			if err != nil {
				if tc.IsError {
					assert.Equal(t, tc.Error.GetCode(), err.GetCode())
				} else {
					t.Errorf("Test failing: %v", err)
				}
			} else {
				if tc.IsError {
					t.Error("Test failing: expected not nil error")
				} else {
					assert.Equal(t, tc.User.Id, toCreate.Id)
				}
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	repo, mock := getUser()
	defer mock.Close()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName string
		User     *entity.User
		IsError  bool
		Error    e.Error
	}{
		{
			TestName: "Success",
			User: &entity.User{
				Id:       uint64(1),
				Email:    gofakeit.Email(),
				Name:     gofakeit.Email(),
				Password: gofakeit.Letter(),
				Role:     types.USER,
				Age:      gofakeit.Number(10, 100),
				Verified: true,
			},
			IsError: false,
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if !tc.IsError {
				mock.ExpectBegin()

				query := "UPDATE users "
				mock.ExpectExec(query).
					WithArgs(tc.User.Email, tc.User.Name, tc.User.Password, tc.User.Age, tc.User.Role, tc.User.Verified, tc.User.Id).
					WillReturnResult(pgxmock.NewResult("UPDATE", 1))

				mock.ExpectCommit()
			}

			toUpdate := &entity.User{
				Id:       tc.User.Id,
				Email:    tc.User.Email,
				Name:     tc.User.Name,
				Password: tc.User.Password,
				Age:      tc.User.Age,
				Role:     tc.User.Role,
				Verified: tc.User.Verified,
			}

			err := repo.Update(ctx, toUpdate)
			if err != nil {
				if tc.IsError {
					assert.Equal(t, tc.Error.GetCode(), err.GetCode())
				} else {
					t.Errorf("Test failing: %v", err)
				}
			} else {
				if tc.IsError {
					t.Error("Test failing: expected not nil error")
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	repo, mock := getUser()
	defer mock.Close()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName string
		User     *entity.User
		IsError  bool
		Error    e.Error
	}{
		{
			TestName: "Success",
			User: &entity.User{
				Id: uint64(1),
			},
			IsError: false,
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if !tc.IsError {
				mock.ExpectBegin()

				query := "DELETE FROM users "
				mock.ExpectExec(query).
					WithArgs(tc.User.Id).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))

				mock.ExpectCommit()
			}

			err := repo.Delete(ctx, tc.User.Id)
			if err != nil {
				if tc.IsError {
					assert.Equal(t, tc.Error.GetCode(), err.GetCode())
				} else {
					t.Errorf("Test failing: %v", err)
				}
			} else {
				if tc.IsError {
					t.Error("Test failing: expected not nil error")
				}
			}
		})
	}
}

func TestIdQuery(t *testing.T) {
	id := uint64(1)
	expectedQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	expectedArgs := []any{id}

	query, args := idQuery(id)

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestEmailQuery(t *testing.T) {
	email := "test@example.com"
	expectedQuery := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTable)
	expectedArgs := []any{email}

	query, args := emailQuery(email)

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestCreateQuery(t *testing.T) {
	user := &entity.User{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "securepassword",
		Age:      30,
	}
	expectedQuery := fmt.Sprintf("INSERT INTO %s (email,name,password,age) VALUES ($1,$2,$3,$4) RETURNING id", usersTable)
	expectedArgs := []any{user.Email, user.Name, user.Password, user.Age}

	query, args := createQuery(user)

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestUpdateQuery(t *testing.T) {
	user := &entity.User{
		Id:       1,
		Email:    "updated@example.com",
		Name:     "Updated User",
		Password: "newpassword",
		Age:      25,
		Role:     types.USER,
	}
	expectedQuery := fmt.Sprintf("UPDATE %s SET email = $1, name = $2, password = $3, age = $4, role = $5, verified = $6 WHERE id = $7", usersTable)
	expectedArgs := []any{user.Email, user.Name, user.Password, user.Age, user.Role, user.Verified, user.Id}

	query, args := updateQuery(user)

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestDeleteQuery(t *testing.T) {
	id := uint64(1)
	expectedQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	expectedArgs := []any{id}

	query, args := deleteQuery(id)

	assert.Equal(t, expectedQuery, query)
	assert.Equal(t, expectedArgs, args)
}

func TestRedisKey(t *testing.T) {
	id := uint64(1)
	expectedKey := "users:1"

	key := redisKey(id)

	assert.Equal(t, expectedKey, key)
}

func getUser() (*User, pgxmock.PgxPoolIface) {
	server, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	parts := strings.Split(server.Addr(), ":")
	port, _ := strconv.ParseInt(parts[1], 10, 32)

	cfg := &rs.Config{
		Host: parts[0],
		Port: int(port),
	}

	client, err := rs.New(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		panic(err)
	}

	pg := gopg.NewWithMock(mock)

	return New(pg, client), mock
}
