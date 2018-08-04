package ig

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type CountWrapper struct {
	Count int `json:"count"`
}

type IdWrapper struct {
	ID string `json:"id"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"has_next_page"`
	EndCursor       string `json:"end_cursor"`
	StartCursor     string `json:"start_cursor"`
	HasPreviousPage bool   `json:"has_previous_page"`
}

type NodeWrapper struct {
	Node Node `json:"node"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	ProfilePicURL string `json:"profile_pic_url"`
}

type PostResponse struct {
	ID                   string     `json:"id"`
	ShortCode            string     `json:"shortcode"`
	MediaPreview         string     `json:"media_preview"`
	DisplayURL           string     `json:"display_url"`
	ShouldLogClientEvent bool       `json:"should_log_client_event"`
	IsVideo              bool       `json:"is_video"`
	TrackingToken        string     `json:"tracking_token"`
	TakenAtTimestamp     int        `json:"taken_at_timestamp"`
	CaptionIsEdited      int        `json:"caption_is_edited"`
	Dimensions           Dimensions `json:"dimensions"`
	User                 User       `json:"owner"`
	DisplayResources     []struct {
		Src          string `json:"src"`
		ConfigWidth  int    `json:"config_width"`
		ConfigHeight int    `json:"config_height"`
	} `json:"display_resources"`
	EdgeMediaToTaggedUser struct {
		Edges []struct {
			Node struct {
				User struct {
					id       string `json:"id"`
					username string `json:"username"`
				} `json:"user"`
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_tagged_user"`
	EdgeMediaToCaption struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
}

type TagResponse struct {
	Name           string `json:"name"`
	IsTopMediaOnly bool   `json:"is_top_media_only"`
	ProfilePicURL  string `json:"profile_pic_url"`
	Medias         struct {
		Count int           `json:"count"`
		Nodes []NodeWrapper `json:"edges"`
	} `json:"edge_hashtag_to_media"`
	TopPosts struct {
		Nodes []NodeWrapper `json:"edges"`
	} `json:"edge_hashtag_to_top_posts"`
}

type LoginResponse struct {
	Status        string `json:"status"`
	Authenticated bool   `json:"authenticated"`
	User          bool   `json:"user"`
	UserID        string `json:"userId"`
}
