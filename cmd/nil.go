package main

import "fmt"

type BillionDollarMistake struct {
	wat bool
}

func (bdm *BillionDollarMistake) Hello() {
	fmt.Println("Thanks Tony!")
}

func (bdm *BillionDollarMistake) Wat() bool {
	return bdm.wat
}

func (bdm *BillionDollarMistake) New() *BillionDollarMistake {
	return &BillionDollarMistake{}
}

type NilOrNotNil interface {
}

func Question() NilOrNotNil {
	var b *NilOrNotNil
	return b
}

func main() {
	var bdm *BillionDollarMistake
	bdm.Hello()

	b := bdm.New()
	b.Wat()
	mistake := Question()
	fmt.Println(mistake)
	if mistake != (*NilOrNotNil)(nil) {
		fmt.Println("billion dollar mistake is not nil")
	}
}
