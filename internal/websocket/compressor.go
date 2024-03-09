package websocket

import (
	"encoding/binary"
	"github.com/Lutz-Pfannenschmidt/tichu/internal/tichu"
)

func Compress(player *tichu.Player) []byte {
	var res []byte

	handCards := CompressCombination(&player.Hand)
	res = append(res, handCards[:]...)

	playedCards := CompressCombination(&player.Hand)
	res = append(res, playedCards[:]...)

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
