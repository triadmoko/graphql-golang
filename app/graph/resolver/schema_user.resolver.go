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
	panic(fmt.Errorf("not implemented: UpdatePost - updatePost"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, id string, input request.NewComment) (*request.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, input request.NewComment) (*request.Comment, error) {
	panic(fmt.Errorf("not implemented: UpdateComment - updateComment"))
}

// Like is the resolver for the Like field.
func (r *mutationResolver) Like(ctx context.Context, id string, input request.NewLike) (*request.Like, error) {
	sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	if !oke {
		return nil, errors.New("session invalid")
	}
	fmt.Println(helpers.PrettyPrint(sess))
	return &request.Like{}, nil
}

// UnLike is the resolver for the UnLike field.
func (r *mutationResolver) UnLike(ctx context.Context, id string, input request.NewLike) (*request.Like, error) {
	panic(fmt.Errorf("not implemented: UnLike - UnLike"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) ([]*request.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context) ([]*request.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}

// Like is the resolver for the like field.
func (r *queryResolver) Like(ctx context.Context) ([]*request.Like, error) {
	// sess, oke := ctx.Value("sess").(*helpers.MetaToken)
	// if !oke {
	// 	return nil, errors.New("session invalid")
	// }
	return nil, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
