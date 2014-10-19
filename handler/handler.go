package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zacharycarter/monster-codex/db"
	"github.com/zacharycarter/monster-codex/entity"
	"github.com/zacharycarter/monster-codex/repository"
	"github.com/zacharycarter/monster-codex/template"
	"github.com/zacharycarter/monster-codex/util"
)

var (
	monsters = &repository.MonsterRepository{db.DbMap}
)

func Home(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, []string{"header.html", "home.html", "footer.html"})
}

func Monsters(w http.ResponseWriter, r *http.Request) {
	m, err := monsters.List()
	util.CheckErr(err, "Error handling monsters api request: ")

	if m == nil {
		m = []*entity.Monster{}
	}

	jsonResponse(w, m)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	util.CheckErr(err, "Error marshalling data: ")

	fmt.Fprintf(w, string(payload))
}
