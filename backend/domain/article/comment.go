package article

import "github.com/mopeneko/qiita-clone/backend/domain/meta"

// Comment is a comment of a Article
type Comment struct {
	ID      string
	Meta    meta.CommentMeta
	Content string
}
