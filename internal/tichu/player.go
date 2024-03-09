package tichu

import "sort"

type Player struct {
	Name  string
	Hand  []*Card
	Ready bool
	Team  *Team
}

func NewPlayer(name string, team *Team) *Player {
	return &Player{
		Name: name,
		Hand: make([]*Card, 0),
		Team: team,
	}
}

func (p *Player) SortHand() {
	sort.Slice(p.Hand, func(i, j int) bool {
		return p.Hand[i].Type.Value() < p.Hand[j].Type.Value()
	})
}
