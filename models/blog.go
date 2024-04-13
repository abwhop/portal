package models

type BlogGQLRespond struct {
	Data struct {
		Blogs []*BlogAPI `json:"blogs"`
	} `json:"data"`
}

type BlogAPI struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DateCreate  int        `json:"date_create"`
	Author      *UserAPI   `json:"author"`
	Subscribers []*UserAPI `json:"subscribers"`
}

type BlogDB struct {
	BitrixId         int            `gorm:"column:id;primaryKey" json:"id"`
	Name             string         `gorm:"column:name;" json:"name"`
	Description      string         `gorm:"column:description;" json:"description"`
	DateCreated      int            `gorm:"column:dateCreated;" json:"dateCreated"`
	Author           *UserDB        `gorm:"column:author;" json:"author"`
	Subscribers      *ListOfUsersDB `gorm:"column:subscribers;" json:"subscribers"`
	SubscribersCount int            `gorm:"column:subscribersCount;" json:"subscribersCount"`
}

func (BlogDB) TableName() string {
	return "public.bitrix_blogs"
}

/*type News struct {
	Id           uint          `msgpack:"id" json:"id"`
	Type         string        `msgpack:"type" json:"type"`
	PublishDate  int           `msgpack:"publish_date" json:"publish_date"`
	CreateDate   int           `msgpack:"create_date" json:"create_date"`
	ChangeDate   int           `msgpack:"date_updated" json:"change_date"`
	Name         string        `msgpack:"title" json:"name"`
	LogId        int           `msgpack:"log_id" json:"log_id"`
	CanComment   bool          `msgpack:"can_comment" json:"can_comment"`
	Descriptions []Description `msgpack:"descriptions" json:"descriptions"`
	Text         string        `msgpack:"description" json:"text"`
	Published    bool          `msgpack:"published" json:"published"`
	Rights       []string      `msgpack:"rights" json:"rights"`
	Img          string        `msgpack:"image_url" json:"img"`
	PreviewText  string        `msgpack:"preview_description" json:"preview_text"`
	XmlId        interface{}   `msgpack:"xml_id" json:"xml_id"`
	SliderFile   interface{}   `msgpack:"slider_file" json:"slider_file"`
	Rubric       struct {
		Name string      `msgpack:"name" json:"name"`
		Id   int         `msgpack:"id" json:"id"`
		Code interface{} `msgpack:"code" json:"code"`
	} `msgpack:"rubric" json:"rubric"`
	User UserShort `msgpack:"author" json:"author"`
	Likes  Likes     `msgpack:"likes" json:"likes"`
	Views  struct {
		Count int `msgpack:"count" json:"count"`
		Users []struct {
			Id             int    `msgpack:"id" json:"id"`
			Name           string `msgpack:"name" json:"name"`
			LastName       string `msgpack:"last_name" json:"last_name"`
			SecondName     string `msgpack:"second_name" json:"second_name"`
			Email          string `msgpack:"email" json:"email"`
			Position       string `msgpack:"position" json:"position"`
			PersonalNumber int    `msgpack:"personal_number" json:"personal_number"`
		} `msgpack:"users" json:"users"`
	} `msgpack:"views" json:"views"`
	Comments []Comment `msgpack:"comments" json:"comments"`
	Files    []File    `msgpack:"files" json:"files"`
	VoteId   []int     `msgpack:"vote_id" json:"vote_id"`
	FormId   []int     `msgpack:"form_id" json:"form_id"`
}*/
