package data

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Name     string
	Password password
	Email    string
}

type password struct {
	plaintext string
	hash      []byte
}

type UserModel struct {
	DB *sql.DB
}

// =============== User Functions ==================

func NewUser(name string, pwd string, email string) *User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	pw := password{
		plaintext: pwd,
		hash:      hash,
	}
	user := &User{
		Name:     name,
		Password: pw,
		Email:    email,
	}
	return user
}

func (user *User) CompareHashAndPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(user.Password.hash, []byte(password))
	return err == nil
}

// ============ User Model Functions ================

func (um *UserModel) Insert(user *User) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id;`
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password.plaintext), bcrypt.DefaultCost)
	user.Password.hash = hash[:]
	args := []interface{}{user.Name, user.Email, user.Password.hash}
	err := um.DB.QueryRow(query, args...).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (um *UserModel) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT * FROM users WHERE email = ($1);`
	err := um.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password.hash,
	)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (um *UserModel) GetUserById(userId int) (*User, error) {
	user := &User{}
	query := `SELECT * FROM users WHERE id = ($1);`
	err := um.DB.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.Password.hash,
	)
	if err != nil {
		return nil, err
	}
	return user, err
}
