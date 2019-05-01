package article

import (
	"html/template"
	"io/ioutil"
	"path"
	"sort"
	"strings"
	"sync/atomic"

	"gopkg.in/russross/blackfriday.v2"

	"github.com/Moekr/gopkg/logs"
	"github.com/huichen/wukong/engine"
)

type Store struct {
	dirPath  string
	i2a      []*Article
	n2a      map[string]*Article
	searcher *engine.Engine
}

var (
	s = &atomic.Value{}
)

func GetStore() *Store {
	if x := s.Load(); x != nil {
		if v, ok := x.(*Store); ok {
			return v
		}
	}
	return nil
}

func LoadArticles(dataPath string) {
	previous := GetStore()
	if previous != nil {
		defer previous.searcher.Close()
	}
	store := &Store{
		dirPath: dataPath,
	}
	store.loadArticles()
	store.buildIndex()
	for _, article := range store.n2a {
		article.HTML = template.HTML(blackfriday.Run([]byte(article.Content)))
	}
	s.Store(store)
}

func (s *Store) loadArticles() {
	if s.dirPath == "" {
		logs.Warn("[Store] dir path is empty")
		return
	}
	infos, err := ioutil.ReadDir(s.dirPath)
	if err != nil {
		logs.Error("[Store] read dir error: %s", err.Error())
	}
	s.i2a = make([]*Article, 0, len(infos)/2)
	s.n2a = make(map[string]*Article, len(infos)/2)
	for _, info := range infos {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			continue
		}
		name := strings.TrimSuffix(info.Name(), ".md")
		logs.Info("[Store] loading article: %s", name)
		content, err := ioutil.ReadFile(path.Join(s.dirPath, info.Name()))
		if err != nil {
			logs.Error("[Store] read content error: %s", err.Error())
			continue
		}
		metadata, err := ioutil.ReadFile(path.Join(s.dirPath, name+".meta"))
		if err != nil {
			logs.Error("[Store] read metadata error: %s", err.Error())
			continue
		}
		if article, err := NewArticle(name, content, metadata, info.ModTime()); err != nil {
			logs.Error("[Store] load error: %s", err.Error())
		} else {
			if article.IsHidden {
				continue
			}
			s.n2a[name] = article
			if article.IsPage {
				continue
			}
			s.i2a = append(s.i2a, article)
		}
	}
	sort.Slice(s.i2a, func(i, j int) bool {
		return s.i2a[i].CreatedAt > s.i2a[j].CreatedAt
	})
}

func (s *Store) GetPage(number int) *Page {
	return NewPage(s.i2a, number)
}

func (s *Store) Get(name string) *Article {
	return s.n2a[name]
}

func (s *Store) Archives() []*Archive {
	return NewArchives(s.i2a)
}
