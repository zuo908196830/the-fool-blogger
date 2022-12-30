package request

type BlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type SearchBlogRequest struct {
	Limit  int `json:"limit" uri:"limit"`
	Offset int `json:"offset" uri:"offset"`
}

type BlogIdRequest struct {
	ID int `json:"id" uri:"id"`
}

type UpdateBlogRequest struct {
	ID      int    `json:"id" bind:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
