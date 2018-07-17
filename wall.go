package main

import (
    "net/http"
    "log"
    "html/template"
    _ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type wallPost struct  {
    Message []string
}

type pageVariables struct {
	Title string
}

func main() {
    http.HandleFunc("/", DisplayWebsite)
    http.HandleFunc("/send", SendMessage)
    log.Fatal(http.ListenAndServe(":80", nil))
    //http.HandleFunc("/submit", Submit)
}

func DisplayWebsite(w http.ResponseWriter, r *http.Request) {

	// ------------------------
	// Recibimos datos de MySQL
	// ------------------------

	var Data []string // Conjunto de variables de tipo wallPost
	var Wall wallPost // Variable de tipo wallPost

	// Abrimos conexion a MySQL
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/thewall")

	// Por si hay algun error
	if err != nil {
		panic(err.Error())
	}

	// Recibe datos de MySQL
	results, err := db.Query("SELECT valueMessages FROM messages")
	if err != nil {
		panic(err.Error())
	}

	// Recorre los datos de MySQL y los devuelve como una variable data
	for results.Next() {
		var message string

		err = results.Scan(&message)


		Data = append(Data, message)
		//log.Printf(Wall.Message)

		if err != nil {
			panic(err.Error())
		}
	}

	Wall.Message = Data

	// Cuando termine la funcion Main, cierra la conexion
	defer db.Close()

	// ------------------------------------
	// Usamos el template para enviar datos
	// ------------------------------------

	//p := pageVariables{Title: "pene"}

    t, err := template.ParseFiles("select.html")

    if err != nil {
        log.Print("Hubo un eror mostrando esta pagina: ", err)
    }

    err = t.Execute(w, Wall)

    if err != nil {
        log.Printf("Problema al enviar a HTML ", err)
    }
}

func SendMessage(w http.ResponseWriter, r *http.Request) {

}