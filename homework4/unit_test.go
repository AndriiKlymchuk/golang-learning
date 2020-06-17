package homework4

import (
	"math/rand"
	"strings"
	"testing"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func Test_calculateRank(t *testing.T) {
	if initRank := calculateRank("", ""); !isValidRank(initRank) {
		t.Errorf("Ivalid initial rank %v", initRank)
	}
	var smaller, bigger rank
	for i := 0; i < 100000; i++ {
		a, b := randRanksPair()
		rank := calculateRank(a, b)
		if a < b {
			smaller, bigger = a, b
		} else {
			smaller, bigger = b, a
		}
		if !isValidRank(rank) || !(smaller < rank && rank < bigger) {
			t.Errorf("calculateRank(%v, %v) returned incorrect result %v", a, b, rank)
		}
	}
}

func randRanksPair() (rank, rank) {
	var a, b rank
	a = randRank()
	for {
		if b = randRank(); a != b {
			break
		}
	}
	return a, b
}

func randRank() rank {
	b := make([]byte, rand.Intn(6))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	if len(b) > 0 && b[len(b)-1] == 'a' {
		b = append(b, letters[1:][rand.Intn(len(letters)-1)])
	}
	return rank(b)
}

func isValidRank(r rank) bool {
	if r == "" && r[len(r)-1] == 'a' {
		return false
	}
	for _, c := range r {
		if !strings.ContainsRune(letters, rune(c)) {
			return false
		}
	}
	return true
}