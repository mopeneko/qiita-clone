package meta

// UserMeta is meta info of user.User
type UserMeta struct {
	Meta
	Name string
	Bio  string
	URLs []string
}
