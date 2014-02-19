package controllers

import (
	"fmt"
	"html/template"
)

const (
	PagesPerView = 11 //最大页码数
	ItemsPerPage = 10 //每页记录数
)

type Pagination struct {
	page  int //当前页码
	rows  int //记录总数
	url   string
	pages int //总页数
}

func NewPagination(page int, rows int, url string) *Pagination {
	return &Pagination{
		page: page,
		rows: rows,
		url:  url,
	}
}

func (p *Pagination) Html() template.HTML {
	html := ""
	p.pages = p.rows / ItemsPerPage
	if p.pages*ItemsPerPage < p.rows {
		p.pages += 1
	}
	if p.pages == 1 {
		return template.HTML(html)
	}

	page := p.page
	page -= PagesPerView / 2
	if page < 0 {
		page = 0
	}

	count := page + PagesPerView
	if count > p.pages {
		count = p.pages
	}

	pageNum := 0
	for ; page < count; page++ {
		pageNum = page + 1
		if page != p.page {
			html += fmt.Sprintf(`<li><a href="%s%d">%d</a></li>`, p.url, pageNum, pageNum)
		} else {
			html += fmt.Sprintf(`<li class="active"><a href="#">%d</a></li>`, pageNum)
		}
	}

	return template.HTML(html)
}
