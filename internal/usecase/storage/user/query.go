package user

import (
	"fmt"

	"app/internal/entity"

	sq "github.com/Masterminds/squirrel"
)

func idQuery(id uint64) (string, []interface{}) {
	builder := sq.Select("*").From(usersTable).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func emailQuery(email string) (string, []interface{}) {
	builder := sq.Select("*").From(usersTable).
		Where(sq.Eq{"email": email}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func createQuery(user *entity.User) (string, []interface{}) {
	builder := sq.Insert(usersTable).
		Columns(
			"email", "name", "password", "age",
		).
		Values(
			user.Email, user.Name, user.Password, user.Age,
		).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func updateQuery(user *entity.User) (string, []interface{}) {
	builder := sq.Update(usersTable)

	if user.Email != "" {
		builder = builder.Set("email", user.Email)
	}

	if user.Name != "" {
		builder = builder.Set("name", user.Name)
	}

	if user.Password != "" {
		builder = builder.Set("password", user.Password)
	}

	if user.Age != 0 {
		builder = builder.Set("age", user.Age)
	}

	builder = builder.Where(sq.Eq{"id": user.Id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func verifyQuery(id uint64, verified bool) (string, []interface{}) {
	builder := sq.Update(usersTable).
		Set("verified", verified).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func deleteQuery(id uint64) (string, []interface{}) {
	builder := sq.Delete(usersTable).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func redisKey(id uint64) string {
	return fmt.Sprintf("users:%d", id)
}
