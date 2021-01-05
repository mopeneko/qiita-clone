package user

import "github.com/mopeneko/qiita-clone/backend/domain/meta"

// Auth is an auth info of User
type Auth struct {
	ID       string
	Meta     meta.AuthMeta
	Email    string
	Password string
}
