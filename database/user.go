package database

import (
	"log"
	structs "rto/struct"

	"github.com/google/uuid"
)

// Not using any function from below
func FetchUserInfo(username string) (structs.User, error) {
	var user structs.User

	query := `select userid , username , userfirstname , userlastname , email, "password"  from public.master_user mu where username = $1`

	rows, err := DB.Query(query, username)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		rows.Scan(
			&user.UserId,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)
	}

	return user, nil
}

func Upsertuser(input structs.UserRequest, pass string) error {
	var inputArgs []interface{}

	id := uuid.New()

	query := `insert into public.master_user (userid, username, userfirstname, userlastname, createdat, email, isactive, "password" ) values(
		$1, $2, $3, $4, current_timestamp, $5, true, $6) returning userid;`

	inputArgs = append(inputArgs, id)
	inputArgs = append(inputArgs, input.UserName)
	inputArgs = append(inputArgs, input.FirstName)
	inputArgs = append(inputArgs, input.LastName)
	inputArgs = append(inputArgs, input.Email)
	inputArgs = append(inputArgs, pass)

	err2 := DB.QueryRow(query, inputArgs...).Scan(&id)
	if err2 != nil {
		log.Println("Error from DB")
	}

	return nil
}
