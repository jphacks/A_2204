package myMiddleware

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

func Auth0(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Auth0のトークンがセットされているか確認
		if !strings.HasPrefix(c.Request().Header.Get("Authorization"), "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Auth0 token required")
		}

		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
		if err != nil {
			log.Fatalf("Failed to parse the issuer url: %v", err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatalf("Failed to set up the jwt validator")
		}

		token := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
		claims, err := jwtValidator.ValidateToken(c.Request().Context(), token)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Auth0 token")
		}
		c.Set("claims", claims.(*validator.ValidatedClaims))

		return next(c)
	}
}
