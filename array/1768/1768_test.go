package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{"same length", "abc", "pqr", "apbqcr"},
		{"longer second word", "ab", "pqrs", "apbqrs"},
		{"longer first word", "abcd", "pq", "apbqcd"},
		{"blank first word", "", "pqr", "pqr"},
		{"blank second word", "abc", "", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeAlternately(tt.word1, tt.word2))
		})
	}
}
