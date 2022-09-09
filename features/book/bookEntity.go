package book

import "time"

type BookCore struct {
	ID        uint
	Title     string
	Author    string
	Publisher string
	Page      int
	CreatedAt time.Time
	UserID    uint
	UpdatedAt time.Time
	DeletedAt time.Time
}

type DataInterface interface {
	SelectAll() (data []BookCore, err error)
	SelectById(param int) (data BookCore, err error)
	CreateData(data BookCore, token int) (int, error)
	UpdateData(param, token int, data BookCore) (int, error)
	DelData(param, token int) (int, error)
}

type ServiceInterface interface {
	GetAll() (data []BookCore, err error)
	GetById(param int) (data BookCore, err error)
	PostData(data BookCore, token int) (int, error)
	PutData(param, token int, data BookCore) (int, error)
	DeleteData(param, token int) (int, error)
}
