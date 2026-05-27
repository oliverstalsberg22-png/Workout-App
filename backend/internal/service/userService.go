package service

import (
	"context"

	"github.com/Oliverstalsy/egolifter/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Register(ctx context.Context, email string, password string) (*domain.User, error) {
	newID := uuid.NewString()
	pHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	date := time.Now()
	return u.repo.CreateUser(ctx, &domain.User{
		ID:           newID,
		Email:        email,
		PasswordHash: pHash,
		Created:      date,
	})
}

type Claim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(user *domain.User) (string, error) {
	key := []byte(os.Getenv("JWTKEY"))

	claim := Claim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString(key)
}

func (u *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return "Can't get user by email", err
	}
	pswd := user.PasswordHash
	err = bcrypt.CompareHashAndPassword(pswd, []byte(password))
	if err != nil {
		return "", err
	}
	token, err := CreateToken(user)
	if err != nil {
		return "failed to create token", err
	}
	return token, nil
}
