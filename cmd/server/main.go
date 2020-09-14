package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dangnp/go-backend-api/cfg"

	"github.com/dangnp/go-backend-api/internal/services"
	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("cannot load .env file")
	}
}

func main() {

	serviceCfg := cfg.ServiceCfg{
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		PostgresURI:  os.Getenv("POSTGRES_URI"),
	}
	e := echo.New()

	// DB pool connection
	dbPool := NewDBPool(serviceCfg)

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	jwtConfig := middleware.JWTConfig{
		Claims:     &services.JWTCustomClaim{},
		SigningKey: serviceCfg.JWTSecretKey,
		Skipper: func(c echo.Context) bool {
			path := c.Path()
			for _, skip := range services.SkipJWTAuth {
				if strings.HasPrefix(path, skip) {
					return true
				}
			}
			return false
		},
	}
	e.Use(middleware.JWTWithConfig(jwtConfig))

	// Start server
	// e.Logger.Fatal(e.Start(":1323"))
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

}

// NewDBPool connection to postgres
func NewDBPool(cfg cfg.ServiceCfg) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(cfg.PostgresURI)
	if err != nil {
		panic(err.Error())
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		panic(err.Error())
	}

	return pool
}
