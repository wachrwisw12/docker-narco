package services

import (
	"context"
	"errors"
	"time"

	"api-naco/db"
	"api-naco/models"

	"golang.org/x/crypto/bcrypt"
)

func AuthRegisterService(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if len(user.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	// 1️⃣ hash password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users (username, password, role_id, fullname)
		VALUES ($1, $2, $3, $4)
		RETURNING username, role_id, fullname
	`

	// 2️⃣ insert + returning
	err = db.DB.QueryRow(
		ctx,
		query,
		user.Username,
		hashedPassword,
		user.RoleId,
		user.Fullname,
	).Scan(
		&user.Username,
		&user.RoleId,
		&user.Fullname,
	)
	if err != nil {
		return nil, err
	}

	// 3️⃣ ไม่ส่ง password กลับ
	user.Password = "d"

	return &user, nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // cost = 10
	)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
