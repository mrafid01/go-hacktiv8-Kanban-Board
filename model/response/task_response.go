package response

import "time"

type TaskCreateResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskGetResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        TaskUser  `json:"User"`
}

type TaskUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type TaskUpdateResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskPatchStatusResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskPatchCategoryResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskDeleteResponse struct {
	Message string `json:"message"`
}
