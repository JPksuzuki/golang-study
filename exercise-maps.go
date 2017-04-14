package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var want = make(map[string]int) //戻り値

	var words []string = strings.Fields(s) //単語ごとに分解

	for _, word := range words { //各単語について
		_, ok := want[word] //単語に対する要素が存在するか
		if !ok {
			want[word] = 1 //存在しない場合 → wordの数を1に初期化
		} else {
			want[word]++ //存在する場合 → wordの数を+1
		} 
	}
	return want
}

func main() {
	wc.Test(WordCount)
}
