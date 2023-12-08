package bid

type Repo struct {
	Bid
}

func NewUsecase(repoBid Bid) *Repo {
	return &Repo{
		Bid: repoBid,
	}
}
