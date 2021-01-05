package user

import "github.com/mopeneko/qiita-clone/backend/domain/meta"

// User is an user of the site
type User struct {
	ID   string
	Meta meta.UserMeta
	Auth Auth
}
