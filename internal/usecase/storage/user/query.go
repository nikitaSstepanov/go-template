package user

import (
	"fmt"
	"strings"

	"github.com/nikitaSstepanov/templates/golang/internal/entity"
)

func idQuery(id uint64) string {
	return fmt.Sprintf(
		`
			SELECT * FROM %s 
			WHERE id = %d;
		`, usersTable, id,
	)
}

func emailQuery(email string) string {
	return fmt.Sprintf(
		`
			SELECT * FROM %s 
			WHERE email = '%s';
		`, usersTable, email,
	)
}

func createQuery(user *entity.User) string {
	return fmt.Sprintf(
		`
			INSERT INTO %s 
				(email, name, password, age) 
			VALUES 
				('%s', '%s', '%s', %d) 
			RETURNING id;
		`,
		usersTable, user.Email, user.Name, user.Password, user.Age,
	)
}

func updateQuery(user *entity.User) string {
	toUpd := setupValues(user)

	return fmt.Sprintf(
		`
			UPDATE %s 
			SET %s 
			WHERE id = %d;
		`, usersTable, toUpd, user.Id,
	)
}

func verifyQuery(verified bool, id uint64) string {
	return fmt.Sprintf(
		`
			UPDATE %s 
			SET verified = %t
			WHERE id = %d;
		`, usersTable, verified, id,
	)
}

func deleteQuery() string {
	return fmt.Sprintf(
		`
			DELETE FROM %s 
			WHERE id = $1;
		`, usersTable,
	)
}

func redisKey(id uint64) string {
	return fmt.Sprintf("users:%d", id)
}

func setupValues(user *entity.User) string {
	toUpd := make([]string, 0)

	if user.Email != "" {
		toUpd = append(toUpd, fmt.Sprintf("email = '%s'", user.Email))
	}

	if user.Name != "" {
		toUpd = append(toUpd, fmt.Sprintf("name = '%s'", user.Name))
	}

	if user.Password != "" {
		toUpd = append(toUpd, fmt.Sprintf("password = '%s'", user.Password))
	}

	if user.Age != 0 {
		toUpd = append(toUpd, fmt.Sprintf("age = %d", user.Age))
	}

	query := strings.Join(toUpd, ", ")

	return query
}
