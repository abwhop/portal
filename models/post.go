package models

import "github.com/lib/pq"

type PostGQLRespond struct {
	Data struct {
		BlogPosts []*PostAPI `json:"blog_posts"`
	} `json:"data"`
}

type PostAPI struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	VoteNum     interface{}   `json:"vote_num"`
	BlogId      int           `json:"blog_id"`
	Text        string        `json:"text"`
	IsDraft     bool          `json:"is_draft"`
	CreateDate  int64         `json:"create_date"`
	PublishDate int64         `json:"publish_date"`
	Img         string        `json:"img"`
	Rights      []string      `json:"rights"`
	PostRights  []string      `json:"post_rights"`
	Files       []*FileAPI    `json:"files"`
	Author      *UserAPI      `json:"author"`
	Likes       *LikesAPI     `json:"likes"`
	Views       *ViewsAPI     `json:"views"`
	Comments    []*CommentAPI `json:"comments"`
	RepostBlog  *PostAPI      `json:"repost_blog"`
	RepostNews  *NewsAPI      `json:"repost_news"`
	Vote        interface{}   `json:"vote"`
}

/*type PostDB struct {
	BitrixId           int                  `json:"bitrix_id"`
	Type               string               `json:"type"`
	Title              string               `json:"title"`
	Description        string               `json:"description"`
	ImageUrl           string               `json:"image_url"`
	AuthorId           int                  `json:"author_id"`
	Comments           *ListOfCommentDB     `json:"comments"`
	Descriptions       *ListOfDescriptionDB `json:"descriptions"`
	PreviewDescription string               `json:"preview_description"`
	Likes              *LikesDB             `json:"likes"`
	DateUpdated        int                  `json:"date_updated"`
	RubricId           int                  `json:"rubric_id"`
	Published          bool                 `json:"published"`
	Rights             []string             `json:"rights"`
	PublishDate        int                  `json:"publish_date"`
	Views              *ViewsDB             `json:"views"`
	RepostBlogPostId   int                  `json:"repost_blog_post_id"`
	LogId              int                  `json:"log_id"`
	IsDeleted          bool                 `json:"is_deleted"`
	CanComment         bool                 `json:"can_comment"`
	Files              []*FileDB            `json:"files"`
	VoteId             int                  `json:"vote_id"`
	FormId             int                  `json:"form_id"`
}*/

type PostDB struct {
	Id               int                  `gorm:"column:id;primaryKey" json:"id"`
	Title            string               `gorm:"column:title;" json:"title"`
	Text             string               `gorm:"column:text;" json:"text"`
	CreatedDate      int64                `gorm:"column:createDate;" json:"create_date"`
	PublishDate      int64                `gorm:"column:publishDate;" json:"publish_date"`
	Img              string               `gorm:"column:img;" json:"img"`
	Rights           pq.StringArray       `gorm:"column:rights;type:varchar[]" json:"rights"`
	Files            *ListOfFileDB        `gorm:"column:files;" json:"files"`
	Author           *UserDB              `gorm:"column:author;" json:"author"`
	Likes            *LikesDB             `gorm:"column:likes;" json:"likes"`
	Views            *ViewsDB             `gorm:"column:views;" json:"views"`
	Comments         *ListOfCommentDB     `gorm:"column:comments;" json:"comments"`
	Descriptions     *ListOfDescriptionDB `gorm:"column:descriptions;" json:"descriptions"`
	BlogId           int                  `gorm:"column:blogId;" json:"blog_id"`
	RepostBlogPostId int                  `gorm:"column:repostBlogPostId;" json:"repost_blog_post_id"`
	RepostNewsId     int                  `gorm:"column:repostNewsId;" json:"repost_news_id"`
	IsDraft          bool                 `gorm:"column:isDraft;" json:"is_draft"`
	LastUpdateDate   int64                `gorm:"column:lastUpdateDate;" json:"last_update_date"`
	VoteNum          pq.Int64Array        `gorm:"column:voteNum;type:int[]" json:"vote_num"`
	IsDeleted        bool                 `gorm:"column:isDeleted;" json:"is_deleted"`
	PostRights       pq.StringArray       `gorm:"column:postRights;type:varchar[]" json:"post_rights"`
	FormId           pq.Int64Array        `gorm:"column:formId;type:int[]" json:"form_id"`
	CommentsCount    int                  `gorm:"column:commentsCount;" json:"comments_count"`
}

func (PostDB) TableName() string {
	return "public.bitrix_posts"
}
