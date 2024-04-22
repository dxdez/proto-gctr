package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	items, err := fetchItems()
	if err != nil {
		log.Printf("error fetching items %v", err)
		return
	}
	count, err := fetchCount()
	if err != nil {
		log.Printf("error fetching count %v", err)
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("error fetching completed count %v", err)
	}
	data := ItemLists{
		Items:          items,
		Count:          count,
		CountChecked: completedCount,
	}
	currentTemplate.ExecuteTemplate(w, "base", data)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		return
	}
	_, err := insertItem(title)
	if err != nil {
		log.Printf("error inserting item %v", err)
		return
	}
	count, err := fetchCount()
	if err != nil {
		log.Printf("error fetching count %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	currentTemplate.ExecuteTemplate(w, "Form", nil)
	currentTemplate.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
}

func handleToggleTask(_ http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int %v", err)
		return
	}
	_, err = toggleItem(id)
	if err != nil {
		log.Printf("error toggling task %v", err)
		return
	}
}
