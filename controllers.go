package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func getGames(w http.ResponseWriter, r *http.Request) {
	var games []Game
	rows, err := db.Query("select * from games")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Records could not retrieved"))
		return
	}
	for rows.Next() {
		game := Game{}
		rows.Scan(&game.ID, &game.Name, &game.LatestVersion, &game.Category, &game.Descr, &game.Producer)
		games = append(games, game)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(games)
}

func getGame(w http.ResponseWriter, r *http.Request) {

}

func createGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &game)
	defer r.Body.Close()
	vldErr := validate.Struct(game)
	if vldErr != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, vldErr.Error())
		return
	}
	stmt, err := db.Prepare("insert into games(name,latest_version,category,descr,producer) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Something went wrong when creating the game"))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(game.Name, game.LatestVersion, game.Category, game.Descr, game.Producer)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Something went wrong when creating the game"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Success"))

}

func updateGame(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var game Game
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &game)
	defer r.Body.Close()
	vldErr := validate.Struct(game)
	if vldErr != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, vldErr.Error())
		return
	}
	stmt, err := db.Prepare("UPDATE games SET name=$1, latest_version=$2,category=$3,descr=$4,producer=$5 where id=$6")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Something went wrong when updating the game"))
		return
	}
	_, err = stmt.Exec(game.Name, game.LatestVersion, game.Category, game.Descr, game.Producer, id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Something went wrong when updating the game"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Success"))

}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	stmt, err := db.Prepare("DELETE FROM games WHERE id=$1")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Something went wrong when deleting the game"))
		return
	}
	stmt.Exec(id)
	w.WriteHeader(200)
	w.Write([]byte("Success"))
}
