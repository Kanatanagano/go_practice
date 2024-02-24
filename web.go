package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// analyze queryparameters
	query := r.URL.Query()
	name := query.Get("name")

	// create a response map
	response := map[string]string{
		"message": "Hello " + name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// カテゴリを固定する
var categories = []string{"math", "science", "history", "english"}

func categoriesoHandler(w http.ResponseWriter, r *http.Request) {
	// create query parameters
	query := r.URL.Query()
	categoryName := query.Get("category")

	// カテゴリが選択された場合
	if categoryName != "" {
		response := map[string]string{
			"category": categoryName,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return

	}
	// カテゴリが選択されなかった場合
	response := map[string]interface{}{
		"categories": categories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	operator := query.Get("o")
	strx := query.Get("x")
	stry := query.Get("y")

	// x, yが数字かどうかを確認
	// x, yはfloat64に変換して幅広い数値を扱えるようにする

	x, err := strconv.ParseFloat(strx, 64)
	if err != nil {
		http.Error(w, "x is not a number", http.StatusBadRequest)
		return
	}
	y, err := strconv.ParseFloat(stry, 64)
	if err != nil {
		http.Error(w, "y is not a number", http.StatusBadRequest)
		return
	}

	//URLクエリパラメータに+を直接入力するとエラーが発生するため、その例外処理を追加
	if operator == " " {
		operator = "+"
	}

	// x, yを計算する
	var result float64
	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			http.Error(w, "y is 0", http.StatusBadRequest)
			return
		}
		result = x / y
	default:
		http.Error(w, "invalid operator", http.StatusBadRequest)
		return
	}

	//レスポンスマップを作成
	response := map[string]float64{
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	fmt.Println("Starting server on port 8010")
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/categories", categoriesoHandler)
	http.HandleFunc("/api/calculator", calculatorHandler)

	http.ListenAndServe(":8010", nil)
}
