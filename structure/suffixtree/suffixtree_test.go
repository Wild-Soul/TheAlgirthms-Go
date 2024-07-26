package suffixtree_test

import (
	"strings"
	"testing"

	"github.com/TheAlgorithms/Go/structure/suffixtree"
)

// Helper function to count occurrences of a pattern in a text
func countOccurrences(text, pattern string) int {
	return strings.Count(text, pattern)
}

func TestSuffixTreeConstruction(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"Empty string", ""},
		{"Single character", "a"},
		{"Repeated character", "aaaa"},
		{"Simple string", "banana"},
		{"Complex string", "mississippi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := suffixtree.NewSuffixTree(tt.text)
			st.Build()

			// Check if all suffixes are present
			for i := 0; i < len(tt.text); i++ {
				suffix := tt.text[i:]
				if !st.Contains(suffix) {
					t.Errorf("Suffix %s not found in tree", suffix)
				}
			}
		})
	}
}

func TestSuffixTreeSearch(t *testing.T) {
	tests := []struct {
		text    string
		pattern string
		want    bool
	}{
		{"banana", "ana", true},
		{"banana", "anan", false},
		{"mississippi", "ssi", true},
		{"mississippi", "missi", true},
		{"mississippi", "ississi", true},
		{"mississippi", "issippi", true},
		{"mississippi", "ppippi", false},
	}

	for _, tt := range tests {
		t.Run(tt.text+"-"+tt.pattern, func(t *testing.T) {
			st := suffixtree.NewSuffixTree(tt.text)
			st.Build()

			if got := st.Contains(tt.pattern); got != tt.want {
				t.Errorf("SuffixTree.contains(%q) = %v, want %v", tt.pattern, got, tt.want)
			}

			occurrences := countOccurrences(tt.text, tt.pattern)
			if tt.want && occurrences == 0 {
				t.Errorf("Pattern %q should occur in text %q", tt.pattern, tt.text)
			}
			if !tt.want && occurrences > 0 {
				t.Errorf("Pattern %q should not occur in text %q", tt.pattern, tt.text)
			}
		})
	}
}

func TestSuffixTreeEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"Very long string", strings.Repeat("abcdefghijklmnopqrstuvwxyz", 1000)},
		{"String with special characters", "Hello, World! 123 @#$%^&*()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := suffixtree.NewSuffixTree(tt.text)
			st.Build()

			// Check a few random suffixes
			for i := 0; i < 10; i++ {
				start := len(tt.text) / 10 * i
				suffix := tt.text[start:]
				if !st.Contains(suffix) {
					t.Errorf("Suffix starting at position %d not found in tree", start)
				}
			}
		})
	}
}
