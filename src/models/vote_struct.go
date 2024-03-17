package models

type Vote struct {
	Params []string
	PlayerId string
}

func (server *Server) NewVote(v Vote) {
	server.VoteChannel <- v
}