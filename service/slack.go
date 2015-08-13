package service

type SlackTopic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

type SlackPurpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

type SlackChannel struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	IsChannel  bool         `json:"is_channel"`
	Created    int          `json:"created"`
	Creator    string       `json:"creator"`
	IsArchived bool         `json:"is_archived"`
	IsGeneral  bool         `json:"is_general"`
	IsMember   bool         `json:"is_member"`
	Members    []string     `json:"members"`
	Topic      SlackTopic   `json:"topic"`
	Purpose    SlackPurpose `json:"purpose"`
	NumMembers int          `json:"num_members"`
}

type SlackChannels struct {
	Ok       bool           `json:"ok"`
	Channels []SlackChannel `json:"channels"`
}

type SlackIm struct {
	ID            string `json:"id"`
	IsIm          bool   `json:"is_im"`
	User          string `json:"user"`
	Created       int    `json:"created"`
	IsUserDeleted bool   `json:"is_user_deleted"`
}

type SlackMessages struct {
	Ok       bool    `json:"ok"`
	Messages []Slack `json:"messages"`
	Has_more bool    `json:"has_more"`
}

/*
type Message struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Attachments []SlackAttachment `json:"attachments,omitempty"`
	Ts string `json:"ts"`
}
*/

type SkackEdited struct {
	User string `json:"user"`
	Ts   string `json:"ts"`
}

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type SlackReaction struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Users []string `json:"users"`
}

type SlackAttachment struct {
	Title       string       `json:"title,omitempty"`
	TitleLink   string       `json:"title_link,omitempty"`
	Text        string       `json:"text,omitempty"`
	Fallback    string       `json:"fallback,omitempty"`
	Color       string       `json:"color,omitempty"`
	Pretext     string       `json:"pretext,omitempty"`
	AuthorName  string       `json:"author_name,omitempty"`
	AuthorLink  string       `json:"author_link,omitempty"`
	AuthorIcon  string       `json:"author_icon,omitempty"`
	Fields      []SlackField `json:"fields,omitempty"`
	ImageWidth  int          `json:"image_width,omitempty"`
	ImageHeight int          `json:"image_height,omitempty"`
	ImageBytes  int          `json:"image_bytes,omitempty"`
	FromURL     string       `json:"from_url"`
	ImageURL    string       `json:"image_url,omitempty"`
	ThumbURL    string       `json:"thumb_url,omitempty"`
}

type Slack struct {
	Id          int32  `json:"id"`
	Created     int32  `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type" binding:"required"`
	Channel     string `json:"channel"`
	User        string `json:"user"`
	Text        string `json:"text"`
	Ts          string `json:"ts"`
	Subtype     string `json:"subtype,omitempty"`
	//Edited SkackEdited `json:"edited,omitempty"`
	Hidden    bool   `json:"hidden,omitempty"`
	DeletedTs string `json:"deleted_ts,omitempty"`
	EventTs   string `json:"event_ts,omitempty"`
	IsStarred bool   `json:"is_starred,omitempty"`
	//Attachments []SlackAttachment `json:"attachments,omitempty"`
	//PinnedTo  []string `json:"pinned_to,omitempty"`
	//Reactions []SlackReaction `json:"reactions,omitempty"`
}

const (
	SlackStatus string = "todo"
	DoingStatus string = "doing"
	DoneStatus  string = "done"
)
