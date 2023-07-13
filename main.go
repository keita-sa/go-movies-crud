package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // Movieのデータを保持するスライスの作成

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダの設定
	json.NewEncoder(w).Encode(movies)                  // Goのデータ型からJSONに変換
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダの設定
	params := mux.Vars(r)                              // gorilla/muxの機能を使ってパスパラメータを取得
	for _, item := range movies {                      // for _, v := range でループ。_,は回数が必要ない場合（通常はi）
		if item.ID == params["id"] { // itemのIDと、パラメータのidが同じ場合
			json.NewEncoder(w).Encode(item) // JSONデータを書き込む
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダの設定
	var movie Movie                                    // structで定義された新しい型(type)、Movie型の変数movie
	_ = json.NewDecoder(r.Body).Decode(&movie)         // JSONデータを読み込み、その結果を構造体Movieのmovieのアドレスに格納する
	movie.ID = strconv.Itoa(rand.Intn(100))            // strconv.Itoaで数値を文字列に変換
	movies = append(movies, movie)                     // moviesスライスの最後に要素を追加
	json.NewEncoder(w).Encode(movie)                   // JSONデータを書き込む
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダの設定
	params := mux.Vars(r)                              // クエリパラメータの取得

	// delete the movie with the i.d that you've sent
	// add a new movie - the movie that we send in the body of postman
	for index, item := range movies { // loop over the movies, range
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // スライスをappendで追加する（第二引数の後ろに...を加える）
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // JSONデータを読み込み、その結果を構造体Movieのmovieのアドレスに格納する
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // JSONデータを書き込む
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダの設定
	params := mux.Vars(r)                              // gorilla/muxの機能を使ってパスパラメータを取得
	for index, item := range movies {                  // loop over the movies, range

		if item.ID == params["id"] { // itemのIDと、パラメータのidが同じ場合
			movies = append(movies[:index], movies[index+1:]...) // スライスをappendで追加
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter() // NewRouter関数でマルチプレクサを生成、ルーターのイニシャライズ

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454369", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
