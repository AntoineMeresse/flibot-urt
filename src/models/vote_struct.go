package models

type Vote struct {
	Params []string
	PlayerId string
}

func (server *Context) NewVote(v Vote) {
	server.VoteChannel <- v
}