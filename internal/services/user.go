package services

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dangnp/go-backend-api/internal/entities"
	"github.com/dangnp/go-backend-api/internal/repositories"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
)

type UserService struct {
	BaseService
	DB       *pgxpool.Pool
	UserRepo interface {
		FindUserFilter(ctx context.Context, db *pgxpool.Pool, filter repositories.UserFilter) (*entities.User, error)
	}
}

func (s *UserService) RegisterEchoHandler(c *echo.Echo) error {

	return nil
}

func (s *UserService) Login(c echo.Context) error {

	loginRequest := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadGateway, "Bad request!")
	}

	findUserFilter := repositories.UserFilter{Email: loginRequest.Email}
	existUser, err := s.UserRepo.FindUserFilter(context.Background(), s.DB, findUserFilter)

	if err != nil {
		// Respond bad request
	}

	if existUser.ID.Int == 0 {
		// Respond bad request
	}

	// Compare password was correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password.String), []byte(loginRequest.Password)); err != nil {
		// Respones password incorrect
	}

	// Set Claim, generate token
	t, err := s.GenerateToken(existUser.ID.Int, existUser.Email.String)
	if err != nil {
		// Response cannot generate token
	}

	responseObj := struct {
		ID        int64  `json:"id"`
		CreatedAt int64  `json:"created_at"`
		Token     string `json:"token"`
	}{
		ID:        existUser.ID.Int,
		CreatedAt: existUser.CreatedAt.Int,
		Token:     t,
	}
	return c.JSON(http.StatusOK, responseObj)
}
