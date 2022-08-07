package api

import (
	"net/http"
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
)

const (
	USER      = "UserId"
	GAME      = "GameId"
)

func newDownloadHandler( c echo.Context) error {
	newDownload := models.Download{
		UserId: c.FormValue(USER),
		GameId: c.FormValue(GAME),
		CreationDate: time.Now(),
	}

	_db, getDbErr := db.NewDatabaseFromEnv()
	_, err := _db.AddDownload(newDownload)

	// error handling
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	if err != nil {
		log.Error("An error occurred while getting download records: %v", err)
		return c.JSON(http.StatusInternalServerError, "Database add download error")
	}

	message := "Download added"
	return c.JSON(http.StatusOK, message)
}

func init() {
	registerRoute(route{method: http.MethodPost, path: "/downloads", handler: newDownloadHandler})
}