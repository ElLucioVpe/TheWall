package main

import (
    "net/http"
    "log"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type wallPost struct  {
    Message string
}

type pageVariables struct {
	Title string
	DBStatus bool
}

func main() {

    http.HandleFunc("/", DisplayWall)
    log.Fatal(http.ListenAndServe(":80", nil))
    //http.HandleFunc("/submit", Submit)
}

func DisplayWall(w http.ResponseWriter, r *http.Request) {
    var Wall wallPost
    var data []string

    log.Printf("xDDDDD")

    p := pageVariables{Title: "pene"}

    // Abrimos conexion a MySQL
    db, err := sql.Open("mysql", "root@tcp(localhost:3306)/thewall")

    p.DBStatus = db.Ping() == nil

    // Por si hay algun error
    if err != nil {
        panic(err.Error())
    }

    t, err := template.ParseFiles("select.html")

    if err != nil {
        log.Print("Hubo un eror mostrando esta pagina: ", err)
    }

    results, err := db.Query("SELECT valueMessages FROM messages")

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        var message string

        log.Print("xD")

        err = results.Scan(&message)

        Wall.Message = message
        data = append(data, Wall.Message)
        log.Printf(Wall.Message)

        if err != nil {
            panic(err.Error())
        }
    }

    err = t.Execute(w, p)
    if err != nil {
        log.Printf("Problema al enviar a HTML ", err)
    }

    // Cuando termine la funcion Main, cierra la conexion
    defer db.Close()
}