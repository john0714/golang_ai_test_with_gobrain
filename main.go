package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain"
)

var learningCount int = 1;
var learningRate float64 = 0.6;
var momentumFactor float64 = 0.4;

func main() {
	// set the random seed to 0
	// gobrainは初期化の際randの(-1, 1)範囲を使うためSeedを0に設定する必要がある
	rand.Seed(0)

	// 学ぶパターン(例：AND)
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {0}},
		{{1, 0}, {0}},
		{{1, 1}, {1}},
	}

	// モデル初期化
	brain := &gobrain.FeedForward{}

	// init the Neural Network
	// 引数は左から順に、入力数, 中間層の次元数, 答えの数を設定する
	// 2 inputs, 2 hidden nodes and 1 output
	brain.Init(2, 2, 1)

	for i := 0; i < 5; i++ {
		fmt.Printf("今回の学習回数 : %v\n", learningCount)
		
		// 学習メソッド
	    // 引数は左から順に学習させるパターン,　学習回数, 学習率(学習速度), 運動量係数,　エラー出力
		brain.Train(patterns, learningCount, learningRate, momentumFactor, false)
		
		/*
		// テスト
		fmt.Println("gobrainテスト結果")
		fmt.Println("========================================")
		brain.Test(patterns)
		fmt.Println("========================================")
		fmt.Println("")
		*/
		
		// 実際データ作成
	    inputs := [][]float64{
	      {0,0},
	      {0,1},
	      {1,0},
	      {1,1},
	    }
	    
	    // 結果表示
	    fmt.Println("実際データの結果")
	    fmt.Println("========================================")
	    for _, input := range inputs{
	      result := brain.Update(input) // データ入力
	      fmt.Println(input[0],",",input[1],":",result[0]) // 表示
	    }
	    fmt.Println("========================================")	
	    fmt.Println("")
	    
	    //　学習回数増加
		learningCount = learningCount*10
	}
}