package deliv

import "alta/project2/features/book"

type Request struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Page      int    `json:"page" form:"page"`
}

func toCore(req Request) book.BookCore {

	var res = book.BookCore{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Page:      req.Page,
	}

	return res

}
