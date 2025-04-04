package user

import (
	"fmt"

	"app/internal/entity"

	sq "github.com/Masterminds/squirrel"
)

func idQuery(id uint64) (string, []any) {
	builder := sq.Select("*").From(usersTable).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func emailQuery(email string) (string, []any) {
	builder := sq.Select("*").From(usersTable).
		Where(sq.Eq{"email": email}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func createQuery(user *entity.User) (string, []any) {
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

func updateQuery(user *entity.User) (string, []any) {
	builder := sq.Update(usersTable).
		Set("email", user.Email).Set("name", user.Name).
		Set("password", user.Password).Set("age", user.Age).
		Set("role", user.Role).Set("verified", user.Verified).
		Where(sq.Eq{"id": user.Id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func deleteQuery(id uint64) (string, []any) {
	builder := sq.Delete(usersTable).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, _ := builder.ToSql()

	return query, args
}

func redisKey(id uint64) string {
	return fmt.Sprintf("users:%d", id)
}
