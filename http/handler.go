package http

import (
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/Moekr/lightning/article"
	"github.com/Moekr/lightning/util/version"
)

func handleIndex(ctx *macaron.Context) {
	page := article.GetStore().GetPage(ctx.QueryInt("p"))
	ctx.Data["title"] = "首页"
	ctx.Data["articles"] = page.Articles
	ctx.Data["page"] = page
	ctx.Data["version"] = version.Code
	ctx.HTML(http.StatusOK, "index")
}

func handlePage(ctx *macaron.Context) {
	handlePost(ctx, true, false)
}

func handleArticle(ctx *macaron.Context) {
	handlePost(ctx, false, false)
}

func handleMarkdown(ctx *macaron.Context) {
	handlePost(ctx, false, true)
}

func handlePost(ctx *macaron.Context, isPage, isMarkdown bool) {
	_article := article.GetStore().Get(ctx.Params(":name"))
	if _article == nil || _article.IsPage != isPage {
		ctx.Status(http.StatusNotFound)
		return
	}
	if isMarkdown {
		ctx.Header().Set("Content-Type", "text/markdown; charset=UTF-8")
		ctx.PlainText(http.StatusOK, []byte(_article.Content))
		return
	}
	ctx.Data["title"] = _article.Title
	ctx.Data["article"] = _article
	ctx.Data["version"] = version.Code
	ctx.HTML(http.StatusOK, "article")
}

func handleArchive(ctx *macaron.Context) {
	ctx.Data["title"] = "文章归档"
	ctx.Data["archives"] = article.GetStore().Archives()
	ctx.Data["version"] = version.Code
	ctx.HTML(http.StatusOK, "archive")
}

func handleSearch(ctx *macaron.Context) {
	query, p := ctx.Query("q"), ctx.QueryInt("p")
	page := article.GetStore().Search(query, p)
	ctx.Data["title"] = "搜索"
	ctx.Data["articles"] = page.Articles
	ctx.Data["page"] = page
	ctx.Data["query"] = query
	ctx.Data["version"] = version.Code
	ctx.HTML(http.StatusOK, "search")
}

func handleNotFound(ctx *macaron.Context) {
	handleError(ctx, http.StatusNotFound, "Not Found")
}

func handleInternalServerError(ctx *macaron.Context, err error) {
	handleError(ctx, http.StatusInternalServerError, err.Error())
}

func handleError(ctx *macaron.Context, code int, message string) {
	ctx.Data["title"] = "出错了"
	ctx.Data["code"] = code
	ctx.Data["message"] = message
	ctx.Data["version"] = version.Code
	ctx.HTML(code, "error")
}
