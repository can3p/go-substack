package types

import (
	"encoding/json"
	"time"
)

type ImageRequest struct {
	Image string `json:"image,omitempty"`
}

type ImageResponse struct {
	ID          string `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Bytes       int64  `json:"bytes,omitempty"`
}

type DraftBylines struct {
	ID      int  `json:"id"`
	IsGuest bool `json:"is_guest"`
}

type DraftBody Node

func (db DraftBody) MarshalJSON() ([]byte, error) {
	node := (Node)(db)
	b, err := json.Marshal(node)

	if err != nil {
		return nil, err
	}

	// encode json into a string to put inside of json
	return json.Marshal(string(b))
}

func (db *DraftBody) UnmarshalJSON(b []byte) error {
	var unwrapped string

	if err := json.Unmarshal(b, &unwrapped); err != nil {
		return err
	}

	var n Node

	if err := json.Unmarshal([]byte(unwrapped), &n); err != nil {
		return err
	}

	converted := DraftBody(n)

	*db = converted

	return nil
}

type DraftRequest struct {
	Draft
}

type DraftResponse struct {
	Draft
}

type Draft struct {
	Type                        string         `json:"type,omitempty"`
	DraftTitle                  string         `json:"draft_title,omitempty"`
	DraftSubtitle               string         `json:"draft_subtitle,omitempty"`
	DraftPodcastURL             string         `json:"draft_podcast_url,omitempty"`
	Audience                    string         `json:"audience,omitempty"`
	SectionChosen               bool           `json:"section_chosen,omitempty"`
	DraftPodcastDuration        int64          `json:"draft_podcast_duration,omitempty"`
	DraftVideoUploadID          string         `json:"draft_video_upload_id,omitempty"`
	DraftPodcastUploadID        string         `json:"draft_podcast_upload_id,omitempty"`
	DraftPodcastPreviewUploadID string         `json:"draft_podcast_preview_upload_id,omitempty"`
	DraftVoiceoverUploadID      string         `json:"draft_voiceover_upload_id,omitempty"`
	PublicationID               int            `json:"publication_id,omitempty"`
	WordCount                   int            `json:"word_count,omitempty"`
	WriteCommentPermissions     string         `json:"write_comment_permissions,omitempty"`
	ShouldSendEmail             bool           `json:"should_send_email,omitempty"`
	ShowGuestBios               bool           `json:"show_guest_bios,omitempty"`
	CoverImage                  string         `json:"cover_image,omitempty"`
	Description                 string         `json:"description,omitempty"`
	SearchEngineDescription     string         `json:"search_engine_description,omitempty"`
	SearchEngineTitle           string         `json:"search_engine_title,omitempty"`
	Slug                        string         `json:"slug,omitempty"`
	SocialTitle                 string         `json:"social_title,omitempty"`
	PodcastDescription          string         `json:"podcast_description,omitempty"`
	FreeUnlockRequired          bool           `json:"free_unlock_required,omitempty"`
	SyndicateVoiceoverToRss     bool           `json:"syndicate_voiceover_to_rss,omitempty"`
	AudienceBeforeArchived      string         `json:"audience_before_archived,omitempty"`
	ExemptFromArchivePaywall    bool           `json:"exempt_from_archive_paywall,omitempty"`
	Explicit                    string         `json:"explicit,omitempty"`
	DefaultCommentSort          string         `json:"default_comment_sort,omitempty"`
	DraftBody                   DraftBody      `json:"draft_body,omitempty"`
	DraftSectionID              string         `json:"draft_section_id,omitempty"`
	ShouldSendFreePreview       bool           `json:"should_send_free_preview,omitempty"`
	Body                        string         `json:"body,omitempty"`
	DraftCreatedAt              *time.Time     `json:"draft_created_at,omitempty"`
	DraftUpdatedAt              *time.Time     `json:"draft_updated_at,omitempty"`
	EmailSentAt                 string         `json:"email_sent_at,omitempty"`
	ID                          int            `json:"id,omitempty"`
	IsPublished                 bool           `json:"is_published,omitempty"`
	PodcastDuration             string         `json:"podcast_duration,omitempty"`
	PodcastURL                  string         `json:"podcast_url,omitempty"`
	VideoUploadID               string         `json:"video_upload_id,omitempty"`
	PodcastUploadID             string         `json:"podcast_upload_id,omitempty"`
	PodcastPreviewUploadID      string         `json:"podcast_preview_upload_id,omitempty"`
	VoiceoverUploadID           string         `json:"voiceover_upload_id,omitempty"`
	PostDate                    string         `json:"post_date,omitempty"`
	ReplyToPostID               string         `json:"reply_to_post_id,omitempty"`
	SectionID                   string         `json:"section_id,omitempty"`
	SubscriberSetID             string         `json:"subscriber_set_id,omitempty"`
	Subtitle                    string         `json:"subtitle,omitempty"`
	Title                       string         `json:"title,omitempty"`
	UUID                        string         `json:"uuid,omitempty"`
	EditorV2                    bool           `json:"editor_v2,omitempty"`
	DraftVideoUpload            string         `json:"draftVideoUpload,omitempty"`
	DraftPodcastUpload          string         `json:"draftPodcastUpload,omitempty"`
	PodcastEpisodeNumber        string         `json:"podcast_episode_number,omitempty"`
	PodcastSeasonNumber         string         `json:"podcast_season_number,omitempty"`
	PodcastEpisodeType          string         `json:"podcast_episode_type,omitempty"`
	ShouldSyndicateToOtherFeed  string         `json:"should_syndicate_to_other_feed,omitempty"`
	SyndicateToSectionID        string         `json:"syndicate_to_section_id,omitempty"`
	HideFromFeed                bool           `json:"hide_from_feed,omitempty"`
	DraftBylines                []DraftBylines `json:"draft_bylines,omitempty"`
}
