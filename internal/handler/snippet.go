package handler

import (
	"fmt"
	"net/http"

	"github.com/hungtrd/snippetstorage/internal/database"
	"github.com/hungtrd/snippetstorage/internal/model"
	"github.com/labstack/echo/v4"
)

type SnippetHandler struct {
	db database.Service
}

func NewSnippetHandler(e *echo.Echo, db database.Service) {
	handler := &SnippetHandler{
		db: db,
	}

	e.POST("/snippets", handler.CreateSnippetHandler)
	e.GET("/snippets", echo.WrapHandler(http.HandlerFunc(handler.GetSnippetsHandler)))
	e.GET("/snippet/:id", echo.WrapHandler(http.HandlerFunc(handler.GetDetailSnippetHandler)))
}

func (h *SnippetHandler) CreateSnippetHandler(c echo.Context) error {
	ctx := c.Request().Context()

	snippet := model.Snippet{}

	collection := h.db.Client().Database("snippetstorage").Collection(model.SnippetCollection)

	result, err := collection.InsertOne(ctx, snippet)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.InsertedID)

	return c.JSON(200, snippet)
}

func (h *SnippetHandler) GetDetailSnippetHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *SnippetHandler) GetSnippetsHandler(w http.ResponseWriter, r *http.Request) {
}
