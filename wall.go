package main

import (
    "net/http"
    "log"
    "html/template"
    _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
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
	switch r.Method {
	case "GET":

	case "POST":
		r.ParseForm()
		message := r.FormValue("txtMensaje")

		// Abrimos conexion a MySQL
		db, err := sql.Open("mysql", "root@tcp(localhost:3306)/thewall")

		// Por si hay algun error
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("it works :P")

		sqlvalue := fmt.Sprintf("INSERT INTO messages VALUES ( NULL, '%s')", message)
		fmt.Printf(sqlvalue)

		// perform a db.Query insert
		insert, err := db.Query(sqlvalue)

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}

		// be careful deferring Queries if you are using transactions
		defer insert.Close()
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

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

	Wall.Message = Data // Enviamos los mensajes al

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
	r.ParseForm()
	message := r.FormValue("txtMensaje")

	// Abrimos conexion a MySQL
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/thewall")

	// Por si hay algun error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("it works :P")

	sqlvalue := fmt.Sprintf("INSERT INTO messages VALUES ( NULL, '%s')", message)
	fmt.Printf(sqlvalue)

	// perform a db.Query insert
	insert, err := db.Query(sqlvalue)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}