package main

import (
	"log"
	"net/http"
)

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	items, err := fetchItems()
	if err != nil {
		log.Printf("error fetching tasks %v", err)
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
