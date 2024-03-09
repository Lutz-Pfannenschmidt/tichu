package websocket

import (
	"encoding/binary"
	"github.com/Lutz-Pfannenschmidt/tichu/internal/tichu"
)

func Compress(player *tichu.Player) []byte {
	// H H H H H H H P P P P P P P aaaabbbb ccccpqrst wwwwxxxx 000000tt
	// H = 1 byte of your hand cards
	// P = 1 byte of the cards lying in the middle
	// a = 1 bit of the count of hand cards of the opponent with lower game id
	// b = 1 bit of the count of hand cards of the opponent with higher game id
	// c = 1 bit of the count of hand cards of your teammate
	// p = the bit representing whether the player with game id 0 has won tricks
	// q = the bit representing whether the player with game id 1 has won tricks
	// r = the bit representing whether the player with game id 2 has won tricks
	// s = the bit representing whether the player with game id 3 has won tricks
	// w = 1 bit of the wished Value
	// x = 1 bit of the card value the phoenix became if it was played as single card
	// t = 1 bit of the game id of the player whose turn it is

	var res []byte

	handCards := CompressCombination(&player.Hand)
	res = append(res, handCards[:]...)

	playedCards := CompressCombination(&player.Team.Game.Stack[len(player.Team.Game.Stack)].Cards)
	res = append(res, playedCards[:]...)

	ownTeam := player.Team
	var teamMate *tichu.Player

	if player.Team.Players[0] == player {
		teamMate = player.Team.Players[1]
	} else {
		teamMate = player.Team.Players[0]
	}

	var oppositeTeam *tichu.Team

	if player.Team.Game.Teams[0] == ownTeam {
		oppositeTeam = player.Team.Game.Teams[1]
	} else {
		oppositeTeam = player.Team.Game.Teams[0]
	}

	oppPlayer1 := oppositeTeam.Players[0]
	oppPlayer2 := oppositeTeam.Players[1]

	oppHandCardsCount := byte((len(oppPlayer1.Hand) << 4) | len(oppPlayer2.Hand))
	res = append(res, oppHandCardsCount)

	teamMateCardCountAndTricks := byte(len(teamMate.Hand) << 4)

	if len(oppPlayer1.Earnings) > 0 {
		teamMateCardCountAndTricks |= 1 << oppPlayer1.ID
	}
	if len(oppPlayer2.Earnings) > 0 {
		teamMateCardCountAndTricks |= 1 << oppPlayer2.ID
	}
	if len(teamMate.Earnings) > 0 {
		teamMateCardCountAndTricks |= 1 << teamMate.ID
	}
	if len(player.Earnings) > 0 {
		teamMateCardCountAndTricks |= 1 << player.ID
	}

	res = append(res, teamMateCardCountAndTricks)

	wishAndPhoenixValue := byte(player.Team.Game.Wish<<4) | byte(player.Team.Game.Stack[len(player.Team.Game.Stack)].Base)
	res = append(res, wishAndPhoenixValue)

	res = append(res, byte(player.Team.Game.Turn&3))

	return res
}

func CompressCombination(comb *[]*tichu.Card) [7]byte {
	var res uint64 = 0
	for _, card := range *comb {
		res |= 1 << CardToId(card)
	}
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, res)
	return [7]byte(bs)
}

func DecompressCombination(target *[]*tichu.Card, data [7]byte) {
	bs := []byte{data[0], data[1], data[2], data[3], data[4], data[5], data[6], 0}
	var combId uint64 = binary.LittleEndian.Uint64(bs)
	*target = []*tichu.Card{}
	for i := 0; i < 56; i++ {
		if combId&1 == 1 {
			*target = append(*target, IdToCard(i))
		}
		combId >>= 1
	}
}

func CardToId(c *tichu.Card) int {
	if 2 <= c.Type && c.Type <= 14 {
		return (int(c.Type)-1)*4 + int(c.Color)
	}
	if int(c.Type) <= 1 {
		return int(c.Type)
	} else {
		return int(c.Type) - 13
	}
}

func IdToCard(id int) *tichu.Card {
	if id >= 4 {
		return &tichu.Card{Type: tichu.CardType((id >> 2) + 1), Color: tichu.CardColor(id & 3)}
	}
	if id < 2 {
		return &tichu.Card{Type: tichu.CardType(id), Color: tichu.Special}
	} else {
		return &tichu.Card{Type: tichu.CardType(id + 13), Color: tichu.Special}
	}
}
