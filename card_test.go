package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Heart})
	fmt.Println(Card{Rank: Five, Suit: Spade})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})
	//Output:
	//Ace of Clubs
	//Jack of Hearts
	//Five of Spades
	//Ten of Diamonds
	//Joker
}

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 52 {
		t.Error("Wrong no of cards")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	card := Card{Rank: Three, Suit: Heart}
	if cards[15] != card {
		t.Error("Expected Two of Hearts and got", cards[15])
	}
}

func TestSort(t *testing.T) {
	cards := NewDeck(Sort(Less))
	card := Card{Rank: Three, Suit: Heart}
	if cards[15] != card {
		t.Error("Expected Two of Hearts and got", cards[15])
	}
}

// func TestShuffleDec(t *testing.T)  {
// 	cards := NewDeck()
// 	if cards == ShuffleDec(cards){
// 		t.Error("cards not shuffled")
// 	}
// }

func TestAddJoker(t *testing.T) {
	n := 3
	count := 0
	cards := NewDeck(AddJoker(n))
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != n {
		t.Error(fmt.Sprintf("Expected %d jokers and got %d", n, count))
	}
}

func TestFilter(t *testing.T) {
	f := func(card Card) bool {
		return card.Rank == Three || card.Rank == Jack
	}
	cards := NewDeck(Filter(f))

	for _, card := range cards {
		if card.Rank == Three || card.Rank == Jack {
			t.Error("Not Expected Three and Five")
		}

	}
}

func TestMultiDeck(t *testing.T){
	n:=4
	cards:=NewDeck(MultiDeck(n))
	if len(cards)!= n*52{
		t.Errorf("Expected %d and got %d",n*52,len(cards))
	}
}