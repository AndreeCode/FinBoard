package domains

type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ParentId    *string `json:"parent_id"`
	UserId      *string `json:"user_id"`
}
type UpdateCategoryRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
