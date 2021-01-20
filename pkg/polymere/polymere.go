// Package polymere implement functions to manipulate a polymere
//  A polymere is composed by letters in alphabet,
//  minuscule (consider as negative) or Capital (considered as positive).
package polymere

import (
	"bytes"
	"fmt"
	"unicode"
)

const (
	// ElementError the "error" if not letter
	ElementError = "arguments should be only letters"
)

// Check if elements are opposites
func isOpposite(a, b uint8) (bool, error) {
	var r, ar, br rune
	ar = rune(a)
	br = rune(b)
	if !unicode.IsLetter(ar) || !unicode.IsLetter(br) {
		return false, fmt.Errorf("isOpposite(%c, %c). err: %s", a, b, ElementError)
	}
	if unicode.IsUpper(ar) && unicode.IsLower(br) {
		r = unicode.ToLower(ar)
	} else if unicode.IsUpper(br) && unicode.IsLower(ar) {
		r = unicode.ToUpper(ar)
	}
	return r == br, nil
}

// React method to acquire the reaction between opposite polarisation
func React(polymere string) (string, error) {
	if len(polymere) == 1 {
		return polymere, nil
	}
	var buf bytes.Buffer
	idx := 0
	for idx < len(polymere)-1 {
		//check if are opposite polymere
		ok, err := isOpposite(polymere[idx], polymere[idx+1])
		//check error
		if err != nil {
			return "", err
		}
		if ok {
			//skip
			idx += 2
			// write to buffer
			_, err := buf.WriteString(polymere[idx:])
			if err != nil {
				return "", fmt.Errorf("React:WriteString:%s", err.Error())
			}
			// replace polymere
			polymere = buf.String()
			// restart
			idx = 0
			// reset buffer
			buf.Reset()
			continue
		}

		buf.WriteByte(polymere[idx])
		idx++
	}
	//return polymere
	return polymere, nil
}

//build a map of unique elements in polymer. all the element will be consider as
func getBridleMap(polymer string) map[rune]string {
	m := make(map[rune]string)
	for _, r := range polymer {
		e := unicode.ToLower(r)
		if _, exist := m[e]; !exist {
			// remove the bridle element
			m[e] = trimPolymer(polymer, e)
		}
	}
	//return elements
	return m
}

//trim return a slice of the string white the cutset removed
func trimPolymer(polymer string, cutset rune) string {
	var b bytes.Buffer

	//get opposite
	o := unicode.ToUpper(cutset)
	for _, l := range polymer {
		if l != cutset && l != o {
			b.WriteRune(l)
		}
	}
	return b.String()
}

// LenMinimalReaction return the length of the smallest reaction removing one element
func LenMinimalReaction(polymer string) (int, error) {
	firstReaction, err := React(polymer)
	if err != nil {
		return 0, fmt.Errorf("react err: %s", err.Error())
	}
	// look for unique elements in polymer.
	m := getBridleMap(firstReaction)
	smallest := len(firstReaction)
	// loop in posibilities
	for _, s := range m {
		r, err := React(s)
		if err != nil {
			return 0, fmt.Errorf("react err: %s", err.Error())
		}
		reactionLen := len(r)
		// check if reaction is the smallest
		if reactionLen < smallest {
			smallest = reactionLen
		}
	}
	// return smallest
	return smallest, nil
}
