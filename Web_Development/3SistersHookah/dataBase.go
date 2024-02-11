package main

import(
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"os"
)
var db *sql.DB



func main() {
	//Capture connection properties
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "12.0.0.1:3306",
		DBName: "hookah",
	}
	//Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected")
	
	menus, err := menusByItem("ATL Hoe")
	if err != nil {
	    log.Fatal(err)
	}
    fmt.Println("Menus found: %v\n", items)

	//Hard-code ID 2 here to test the query
	men, err := menuByID(2)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Menu found: %v\n", men)

	menID, err := addMenu(Menu{
		Item: "",
		Ingredients: "",
		Price:  35.00
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added menu: %v\n", menID)
}

// menusByItems queries for menus that have the specified item name.
func menusByItem(name string) ([]Menu, error) {
    // An items slice to hold data from returned rows.
    var menus []Menu  
    
    rows, err := db.Query("SELECT * FROM menu WHERE item = ?", name)
    if err != nil {
        return nil, fmt.Errorf("menusByItem %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for row.Next() {
        var men menusByItem
        if err := rows.Scan(&men.ID, &men.Item, &men.Ingredients, &men.Price); err != nil {
            return nil, fmt.Errorf("menusByItem %q: %v, name,err")
        }
        menus = append(menus, men)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("menusByItem %q: %v, name, err")
    }
    return menus, nil
}

// albumByID queries for the album with the specified ID.
func menuByID(id int64) (Menu, error) {
	//An menu to hold data from the returned row.
	var men Menu

	row := db.QueryRow("SELECT * FROM menu WHERE id = ?", id)
	if err := row.Scan(&men.ID, &men.Title, &men.Artist, &men.Price); err !=nil {
		if err == sql.ErrNoRows {
			return men, fmt.Errorf("menusByID %d: no such menu", id)
		}
		return men, fmt.Errorf("menusByID %d: %v", id, err)
	}
	return men, nil
}

//addMenus adds the specified menu to the dataase,
//returning the menu ID of the new entry
func addMenu(men Menu) (int64, error) {
	result, err := db.Exec("INSERT INTO menu(items, ingredients, price) VALUES (?,?,?)", men.Item, men.Ingredients, men.Price)
	if err != nill {
		return 0, fmt.Errorf("addMenu: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addMenu: %v", err)
	}
	return id, nil
}