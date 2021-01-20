package polymere

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type reactTest struct {
	name     string
	polymere string
	reaction string
}

// Test case for function with reaction
func drivenReactTest() []reactTest {
	//driven test
	return []reactTest{
		{
			name:     "one",
			polymere: "baAC",
			reaction: "bC",
		},
		{
			name:     "two",
			polymere: "cabBA",
			reaction: "c",
		},
		{
			name:     "three",
			polymere: "CaAC",
			reaction: "CC",
		},
		{
			name:     "four",
			polymere: "abAB",
			reaction: "abAB",
		},
		{
			name:     "five",
			polymere: "a",
			reaction: "a",
		},
		{
			name:     "six",
			polymere: "dabBADcC",
			reaction: "",
		},
		{
			name:     "seven",
			polymere: "dabAcCaCBAcCcaDA",
			reaction: "dabCBAcaDA",
		},
		{
			name:     "eight",
			polymere: "LONGSTRINGWITHMIXOFSMALLANDCAPS",
			reaction: "LONGSTRINGWITHMIXOFSMALLANDCAPS",
		},
		{
			name:     "nine",
			polymere: "LlONGsSTRINGaODUIpytThHvVbBnNYPiudoAgWITHMIXOfFasSMALlANDcCAPS",
			reaction: "ONGTRINWITHMIXOaMAANDAPS",
		},
		{
			name:     "ten",
			polymere: "bBtTRRAC",
			reaction: "RRAC",
		},
	}
}

//TestReaction calls react with a polymers, checking his reaction
func TestReaction(t *testing.T) {
	for _, c := range drivenReactTest() {
		t.Run(c.name, func(t1 *testing.T) {
			r, err := React(c.polymere)
			require.NoError(t, err)
			require.Equal(t1, c.reaction, r)
		})
	}
}

type isOppositeTest struct {
	name string
	a    uint8
	b    uint8
	ok   bool
	err  error
}

func drivenOppositeTest() []isOppositeTest {
	return []isOppositeTest{
		{
			name: "one",
			a:    'A',
			b:    'a',
			ok:   true,
		},
		{
			name: "two",
			a:    'a',
			b:    'A',
			ok:   true,
		},
		{
			name: "three",
			a:    'a',
			b:    'a',
		},
		{
			name: "four",
			a:    'A',
			b:    'A',
		},
		{
			name: "five",
			a:    'b',
			b:    'A',
		},
		{
			name: "six",
			a:    'a',
			b:    'B',
		},
		{
			name: "seven",
			a:    '9',
			b:    's',
			err:  fmt.Errorf("isOpposite(9, s). err: arguments should be only letters"),
		},
	}
}
func TestIsOposite(t *testing.T) {
	for _, c := range drivenOppositeTest() {
		t.Run(c.name, func(t1 *testing.T) {
			ok, err := isOpposite(c.a, c.b)
			if err != nil {
				assert.Equal(t, err, c.err)
			}
			require.Equal(t1, c.ok, ok)
		})
	}
}

type LenMinTest struct {
	name     string
	polymere string
	smallest int
}

// Test case for function to compute minimal len
func lenMimalTest() []LenMinTest {
	//driven test
	return []LenMinTest{
		{
			name:     "one",
			polymere: "baBtTRRBAC",
			smallest: 4,
		},
		{
			name:     "two",
			polymere: "cabBA",
			smallest: 0,
		},
		{
			name:     "three",
			polymere: "CaFghjuipPoOIUJHGFAC",
			smallest: 2,
		},
		{
			name:     "four",
			polymere: "abYPPyAB",
			smallest: 4,
		},
		{
			name:     "five",
			polymere: "a",
			smallest: 0,
		},
		{
			name:     "six",
			polymere: "dabBADcC",
			smallest: 0,
		},
		{
			name:     "seven",
			polymere: "dabAcCaCBAcCcaDA",
			smallest: 4,
		},
		{
			name:     "seven",
			polymere: "dabAcCaCBAcCcaDA",
			smallest: 4,
		},
	}
}

func TestLenMinimalReaction(t *testing.T) {
	for _, c := range lenMimalTest() {
		t.Run(c.name, func(t1 *testing.T) {
			sm, err := LenMinimalReaction(c.polymere)
			assert.NoError(t, err)
			require.Equal(t1, c.smallest, sm)
		})
	}
}

func benchMarkReact(b *testing.B, bc reactTest) {
	b.Run(bc.name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := React(bc.polymere)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func BenchmarkReact(b *testing.B) {
	for _, bc := range drivenReactTest() {
		benchMarkReact(b, bc)
	}
}

func benchMarkLenMininalReaction(b *testing.B, bc reactTest) {
	b.Run(bc.name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := LenMinimalReaction(bc.polymere)
			assert.NoError(b, err)
		}
	})
}
func BenchmarkLenMinimalReaction(b *testing.B) {
	for _, bc := range drivenReactTest() {
		benchMarkLenMininalReaction(b, bc)
	}
}
