package main

type Item struct {
	ID int
	Title string
	Checked bool
}

type ItemLists struct {
	Items []Item
	Count Int
	CountChecked Int
}

func fetchItems() ([]Item, error) {
	var itemList []Item
	
	itemRows, err := DB.Query("SELECT id, title, checked FROM items ORDER BY position;")
	if err != nil {
		return []Item{}, err
	}
	defer itemRows.Close()

	for itemRows.Next() {
		currentItem := Item{}
		err := itemRows.Scan(&currentItem.ID, &currentItem.Title, &currentItem.Checked)
		if err != nil {
			return []Item{}, err
		}
		itemList = append(itemList, currentItem)
	}

	return itemList, nil
}

func fetchItem(ID int) (Item, error) {
	var currentItem Item
	err := DB.QueryRow("SELECT id, title, checked FROM items WHERE id = (?)", ID).Scan(&currentItem.ID, &currentItem.Title, &currentItem.Completed)
	if err != nil {
		return Item{}, err
	}
	return currentItem, nil
}
