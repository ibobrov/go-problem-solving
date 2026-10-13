package main

import "fmt"

type MemberStat struct {
	ID     int
	Amount int
}

type Champs struct {
	IDS    []int
	Amount int
}

func main() {
	stat := [][]MemberStat{
		{{ID: 1, Amount: 5}, {ID: 2, Amount: 6}},
		{{ID: 1, Amount: 15}, {ID: 2, Amount: 10}},
	}
	champs := getChamps(stat)
	fmt.Printf("Champs: %v, %d", champs.IDS, champs.Amount)
}

func getChamps(stat [][]MemberStat) Champs {
	scores := make(map[int][]int)

	for _, day := range stat {
		for i := 0; i < len(day); i++ {
			memberId := day[i].ID
			memberAmount := day[i].Amount

			scores[memberId] = append(scores[memberId], memberAmount)
		}
	}

	var champs Champs
	found := false

	for memberId, slice := range scores {
		if len(slice) != len(stat) {
			continue
		}
		totalAmount := 0

		for _, dayScore := range slice {
			totalAmount += dayScore
		}

		if !found || totalAmount > champs.Amount {
			champs = Champs{
				IDS:    []int{memberId},
				Amount: totalAmount,
			}
			found = true
		} else if totalAmount == champs.Amount {
			champs.IDS = append(champs.IDS, memberId)
		}
	}

	return champs
}
