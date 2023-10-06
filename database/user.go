package database

import structs "rto/struct"

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
