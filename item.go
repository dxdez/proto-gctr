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
	err := DB.QueryRow("SELECT id, title, checked FROM items WHERE id = (?)", ID).Scan(&currentItem.ID, &currentItem.Title, &currentItem.Checked)
	if err != nil {
		return Item{}, err
	}
	return currentItem, nil
}

func updateItem(ID int, title string) (Item, error) {
	var currentItem Item
	err := DB.QueryRow("UPDATE items SET title = (?) WHERE id = (?) returning id, title, checked", title, ID).Scan(&item.ID, &item.Title, &item.Checked)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func fetchCount() (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM items;").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func fetchCompletedCount() (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM items WHERE checked = 1;").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func insertItem(title string) (Item, error) {
	count, err := fetchCount()
	if err != nil {
		return Item{}, err
	}
	var id int
	err = DB.QueryRow("INSERT INTO items (title, position) VALUES (?, ?) returning id", title, count).Scan(&id)
	if err != nil {
		return Item{}, err
	}
	item := Item{ID: id, Title: title, Checked: false}
	return item, nil
}

func deleteItem(ctx context.Context, ID int) error {
	_, err := DB.Exec("DELETE FROM items WHERE id = (?)", ID)
	if err != nil {
		return err
	}
	rows, err := DB.Query("SELECT id FROM items ORDER BY position")
	if err != nil {
		return err
	}
	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()
	for idx, id := range ids {
		_, err := DB.Exec("UPDATE items SET position = (?) WHERE id = (?)", idx, id)
		if err != nil {
			return err
		}
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func orderItem(ctx context.Context, values []int) error {
	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()
	for currentValue, value := range values {
		_, err := transaction.Exec("update items set position = (?) where id = (?)", currentValue, value)
		if err != nil {
			return err
		}
	}
	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func toggleItem(ID int) (Item, error) {
	var item Item
	err := DB.QueryRow("update items set checked = case when checked = 1 then 0 else 1 end where id = (?) returning id, title, checked", ID).Scan(&item.ID, &item.Title, &item.Checked)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}
