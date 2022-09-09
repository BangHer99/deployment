package deliv

import (
	"alta/project2/features/book"
)

type Respon struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Page      int    `json:"page" form:"page"`
}

func toRespon(data book.BookCore) Respon {

	var res = Respon{
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		Page:      data.Page,
	}

	return res

}

func toResponList(data []book.BookCore) []Respon {

	var res []Respon
	for _, v := range data {
		res = append(res, toRespon(v))
	}

	return res

}
