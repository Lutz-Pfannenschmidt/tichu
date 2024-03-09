package tichu

type Game struct {
	Teams [2]*Team
	Deck  *Deck
	Stack *Deck
}

func NewGame() *Game {
	g := &Game{
		Teams: [2]*Team{},
		Deck:  NewDeck(),
		Stack: &Deck{},
	}
	g.Teams[0] = NewTeam(g)
	g.Teams[1] = NewTeam(g)
	return g
}

func (g *Game) Deal() {
	for i := 0; i < 8; i++ {
		for _, team := range g.Teams {
			for _, player := range team.Players {
				player.Hand = append(player.Hand, g.Deck.Draw())
			}
		}
	}

	for _, team := range g.Teams {
		for _, player := range team.Players {
			player.SortHand()
		}
	}
}
