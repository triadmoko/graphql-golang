// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Comment struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Description string     `json:"description"`
	User        []*User    `json:"user"`
	Post        *Post      `json:"post"`
	Like        []*Like    `json:"like"`
}

type Like struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Type      string     `json:"type"`
	User      []*User    `json:"user"`
	Comment   []*Comment `json:"comment"`
	Post      []*Post    `json:"post"`
}

type NewComment struct {
	Title       string `json:"title"`
	Description int    `json:"description"`
}

type NewLike struct {
	Type string `json:"type"`
}

type NewLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewPost struct {
	Title       string `json:"title"`
	Description int    `json:"description"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewVerify struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type Post struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       string     `json:"title"`
	Description int        `json:"description"`
	User        *User      `json:"user"`
	Like        []*Like    `json:"like"`
}

type Token struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	Posts     []*Post    `json:"posts"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}
