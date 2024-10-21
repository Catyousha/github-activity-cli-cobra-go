package domain

type Activity struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Type      string `json:"type" validate:"oneof='CommitCommentEvent' 'CreateEvent' 'DeleteEvent' 'ForkEvent' 'GollumEvent' 'IssueCommentEvent' 'IssuesEvent' 'MemberEvent' 'PublicEvent' 'PullRequestEvent' 'PullRequestReviewCommentEvent' 'PushEvent' 'ReleaseEvent' 'SponsorshipEvent' 'WatchEvent'"`
	Repo      struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload map[string]interface{} `json:"payload"`
}

type Payload interface {
	PushEventPayload
}

type PushEventPayload struct {
	Size int `json:"size"`
}

type CreateEventPayload struct {
	RefType string `json:"ref_type"`
}

type DeleteEventPayload struct {
	RefType string `json:"ref_type"`
}
