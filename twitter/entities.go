package twitter

// Entities represent metadata and context info parsed from Twitter components.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object
// TODO: symbols
type Entities struct {
	Hashtags     []HashtagEntity `json:"hashtags,omitempty"`
	Media        []MediaEntity   `json:"media,omitempty"`
	Urls         []URLEntity     `json:"urls,omitempty"`
	UserMentions []MentionEntity `json:"user_mentions,omitempty"`
}

// HashtagEntity represents a hashtag which has been parsed from text.
type HashtagEntity struct {
	Indices Indices `json:"indices,omitempty"`
	Text    string  `json:"text,omitempty"`
}

// URLEntity represents a URL which has been parsed from text.
type URLEntity struct {
	Indices     Indices `json:"indices,omitempty"`
	DisplayURL  string  `json:"display_url,omitempty"`
	ExpandedURL string  `json:"expanded_url,omitempty"`
	URL         string  `json:"url,omitempty"`
}

// MediaEntity represents media elements associated with a Tweet.
type MediaEntity struct {
	URLEntity
	ID                int64      `json:"id,omitempty"`
	IDStr             string     `json:"id_str,omitempty"`
	MediaURL          string     `json:"media_url,omitempty"`
	MediaURLHttps     string     `json:"media_url_https,omitempty"`
	SourceStatusID    int64      `json:"source_status_id,omitempty"`
	SourceStatusIDStr string     `json:"source_status_id_str,omitempty"`
	Type              string     `json:"type,omitempty"`
	Sizes             MediaSizes `json:"sizes,omitempty"`
	VideoInfo         VideoInfo  `json:"video_info,omitempty"`
}

// MentionEntity represents Twitter user mentions parsed from text.
type MentionEntity struct {
	Indices    Indices `json:"indices,omitempty"`
	ID         int64   `json:"id,omitempty"`
	IDStr      string  `json:"id_str,omitempty"`
	Name       string  `json:"name,omitempty"`
	ScreenName string  `json:"screen_name,omitempty"`
}

// UserEntities contain Entities parsed from User url and description fields.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object#mentions
type UserEntities struct {
	URL         Entities `json:"url,omitempty"`
	Description Entities `json:"description,omitempty"`
}

// ExtendedEntity contains media information.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/extended-entities-object
type ExtendedEntity struct {
	Media []MediaEntity `json:"media,omitempty"`
}

// Indices represent the start and end offsets within text.
type Indices [2]int

// Start returns the index at which an entity starts, inclusive.
func (i Indices) Start() int {
	return i[0]
}

// End returns the index at which an entity ends, exclusive.
func (i Indices) End() int {
	return i[1]
}

// MediaSizes contain the different size media that are available.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object#media-size
type MediaSizes struct {
	Thumb  MediaSize `json:"thumb,omitempty"`
	Large  MediaSize `json:"large,omitempty"`
	Medium MediaSize `json:"medium,omitempty"`
	Small  MediaSize `json:"small,omitempty"`
}

// MediaSize describes the height, width, and resizing method used.
type MediaSize struct {
	Width  int    `json:"w,omitempty"`
	Height int    `json:"h,omitempty"`
	Resize string `json:"resize,omitempty"`
}

// VideoInfo is available on video media objects.
type VideoInfo struct {
	AspectRatio    [2]int         `json:"aspect_ratio,omitempty"`
	DurationMillis int            `json:"duration_millis,omitempty"`
	Variants       []VideoVariant `json:"variants,omitempty"`
}

// VideoVariant describes one of the available video formats.
type VideoVariant struct {
	ContentType string `json:"content_type,omitempty"`
	Bitrate     int    `json:"bitrate,omitempty"`
	URL         string `json:"url,omitempty"`
}
