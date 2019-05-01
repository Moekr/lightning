package article

import (
	"html/template"
	"time"

	"gopkg.in/yaml.v2"
)

type Article struct {
	Name       string        `yaml:"name"`
	Title      string        `yaml:"title"`
	Summary    string        `yaml:"summary"`
	Content    string        `yaml:"-"`
	HTML       template.HTML `yaml:"-"`
	Tags       []string      `yaml:"tags"`
	CreatedAt  string        `yaml:"created-at"`
	ModifiedAt string        `yaml:"-"`
	IsPage     bool          `yaml:"is-page"`
	IsHidden   bool          `yaml:"is-hidden"`
}

func NewArticle(name string, content, metadata []byte, modifiedAt time.Time) (*Article, error) {
	article := &Article{
		Name: name,
	}
	if err := yaml.Unmarshal(metadata, article); err != nil {
		return nil, err
	}
	article.Content = string(content)
	article.ModifiedAt = modifiedAt.Format("2006-01-02")
	return article, nil
}
