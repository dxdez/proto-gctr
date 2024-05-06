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

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int %v", err)
		return
	}
	err = deleteItem(r.Context(), id)
	if err != nil {
		log.Printf("error deleting item %v", err)
	}
	count, err := fetchCount()
	if err != nil {
		log.Printf("error fetching count %v", err)
	}
	completedCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("error fetching completed count %v", err)
	}
	currentTemplate.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
	currentTemplate.ExecuteTemplate(w, "CompletedCount", map[string]any{"Count": completedCount, "SwapOOB": true})
}

func handleEditItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int %v", err)
		return
	}
	item, err := fetchItem(id)
	if err != nil {
		log.Printf("error fetching task with id %d %v", id, err)
		return
	}
	currentTemplate.ExecuteTemplate(w, "Item", map[string]any{"Item": item, "Editing": true})
}

func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int %v", err)
		return
	}
	title := r.FormValue("title")
	if title == "" {
		return
	}
	item, err := updateItem(id, title)
	if err != nil {
		log.Printf("error fetching task with id %d %v", id, err)
		return
	}
	currentTemplate.ExecuteTemplate(w, "Item", map[string]any{"Item": item})
}

func handleOrderItems(_ http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing form %v", err)
	}
	var values []int
	for k, v := range r.Form {
		if k == "item" {
			for _, v := range v {
				value, err := strconv.Atoi(v)
				if err != nil {
					log.Printf("error parsing id into int %v", err)
					return
				}
				values = append(values, value)
			}
		}
	}
	err = orderItem(r.Context(), values)
	if err != nil {
		log.Printf("error ordering tasks %v", err)
	}
}

