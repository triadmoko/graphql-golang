package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"errors"
	"fmt"

	"github.com/triadmoko/grahpql-golang/graph"
	"github.com/triadmoko/grahpql-golang/graph/request"
	"github.com/triadmoko/grahpql-golang/helpers"
)

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input request.NewLogin) (*request.Token, error) {
	token, err := r.User.Login(ctx, input)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input request.NewUser) (*request.User, error) {
	user, err := r.User.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// VerifyUser is the resolver for the verifyUser field.
func (r *mutationResolver) VerifyUser(ctx context.Context, input request.NewVerify) (string, error) {
	status, err := r.User.VerifyEmail(ctx, input)
	if err != nil {
		return "", err
	}
	return status, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *request.UpdateUser) (*request.User, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	user, err := r.User.Update(ctx, sess.ID, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input request.NewPost) (*request.Post, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	preloads := helpers.GetPreloads(ctx)
	result, err := r.Post.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	for _, preload := range preloads {
		switch preload {
		case "user":
			user, err := r.User.Detail(ctx, sess.ID)
			if err != nil {
				return nil, err
			}
			result.User = user
		}
	}

	return result, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input request.NewPost) (*request.Post, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	preloads := helpers.GetPreloads(ctx)
	result, err := r.Post.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}
	for _, preload := range preloads {
		switch preload {
		case "user":
			user, err := r.User.Detail(ctx, sess.ID)
			if err != nil {
				return nil, err
			}
			result.User = user
		}
	}

	return result, nil
}

// DetailPost is the resolver for the detailPost field.
func (r *mutationResolver) DetailPost(ctx context.Context, id string) (*request.Post, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	preloads := helpers.GetPreloads(ctx)
	post, err := r.Post.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, preload := range preloads {
		switch preload {
		case "user":
			user, err := r.User.Detail(ctx, sess.ID)
			if err != nil {
				return nil, err
			}
			post.User = user
		case "comments":
			comments, err := r.Comment.CommentList(ctx, request.FilterComment{PostID: post.ID, Page: 1, PerPage: 100})
			if err != nil {
				return nil, err
			}
			post.Comments = comments
		case "total_like":
			count, err := r.LikeService.Count(ctx, post.ID)
			if err != nil {
				return nil, err
			}
			post.TotalLike = &count
		}
	}
	return post, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (string, error) {
	err := r.Post.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	return "success deleted post", nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input request.NewComment) (*request.Comment, error) {
	result, err := r.Comment.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input request.NewComment) (*request.Comment, error) {
	result, err := r.Comment.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (string, error) {
	err := r.Comment.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	return "success deleted comment", nil
}

// Like is the resolver for the Like field.
func (r *mutationResolver) Like(ctx context.Context, input request.NewLike) (*request.Like, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	input.UserID = &sess.ID
	response, err := r.LikeService.Like(ctx, input)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) ([]*request.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// PostList is the resolver for the postList field.
func (r *queryResolver) PostList(ctx context.Context, filter *request.FilterPost) (*request.PostList, error) {
	response, err := r.Post.List(ctx, *filter)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DetailPost is the resolver for the detailPost field.
func (r *queryResolver) DetailPost(ctx context.Context, id string) (*request.Post, error) {
	panic(fmt.Errorf("not implemented: DetailPost - detailPost"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UnLike(ctx context.Context, input request.NewLike) (*request.Like, error) {
	panic(fmt.Errorf("not implemented: UnLike - UnLike"))
}
func (r *queryResolver) CommentList(ctx context.Context, filter *request.FilterComment) (*request.CommentList, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
func (r *queryResolver) Comment(ctx context.Context) ([]*request.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}
func (r *queryResolver) Like(ctx context.Context) ([]*request.Like, error) {
	// sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	// if !oke {
	// 	return nil, errors.New("session invalid")
	// }
	return nil, nil
}
func (r *queryResolver) List(ctx context.Context, filter *request.FilterPost) ([]*request.PostList, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}
