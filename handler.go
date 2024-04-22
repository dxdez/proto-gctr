package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func handleGetItems(w http.ResponseWriter, _ *http.Request) {
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

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		return
	}
	item, err := insertItem(title)
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
	currentTemplate.ExecuteTemplate(w, "Item", map[string]any{"Item": item, "SwapOOB": true})
	currentTemplate.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
}

func handleToggleItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int %v", err)
		return
	}
	_, err = toggleItem(id)
	if err != nil {
		log.Printf("error toggling item %v", err)
		return
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("error fetching completed count %v", err)
		return
	}
	currentTemplate.ExecuteTemplate(w, "CompletedCount", map[string]any{"Count": completedCount, "SwapOOB": true})
}
