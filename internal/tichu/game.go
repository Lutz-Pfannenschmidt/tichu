package tichu

type Game struct {
	Teams [2]*Team
	Deck  *Deck
	Stack []*Combination
	Turn  int
	Wish  int
}

func NewGame(playerName1 string, playerName2 string, playerName3 string, playerName4 string) *Game {
	g := &Game{
		Teams: [2]*Team{},
		Deck:  NewDeck(),
		Stack: []*Combination{},
	}
	g.Teams[0] = NewTeam(g, 0, playerName1, playerName3)
	g.Teams[1] = NewTeam(g, 1, playerName2, playerName4)
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
