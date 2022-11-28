package input

type TaskCreateInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

type TaskUpdateInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPatchStatusInput struct {
	Status *bool `json:"status" binding:"required"`
}

type TaskPatchCategoryInput struct {
	CategoryID int `json:"category_id" binding:"required"`
}

type TaskIdUri struct {
	ID int `uri:"taskID" binding:"required"`
}
