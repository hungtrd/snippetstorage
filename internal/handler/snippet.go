package handler

import (
	"fmt"
	"time"

	"github.com/a-h/templ"
	websnippet "github.com/hungtrd/snippetstorage/cmd/web/components/snippet"
	"github.com/hungtrd/snippetstorage/internal/database"
	"github.com/hungtrd/snippetstorage/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SnippetHandler struct {
	db database.Service
}

func NewSnippetHandler(e *echo.Echo, db database.Service) {
	handler := &SnippetHandler{
		db: db,
	}

	e.GET("/snippets/new", echo.WrapHandler(templ.Handler(websnippet.NewSnippet())))
	e.POST("/snippets", handler.CreateSnippet)
	e.GET("/snippets", handler.GetSnippets)
	e.GET("/snippet/:id", handler.GetDetailSnippet)
}

type CreateSnippetReq struct {
	Title    string `form:"title" validate:"required"`
	IsPublic bool   `form:"is_public" validate:"required"`
	Content  string `form:"content" validate:"required"`
}

func (h *SnippetHandler) CreateSnippet(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(CreateSnippetReq)
	if err := bindAndValidate(c, req); err != nil {
		return err
	}

	snippet := model.Snippet{
		Title:     req.Title,
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

func (h *SnippetHandler) GetDetailSnippet(c echo.Context) error {
	return nil
}

type GetSnipptesReq struct {
	Search string `query:"search"`
	Limit  int64  `query:"limit"`
	Offset int64  `query:"offset"`
}

func (h *SnippetHandler) GetSnippets(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(GetSnipptesReq)

	if err := c.Bind(req); err != nil {
		return err
	}

	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": req.Search, "$options": "i"}},
			{"content": bson.M{"$regex": req.Search, "$options": "i"}},
		},
	}

	findOptions := options.Find()
	if req.Limit != 0 {
		findOptions.SetLimit(req.Limit)
	}
	if req.Offset != 0 {
		findOptions.SetSkip(req.Offset)
	}
	collection := h.db.Client().Database("snippetstorage").Collection(model.SnippetCollection)
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var snippets []model.Snippet
	for cursor.Next(ctx) {
		var snippet model.Snippet
		if err := cursor.Decode(&snippet); err != nil {
			return err
		}
		fmt.Println(snippet)
		snippets = append(snippets, snippet)
	}

	return render(c, websnippet.ListSnippet(snippets))
}
