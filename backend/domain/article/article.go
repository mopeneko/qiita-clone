package article

import "github.com/mopeneko/qiita-clone/backend/domain/meta"

// Article is an article of the site
type Article struct {
	ID       string
	Meta     meta.ArticleMeta
	Content  string
	Comments []Comment
}
