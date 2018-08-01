# TheWall <small>(by lucius✨💞🌷)</small>

## Introduction

Simple website for writing whatever you want publicly on a website everyone can see. Original idea by [elcomix97edu](https://github.com/elcomix97edu/thewall).

Written mainly in cool, flaming, shade-glasses looking [GO](https://golang.org/).


## Notes
   
  * Importing the database
  
  TheWall currently runs on MySQL. Inside the "db" folder there's everything you need to set it up or modify it.
  
  * Connecting GO to MySQL
  
  Here's a template for the 'conectarDB' package:
  
  ~~~~
  package conectarDB

import (
   "database/sql"
)

func Conectar() (*sql.DB, error) {
		// Abrimos conexion a MySQL
		db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/theWall")

		// In case there's some error
		if err != nil {
			panic(err.Error())
		}
		
		return db, err;
}
 ~~~~

## To-Do List

***

v.0.1🐱‍

- [x] Make basic template

- [x] Make the website fetch data from the BD and show it on the wall

- [x] Let users add new posts to the wall

- [x] Insert date with every new post (possibly user too?)

- [ ] Add fade animation for new posts

- [x] Put some nice font ✨

- [ ] Add functionality to add photos/urls

- [ ] Don't allow HTML code

***




