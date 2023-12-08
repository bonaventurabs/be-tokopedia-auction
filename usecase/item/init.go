package item

type Repo struct {
	Item
}

func NewUsecase(repoItem Item) *Repo {
	return &Repo{
		Item: repoItem,
	}
}
