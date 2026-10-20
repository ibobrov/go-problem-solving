package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChamps_Standard(t *testing.T) {
	stat := [][]MemberStat{
		{{ID: 1, Amount: 5}, {ID: 2, Amount: 6}},
		{{ID: 1, Amount: 15}, {ID: 2, Amount: 10}},
	}

	got := getChamps(stat)
	want := Champs{IDS: []int{1}, Amount: 20}

	assert.ElementsMatch(t, want.IDS, got.IDS, "должен вернуть только ID чемпиона")
	assert.Equal(t, want.Amount, got.Amount, "сумма очков должна быть максимальной")
}
