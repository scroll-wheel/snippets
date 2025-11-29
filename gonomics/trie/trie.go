package trie

import (
	"unicode/utf8"
)

type Trie[T any] struct {
	label    rune
	values   []T
	children []Trie[T]
}

func (t *Trie[T]) Reset() {
	t.label = 0
	t.values = []T{}
	t.children = []Trie[T]{}
}

func (t *Trie[T]) Insert(key string, values ...T) {
	if key == "" {
		t.values = append(t.values, values...)
		return
	}

	runeValue, width := utf8.DecodeRuneInString(key)

	for i := range t.children {
		if t.children[i].label == runeValue {
			t.children[i].Insert(key[width:], values...)
			return
		}
	}

	child := Trie[T]{label: runeValue}
	child.Insert(key[width:], values...)
	t.children = append(t.children, child)
}

func (t Trie[T]) Find(key string, exactMatchOnly bool) []T {
	if key == "" {
		return t.values
	}

	runeValue, width := utf8.DecodeRuneInString(key)

	for _, n := range t.children {
		if n.label == runeValue {
			return n.findHelper(key[width:], exactMatchOnly)
		}
	}

	return []T{}
}

func (t Trie[T]) findHelper(key string, exactMatchOnly bool) []T {
	if key == "" {
		return t.values
	}

	result := []T{}
	runeValue, width := utf8.DecodeRuneInString(key)

	for _, n := range t.children {
		if n.label == runeValue {
			result = append(result, n.findHelper(key[width:], exactMatchOnly)...)
		} else if !exactMatchOnly {
			result = append(result, n.findHelper(key[width:], true)...)
		}
	}

	return result
}
