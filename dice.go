package main

import (
	"crypto/rand"
	"math/big"
)

// GetRandomInt 0から(num -1)までのランダムな整数値を返す
func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

// GetDiceRoll 1からnumまでのランダムな整数値を返す
func GetDiceRoll(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64()) + 1
}

// GetRandomBetween lowからhighの間でランダムな整数を返す
func GetRandomBetween(low int, high int) int {
	var randy int = -1

	for {
		randy = GetDiceRoll(high)
		//	lowよりも大きかったら抽選終了
		if randy >= low {
			break
		}
	}

	return randy
}
