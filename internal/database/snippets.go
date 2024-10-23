package database

import (
	"context"
	"fmt"

	"github.com/hungtrd/snippetstorage/internal/model"
)

func (s *service) InsertSnippet(ctx context.Context, snippet model.Snippet) (model.Snippet, error) {
	collection := s.db.Database(database).Collection(model.SnippetCollection)
	result, err := collection.InsertOne(ctx, snippet)
	if err != nil {
		return model.Snippet{}, err
	}
	fmt.Println(result.InsertedID)

	return snippet, nil
}
