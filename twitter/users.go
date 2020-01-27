package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// User represents a Twitter User.
// https://dev.twitter.com/overview/api/users
type User struct {
	ContributorsEnabled            bool          `json:"contributors_enabled,omitempty"`
	CreatedAt                      string        `json:"created_at,omitempty"`
	DefaultProfile                 bool          `json:"default_profile,omitempty"`
	DefaultProfileImage            bool          `json:"default_profile_image,omitempty"`
	Description                    string        `json:"description,omitempty"`
	Email                          string        `json:"email,omitempty"`
	Entities                       *UserEntities `json:"entities,omitempty"`
	FavouritesCount                int           `json:"favourites_count,omitempty"`
	FollowRequestSent              bool          `json:"follow_request_sent,omitempty"`
	Following                      bool          `json:"following,omitempty"`
	FollowersCount                 int           `json:"followers_count,omitempty"`
	FriendsCount                   int           `json:"friends_count,omitempty"`
	GeoEnabled                     bool          `json:"geo_enabled,omitempty"`
	ID                             int64         `json:"id,omitempty"`
	IDStr                          string        `json:"id_str,omitempty"`
	IsTranslator                   bool          `json:"is_translator,omitempty"`
	Lang                           string        `json:"lang,omitempty"`
	ListedCount                    int           `json:"listed_count,omitempty"`
	Location                       string        `json:"location,omitempty"`
	Name                           string        `json:"name,omitempty"`
	Notifications                  bool          `json:"notifications,omitempty"`
	ProfileBackgroundColor         string        `json:"profile_background_color,omitempty"`
	ProfileBackgroundImageURL      string        `json:"profile_background_image_url,omitempty"`
	ProfileBackgroundImageURLHttps string        `json:"profile_background_image_url_https,omitempty"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile,omitempty"`
	ProfileBannerURL               string        `json:"profile_banner_url,omitempty"`
	ProfileImageURL                string        `json:"profile_image_url,omitempty"`
	ProfileImageURLHttps           string        `json:"profile_image_url_https,omitempty"`
	ProfileLinkColor               string        `json:"profile_link_color,omitempty"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color,omitempty"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color,omitempty"`
	ProfileTextColor               string        `json:"profile_text_color,omitempty"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image,omitempty"`
	Protected                      bool          `json:"protected,omitempty"`
	ScreenName                     string        `json:"screen_name,omitempty"`
	ShowAllInlineMedia             bool          `json:"show_all_inline_media,omitempty"`
	Status                         *Tweet        `json:"status,omitempty"`
	StatusesCount                  int           `json:"statuses_count,omitempty"`
	Timezone                       string        `json:"time_zone,omitempty"`
	URL                            string        `json:"url,omitempty"`
	UtcOffset                      int           `json:"utc_offset,omitempty"`
	Verified                       bool          `json:"verified,omitempty"`
	WithheldInCountries            []string      `json:"withheld_in_countries,omitempty"`
	WithholdScope                  string        `json:"withheld_scope,omitempty"`
}

// UserService provides methods for accessing Twitter user API endpoints.
type UserService struct {
	sling *sling.Sling
}

// newUserService returns a new UserService.
func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("users/"),
	}
}

// UserShowParams are the parameters for UserService.Show.
type UserShowParams struct {
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Show returns the requested User.
// https://dev.twitter.com/rest/reference/get/users/show
func (s *UserService) Show(params *UserShowParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// UserLookupParams are the parameters for UserService.Lookup.
type UserLookupParams struct {
	UserID          []int64  `url:"user_id,omitempty,comma"`
	ScreenName      []string `url:"screen_name,omitempty,comma"`
	IncludeEntities *bool    `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Lookup returns the requested Users as a slice.
// https://dev.twitter.com/rest/reference/get/users/lookup
func (s *UserService) Lookup(params *UserLookupParams) ([]User, *http.Response, error) {
	users := new([]User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// UserSearchParams are the parameters for UserService.Search.
type UserSearchParams struct {
	Query           string `url:"q,omitempty"`
	Page            int    `url:"page,omitempty"` // 1-based page number
	Count           int    `url:"count,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Search queries public user accounts.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/get/users/search
func (s *UserService) Search(query string, params *UserSearchParams) ([]User, *http.Response, error) {
	if params == nil {
		params = &UserSearchParams{}
	}
	params.Query = query
	users := new([]User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("search.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}
