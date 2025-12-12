package commands

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
)

func HandlerBrowse(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return errors.New("browse takes up to one integer argument: limit")
	}
	limit := 2
	if len(cmd.Args) == 1 {
		limitArg := cmd.Args[0]
		var err error
		limit, err = strconv.Atoi(limitArg)
		if err != nil {
			limit = 2
		}
	}

	limit32 := int32(limit)

	ctx := context.Background()

	posts, err := s.DbQueries.GetPostsByUser(ctx, database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  limit32,
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("PublishedAt: %v\nFeedName: %v\nTitle: %v\nDescription: %v\nUrl: %v\n\n\n", post.PublishedAt, post.FeedName, post.Title.String, post.Description.String, post.Url)
	}
	return nil
}
