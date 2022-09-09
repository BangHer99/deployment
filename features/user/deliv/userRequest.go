package deliv

import "alta/project2/features/user"

type UserReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req UserReq) user.UserCore {

	return user.UserCore{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

}
