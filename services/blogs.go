package services

import (
	"context"
	"fmt"
	"portal_sync/gql"
	"portal_sync/models"
	"portal_sync/query"
	"portal_sync/repository"
	"time"
)

func (srv *Service) LoadBlogs(limit int, page int) (int, error) {
	ctx := context.Background()
	var err error
	loafStart := time.Now()
	var respondModel models.BlogGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.BlogQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	blogDB, err := ConvertBlogs(respondModel.Data.Blogs)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetBlogs(blogDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Blogs), nil
}

func ConvertBlogPosts(postsAPI []*models.PostAPI) ([]*models.PostDB, error) {
	var postDB []*models.PostDB
	for _, blog := range postsAPI {
		blogDB, err := ConvertBlogPost(blog)
		if err != nil {
			continue
		}
		postDB = append(postDB, blogDB)
	}
	return postDB, nil
}

func ConvertBlogPost(postAPI *models.PostAPI) (*models.PostDB, error) {
	authorDB, err := ConvertUser(postAPI.Author)
	if err != nil {
		authorDB = nil
	}
	filesDB, err := ConvertFiles(postAPI.Files)
	if err != nil {
		filesDB = nil
	}
	likesDB, err := ConvertLikes(postAPI.Likes)
	if err != nil {
		filesDB = nil
	}

	viewsDB, err := ConvertViews(postAPI.Views)
	if err != nil {
		viewsDB = nil
	}

	commentCount, commentsDB, err := ConvertComments(postAPI.Comments)
	if err != nil {
		commentsDB = nil
	}
	formIds, convertedText, descriptionsDB, err := ConvertDescriptions(postAPI.Text)
	if err != nil {
		descriptionsDB = nil
	}

	repostBlogPostId := 0
	if postAPI.RepostBlog != nil {
		repostBlogPostId = postAPI.RepostBlog.BlogId
	}

	repostNewsId := 0
	if postAPI.RepostNews != nil {
		repostNewsId = postAPI.RepostNews.Id
	}

	return &models.PostDB{
		Id:               postAPI.Id,
		Text:             convertedText,
		Title:            postAPI.Title,
		CreatedDate:      postAPI.CreateDate * 1000,
		PublishDate:      postAPI.PublishDate * 1000,
		Img:              postAPI.Img,
		Rights:           postAPI.Rights,
		Files:            filesDB,
		Author:           authorDB,
		Likes:            likesDB,
		Views:            viewsDB,
		Comments:         commentsDB,
		Descriptions:     descriptionsDB,
		BlogId:           postAPI.BlogId,
		RepostBlogPostId: repostBlogPostId,
		RepostNewsId:     repostNewsId,
		IsDraft:          postAPI.IsDraft,
		LastUpdateDate:   postAPI.PublishDate,
		PostRights:       postAPI.PostRights,
		FormId:           formIds,
		CommentsCount:    commentCount,
	}, nil
}

func ConvertBlogs(blogsAPI []*models.BlogAPI) ([]*models.BlogDB, error) {
	var blogsDB []*models.BlogDB
	for _, blog := range blogsAPI {
		blogDB, err := ConvertBlog(blog)
		if err != nil {
			continue
		}
		blogsDB = append(blogsDB, blogDB)
	}
	return blogsDB, nil
}

func ConvertBlog(blogsAPI *models.BlogAPI) (*models.BlogDB, error) {
	authorDB, err := ConvertUser(blogsAPI.Author)
	if err != nil {
		authorDB = nil
	}
	subscribersDB, err := ConvertUsers(blogsAPI.Subscribers)
	if err != nil {
		subscribersDB = nil
	}
	return &models.BlogDB{
		BitrixId:         blogsAPI.Id,
		Name:             blogsAPI.Name,
		Description:      blogsAPI.Description,
		DateCreated:      blogsAPI.DateCreate,
		Author:           authorDB,
		Subscribers:      subscribersDB,
		SubscribersCount: len(blogsAPI.Subscribers),
	}, nil
}
