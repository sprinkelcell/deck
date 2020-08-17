//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
	Joker
)

var suits = [...]Suit{Spade, Heart, Diamond, Club}

type Rank uint8

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)
const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, Less(cards))
		return cards
	}
}
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}
func absRank(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}

func ShuffleDec(cards []Card) []Card {
	output := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(cards)) {
		output[i] = cards[j]
	}
	return output
}

func AddJoker(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var output []Card
		for _, card := range cards {
			if !f(card) {
				output = append(output, card)
			}
		}
		return output
	}
}

func MultiDeck(n int) func(cards []Card) []Card {
	var output []Card
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			output = append(output, cards...)
		}
		return output
	}

}
