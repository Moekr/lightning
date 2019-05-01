package article

import (
	"strings"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"

	"github.com/Moekr/gopkg/logs"
)

func (s *Store) buildIndex() {
	searcher := &engine.Engine{}
	searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: "dict/dictionary.txt",
	})
	for i, article := range s.i2a {
		searcher.IndexDocument(uint64(i)+1, types.DocumentIndexData{
			Content: article.Content,
			Labels:  article.Tags,
		}, false)
	}
	searcher.FlushIndex()
	logs.Info("[Search] indexed %d articles", searcher.NumDocumentsIndexed())
	s.searcher = searcher
}

func (s *Store) parseSearchQuery(query string) types.SearchRequest {
	var req types.SearchRequest
	var tag, key string
	if strings.HasPrefix(query, "tag:") {
		ss := strings.Fields(query)
		tag = ss[0]
		if len(ss) > 1 {
			key = strings.TrimSpace(strings.TrimPrefix(query, tag))
		}
	} else {
		key = query
	}
	if tag != "" {
		tags := strings.Split(strings.TrimPrefix(tag, "tag:"), ",")
		for _, tag := range tags {
			if tag != "" {
				req.Labels = append(req.Labels, tag)
			}
		}
	}
	if key != "" {
		req.Text = key
	}
	return req
}

func (s *Store) Search(query string, page int) *Page {
	req := s.parseSearchQuery(query)
	resp := s.searcher.Search(req)
	articles := make([]*Article, 0)
	for _, doc := range resp.Docs {
		articles = append(articles, s.i2a[doc.DocId-1])
	}
	return NewPage(articles, page)
}
