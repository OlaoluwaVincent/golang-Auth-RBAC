package services

import (
	"errors"
	"go/auth/entities"
	"go/auth/interfaces"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(name, email, password string) (*entities.User, error)
	Login(email, password string) (accessToken string, refreshToken string, err error)
	ParseToken(tokenString string) (*jwt.Token, error)
	ValidateAccessToken(tokenString string) (*entities.User, error)
	Refresh(refreshToken string) (accessToken string, err error)
	GetByID(id uint) (*entities.User, error)
}

type authService struct {
	repo            interfaces.UserRepository
	jwtSecret       string
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewAuthService(repo interfaces.UserRepository, jwtSecret string, accessDuration, refreshDuration time.Duration) AuthService {
	return &authService{
		repo:            repo,
		jwtSecret:       jwtSecret,
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (s *authService) Register(name, email, password string) (*entities.User, error) {
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u := &entities.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(pw),
		Role:         "user",
	}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	u.PasswordHash = ""
	return u, nil
}

func (s *authService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}
	access, err := s.generateToken(user, s.accessDuration)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.generateToken(user, s.refreshDuration)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (s *authService) generateToken(user *entities.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(duration).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *authService) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
}

func (s *authService) ValidateAccessToken(tokenString string) (*entities.User, error) {
	t, err := s.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid sub")
	}
	id := uint(subFloat)
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *authService) Refresh(refreshToken string) (string, error) {
	t, err := s.ParseToken(refreshToken)
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errors.New("invalid refresh token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return "", errors.New("invalid sub")
	}
	id := uint(subFloat)
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	access, err := s.generateToken(user, s.accessDuration)
	if err != nil {
		return "", err
	}
	return access, nil
}

func (s *authService) GetByID(id uint) (*entities.User, error) {
	return s.repo.FindByID(id)
}
