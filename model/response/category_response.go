package response

import "time"

type CategoryCreateResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryGetResponse struct {
	ID        int            `json:"id"`
	Type      string         `json:"type"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	Tasks     []CategoryTask `json:"Tasks"`
}

type CategoryTask struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryPatchResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryDeleteResponse struct {
	Message string `json:"message"`
}
