package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/can3p/go-substack/client"
	"github.com/can3p/go-substack/markdown"
	"github.com/can3p/go-substack/types"
)

func main() {
	sessionID := flag.String("sid", "", "session id")
	substackName := flag.String("name", "", "substack name")
	substackID := flag.Int("substack-id", 0, "substack id")
	draftID := flag.Int("draft-id", 0, "draft id")
	title := flag.String("title", "", "draft title")
	md := flag.String("md", "", "markdown text to publish")

	flag.Parse()

	if *sessionID == "" || *substackName == "" || *md == "" || *title == "" || *substackID == 0 || *draftID == 0 {
		panic("all arguments are required")
	}

	c, err := client.NewClient(nil, *sessionID, *substackName)

	if err != nil {
		panic(err)
	}

	doc, err := markdown.ToDoc(*md)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	draft := &types.Draft{
		Audience:   "everyone",
		Type:       "newsletter",
		DraftTitle: *title,
		DraftBody:  types.DraftBody(*doc),
		DraftBylines: []types.DraftBylines{
			{ID: *substackID},
		},
	}

	respDraft, err := c.UpdateDraft(ctx, *draftID, draft)

	if err != nil {
		panic(err)
	}

	fmt.Printf("You can checkout updated draft on https://%s.substack.com/publish/post/%d\n", *substackName, respDraft.ID)
}
