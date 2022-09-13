package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-graphql-api/graph/generated"
	"go-graphql-api/graph/model"
	"time"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	Addpost := model.Post{
		Title:       input.Title,
		Content:     input.Content,
		Author:      *input.Author,
		Hero:        *input.Hero,
		PublishedAt: time.Now().Format("01-02-2006"),
		UpdatedAt:   time.Now().Format("01-02-2006"),
	}

	if err := r.Database.Create(&Addpost).Error; err != nil {
		fmt.Println(err)
		return nil, err

	}

	return &Addpost, nil
}

// UpdatePost is the resolver for the UpdatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, postID int, input *model.NewPost) (*model.Post, error) {
	Updatepost := model.Post{
		Title:     input.Title,
		Content:   input.Content,
		UpdatedAt: time.Now().Format("01-02-2006"),
	}

	if err := r.Database.Model(&model.Post{}).Where("id=?", postID).Updates(&Updatepost).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	Updatepost.ID = postID
	return &Updatepost, nil
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}

	GetPosts := r.Database.Model(&posts).Find(&posts)

	if GetPosts.Error != nil {
		fmt.Println(GetPosts.Error)
		return nil, GetPosts.Error
	}
	return posts, nil
}

// GetOnePost is the resolver for the GetOnePost field.
func (r *queryResolver) GetOnePost(ctx context.Context, id int) (*model.Post, error) {
	post := model.Post{}

	if err := r.Database.Find(&post, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
