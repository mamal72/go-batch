package batch

// APIClient is the main struct for handling all batch API calls
type APIClient struct {
	RESTAPIKey  string
	BatchAPIKey string
}

// TransactionalPushResponse is the response type of transactional push notifications
// See https://batch.com/doc/api/transactional.html#_responses
type TransactionalPushResponse struct {
	Token string `json:"token"`
}

// TransactionalPushRecipient is a sturct containing
// tokens, custom ids and install ids of recipients
// See https://batch.com/doc/api/transactional.html#_post-data
type TransactionalPushRecipient struct {
	Tokens     []string `json:"tokens,omitempty"`
	CustomIDs  []string `json:"custom_ids,omitempty"`
	InstallIDs []string `json:"install_ids,omitempty"`
}

// TransactionalPushMessage is a struct containing message title and body
// See https://batch.com/doc/api/transactional.html#_post-data
type TransactionalPushMessage struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body"`
}

// TransactionalPushPriority is a string aliased type for
// transtional push notifications priority
// See https://batch.com/doc/api/transactional.html#_post-data
type TransactionalPushPriority string

// HighPriority is high transactional push priority constant
const HighPriority TransactionalPushPriority = "high"

// NormalPriority is normal transactional push priority constant
const NormalPriority TransactionalPushPriority = "normal"

// TransactionalPushPayload TBA
// See https://batch.com/doc/api/transactional.html#_post-data
type TransactionalPushPayload struct {
	GroupID    string                     `json:"group_id"`
	Labels     []string                   `json:"labels,omitempty"`
	Priority   TransactionalPushPriority  `json:"priority,omitempty"`
	Recipients TransactionalPushRecipient `json:"recipients"`
	Message    TransactionalPushMessage   `json:"message"`
	Deeplink   string                     `json:"deeplink,omitempty"`
}
