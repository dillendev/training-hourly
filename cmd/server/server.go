package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"

	hourly "github.com/dillendev/training-hourly"
)

var _ hourly.ServerInterface = (*server)(nil)

var platformIdRe = regexp.MustCompile(`^[a-z0-9-]+$`)

type server struct{}

func (s *server) ListUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, getUsers())
}

func (s *server) ListProjects(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, getProjects())
}

func (s *server) ListTimeEntries(ctx echo.Context, userId int, params hourly.ListTimeEntriesParams) error {
	var valid bool

	for _, user := range getUsers() {
		if user.Id == userId {
			valid = true
			break
		}
	}

	if !valid {
		return ctx.JSON(http.StatusNotFound, hourly.Error{Message: "user not found"})
	}

	var filtered []hourly.TimeEntry

	entries := getTimeEntries(userId)

	for _, entry := range entries {
		if params.StartDate != nil && entry.StartedAt.Before(*params.StartDate) {
			continue
		}

		if params.EndDate != nil && entry.StoppedAt.After(*params.EndDate) {
			continue
		}

		filtered = append(filtered, entry)
	}

	return ctx.JSON(http.StatusOK, filtered)
}

func (s *server) CreateToken(ctx echo.Context) error {
	var req hourly.TokenRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, hourly.Error{Message: err.Error()})
	}

	if !platformIdRe.MatchString(req.PlatformId) {
		return ctx.JSON(http.StatusBadRequest, hourly.Error{Message: "invalid platform id (can only contain lowercase letters, numbers and hyphens)"})
	}

	buf := make([]byte, 16)

	if _, err := rand.Read(buf); err != nil {
		return ctx.JSON(http.StatusInternalServerError, hourly.Error{Message: err.Error()})
	}

	token := hex.EncodeToString(buf)

	if err := storeToken(token); err != nil {
		return ctx.JSON(http.StatusInternalServerError, hourly.Error{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, hourly.TokenResponse{Token: token})
}
