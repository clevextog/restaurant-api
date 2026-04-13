package main

var user User

func CreateUser(email, passwordHash string) error {
	query := `
	INSERT INTO users(email, password_hash)
	VALUES($1,$2) `
	_, err := db.Exec(query, email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := `
	SELECT id, email, password_hash, created_at
	FROM users
	WHERE email = $1`
	row := db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
