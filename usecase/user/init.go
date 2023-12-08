package user

type Repo struct {
	User
}

func NewUsecase(repoUser User) *Repo {
	return &Repo{
		User: repoUser,
	}
}
