package article

type Archive struct {
	Month    string
	Articles []*Article
}

func NewArchives(articles []*Article) []*Archive {
	archives := make([]*Archive, 0)
	var archive *Archive
	for _, article := range articles {
		month := article.CreatedAt[:7]
		if archive == nil || archive.Month != month {
			if archive != nil {
				archives = append(archives, archive)
			}
			archive = &Archive{}
			archive.Month = month
		}
		archive.Articles = append(archive.Articles, article)
	}
	if archive != nil {
		archives = append(archives, archive)
	}
	return archives
}
