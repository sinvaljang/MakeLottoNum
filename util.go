package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type lottoNum struct {
	nums []int
}

func initLottoNs(l *lottoNum) {
	for i := 0; i < NUMS_MAX; i++ {
		l.nums = append(l.nums, 0)
	}
}

func chkLottoN(l lottoNum) {
	for i, value := range l.nums {
		fmt.Printf("%d _ value = %d\n", i, value)
	}
}

func makeJson(l *lottoNum) []byte {

	val, err := json.Marshal(setLottoNum(*l))
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	return val
}

func setLottoNum(l lottoNum) lottoNum {
	var inputN int
	initLottoNs(&l)
	//r := rand.New(rand.NewSource(MAX_NUM))

	for i, _ := range l.nums {

		rand.Seed(time.Now().UnixNano())

		inputN = rand.Intn(MAX_NUM)

		if i == 0 {
			l.nums[0] = inputN
			continue
		}

		temps := l.nums[:i]

		for ii := 0; ii < i; ii++ {
			if temps[ii] == inputN {
				inputN = rand.Intn(MAX_NUM)
				ii = 0
			}
			if ii == i-1 {
				l.nums[i] = inputN
				break
			}
		}
	}

	chkLottoN(l)

	return l
}
