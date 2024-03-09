package tichu

import "sort"

type Player struct {
	Name     string
	Hand     []*Card
	Ready    bool
	Team     *Team
	Earnings []*Card
	ID       int
}

func NewPlayer(name string, team *Team, id int) *Player {
	return &Player{
		Name:     name,
		Hand:     make([]*Card, 0),
		Team:     team,
		Earnings: make([]*Card, 0),
		ID:       id,
	}
}

func (p *Player) SortHand() {
	sort.Slice(p.Hand, func(i, j int) bool {
		return p.Hand[i].Type.Value() < p.Hand[j].Type.Value()
	})
}
