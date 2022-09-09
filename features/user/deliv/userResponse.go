package deliv

import "alta/project2/features/user"

type Respon struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Book  []user.BookCore
}

func toResponId(data user.UserCore) Respon {

	return Respon{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Book:  data.Book,
	}

}

func toResponList(data []user.UserCore) []Respon {

	var respon []Respon
	for _, v := range data {
		respon = append(respon, Respon{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
			Book:  v.Book,
		})
	}

	return respon
}
