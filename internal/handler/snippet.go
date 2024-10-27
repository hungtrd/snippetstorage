package handler

import (
	"fmt"
	"net/http"
	"time"

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

	e.POST("/snippets", handler.CreateSnippet)
	e.GET("/snippets", echo.WrapHandler(http.HandlerFunc(handler.GetSnippets)))
	e.GET("/snippet/:id", echo.WrapHandler(http.HandlerFunc(handler.GetDetailSnippet)))
}

type CreateSnippetRequest struct {
	Title   string `form:"title" validate:"required"`
	Content string `form:"content" validate:"required"`
}

func (h *SnippetHandler) CreateSnippet(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(CreateSnippetRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	snippet := model.Snippet{
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := h.db.Client().Database("snippetstorage").Collection(model.SnippetCollection)

	result, err := collection.InsertOne(ctx, snippet)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.InsertedID)

	return c.JSON(200, snippet)
}

func (h *SnippetHandler) GetDetailSnippet(w http.ResponseWriter, r *http.Request) {
}

func (h *SnippetHandler) GetSnippets(w http.ResponseWriter, r *http.Request) {
}
