package blackjack
import "slices"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "other":
        return 0
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten":
        return 10
    case "jack":
        return 10
    case "queen":
        return 10
    case "king":
        return 10
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    total :=  ParseCard(card1) +  ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)
    autoWinCounter := []string{"jack", "queen", "king", "ace", "ten"}
   
	switch {
        case (card1 == "ace" && card2 == "ace"):
            return "P"
        case (total == 21 && slices.Contains(autoWinCounter, dealerCard)):
            return "S"
        case (total == 21 && !slices.Contains(autoWinCounter, dealerCard)):
            return "W"
        case (total >= 17 && total <= 20):
            return "S"
        case (total >= 12 && total <= 16 && dealerCardValue < 7):
            return "S"
        case (total <= 11):
            return "H"
        default:
        	return "H"
	}    
}
