package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{
		db: db,
	}
}

func (r *SQLiteUserRepository) SaveUser(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword)
	if err != nil {
		return err
	}
	fmt.Println("USER PASSWORD :", user.Password, hashedPassword)

	return nil
}

func (r *SQLiteUserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User

	err := r.db.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(
		&user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *SQLiteUserRepository) GetUserByUsername(email string) (*models.User, error) {
	var user models.User

	err := r.db.QueryRow("SELECT id,username,email,password FROM users WHERE email = ?", email).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	} else {
		fmt.Println("Test : user found !")
	}

	return &user, nil
}

func (r *SQLiteUserRepository) UpdatePassword(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
	_, err := r.db.Exec("UPDATE users SET password = ? WHERE email = ?", hashedPassword, user.Email)
	if err != nil {
		return err
	}
	return nil
}