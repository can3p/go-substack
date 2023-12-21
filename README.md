# Go-substack - push markdown straight to substack

This package provides a converter to transform markdown
into substack document format and a client to push the document to substack to
create and update drafts, you can check the [notes](https://github.com/can3p/substack-api-notes)

## Example Usage

See [test](test) folder for example scripts.

Create a draft:

```
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
	title := flag.String("title", "", "draft title")
	md := flag.String("md", "", "markdown text to publish")

	flag.Parse()

	if *sessionID == "" || *substackName == "" || *md == "" || *title == "" || *substackID == 0 {
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

	respDraft, err := c.CreateNewDraft(ctx, draft)

	if err != nil {
		panic(err)
	}

	fmt.Printf("You can checkout new draft on https://%s.substack.com/publish/post/%d\n", *substackName, respDraft.ID)
}
```

You can run it like this:

```
$ go run create_draft.go -name substack_name -sid take_from_cookie -title "test draft from the client 333" -md "this is the first paragraph

And this is the second one111" -substack-id 1234
You can checkout new draft on https://substack_name.substack.com/publish/post/1456
```

Update an existing draft:

```
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
```

Will be run like this:

```
$ go run update_draft.go -name substack_name -sid take_from_cookie -title "test draft from the client" -md "this is the first paragraph

And this is the second one111234" -substack-id 12345 -draft-id 456
You can checkout updated draft on https://substack_name.substack.com/publish/post/456

```

where:

* `sid` is a `substack.sid` cookie that you can extract with dev tools
* `substack-id` is a substack internal id probable, you can get it from one of the requests with dev tools

## FAQ

### Can I use it?

Only on your own risk. API is not public and the package can break any time

### Should I use it?

Only if you're adventurous

### Why should I hunt for the values with dev tools?

API is not public and there is no way to use authentenication like OAuth or anything like this. If you know a better way, please let me know!

### Is all markdown supported?

Not at all, you could check [tests](markdown/convert_test.go) to see what's allowed. The good news is that the package should fail conversion in case of unknown markup instead of trying to push garbage

### Are substack specific elements supported?

They can be, since it's clear how to add the to the substack document format, but not implemented yet. The way to go would be to extend markdown a bit for that

## License

MIT
