package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/ChrisTheAbysswalker/rootly-backend/models"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

type Claims struct {
	UserID uint `json:"user_id"`
	RolID  uint `json:"rol_id"`
	Nombre  string `json:"nombre"`
	jwt.RegisteredClaims
}

func (s *AuthService) Register(user *models.Usuario) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	return s.DB.Create(user).Error
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.Usuario
	if err := s.DB.Preload("UsuarioRol").Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	secret := os.Getenv("JWT_SECRET")
	claims := &Claims{
		UserID: user.ID,
		RolID:  user.IDRol,
		Nombre: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}