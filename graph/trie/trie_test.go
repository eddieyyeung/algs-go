// Package trie
// https://www.geeksforgeeks.org/trie-insert-and-search/
package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	keys := []string{"the", "a", "there", "answer", "any", "by", "bye", "their"}

	trie := New()
	for _, key := range keys {
		trie.Insert(key)
	}

	type args struct {
		search string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{"the"}, want: true},
		{name: "", args: args{"these"}, want: false},
		{name: "", args: args{"their"}, want: true},
		{name: "", args: args{"thaw"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trie.Search(tt.args.search); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
