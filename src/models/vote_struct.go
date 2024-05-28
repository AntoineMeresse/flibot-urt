package models

type Vote struct {
	Params []string
	PlayerId string
}

func (context *Context) NewVote(v Vote) {
	context.VoteChannel <- v
}