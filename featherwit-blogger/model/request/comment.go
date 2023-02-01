package request

type AddComment struct {
	BlogId   int    `json:"blogId"`
	ParentId int    `json:"parentId"` //如果是二级评论的话，所属一级评论id，一级评论时为负数
	ReplyId  int    `json:"replyId"`  //如果是回复某二级评论的评论，二级评论id，否则为负数
	Content  string `json:"Content"`
}
