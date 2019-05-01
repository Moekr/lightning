package article

import "github.com/Moekr/lightning/util/algo"

type Page struct {
	Articles []*Article

	Total   int
	Count   int
	Number  int
	IsFirst bool
	IsLast  bool
}

const (
	PageSize = 10
)

func NewPage(articles []*Article, number int) *Page {
	page := &Page{
		Articles: make([]*Article, 0, 10),
	}
	page.Total = len(articles)
	page.Count = (page.Total + PageSize - 1) / PageSize
	page.Number = algo.Max(algo.Min(number, page.Count), 1)
	page.IsFirst = page.Number == 1
	page.IsLast = page.Number == page.Count
	for i := (page.Number - 1) * PageSize; i < len(articles) && len(page.Articles) < PageSize; i++ {
		page.Articles = append(page.Articles, articles[i])
	}
	return page
}
