package main

import "errors"

type Color int

// 範例 1.
const (
	Black Color = iota
	Red
	Blue
)

func (c Color) String() string {
	return [...]string{"Black", "Red", "Blue"}[c]
}

// 範例 2.
type Bank int

const (
	BKTW Bank = iota
	CTBC
	HSBC
)

type bank struct {
	code string
	name string
}

var banks = map[Bank]bank{
	BKTW: {code: "004", name: "臺灣銀行"},
	CTBC: {code: "822", name: "中國信託"},
	HSBC: {code: "081", name: "匯豐銀行"},
}

func (b Bank) Code() string {
	return banks[b].code
}

func (b Bank) Name() string {
	return banks[b].name
}

func All() []Bank {
	return []Bank{BKTW, CTBC, HSBC}
}

func Of(code string) (Bank, error) {
	for _, bk := range All() {
		if bk.Code() == code {
			return bk, nil
		}
	}
	return -1, errors.New("No match")
}
