   package main

   import (
      // local
      "bdutils"

      // go native
      "net/http"
      "log"
      "html/template"
      "fmt"

      // MySQL Driver
      _ "github.com/go-sql-driver/mysql"
   )

   // -------
   // CLASSES
   // -------

   type wallPost struct  {
       Message string
       Date string
   }

   // ---------
   // Functions
   // ---------

   func main() {
       http.HandleFunc("/", DisplayWebsite)
       http.HandleFunc("/jquery", SendJqueryJs)
	   http.HandleFunc("/js", SendJs)
       http.HandleFunc("/style", SendStyle)
       log.Fatal(http.ListenAndServe(":8080", nil))
       

       fmt.Println("La wea esta pronta")
       //http.HandleFunc("/submit", Submit)
   }

   func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	   http.ServeFile(w, r, "js/jquery.js")
   }

   func SendJs(w http.ResponseWriter, r *http.Request) {
	   http.ServeFile(w, r, "js/custom.js")
   }

   func SendStyle(w http.ResponseWriter, r *http.Request) {
	   http.ServeFile(w, r, "css/style.css")
   }

   func DisplayWebsite(w http.ResponseWriter, r *http.Request) {
	   switch r.Method {
	   case "GET":

	   case "POST":
		   r.ParseForm()
		   message := r.FormValue("txtMensaje")
		   date := r.FormValue("txtDate")

         db, err := bdutils.Conectar()

		   fmt.Println("-> Message succesfully written")
	   
	      url := "NULL"

	      //photo := "NULL"
	      
	      sqlStr := "INSERT INTO messages VALUES (NULL, ?, ?, ?, 0);"
		   stmt, _ := db.Prepare(sqlStr)
		   _, err = stmt.Exec(message, date, bdutils.NewNullString(url))
		   
		   if err != nil {
			   panic(err.Error())
		   }

		   // be careful deferring Queries if you are using transactions
		   defer db.Close()
	   default:
		   fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	   }

	   // ------------------------
	   // Recibimos datos de MySQL
	   // ------------------------

	   var message string // Conjunto de variables de tipo wallPost
	   var date string
	   var Wall []wallPost // Variable de tipo wallPost

	   // Abrimos conexion a MySQL
      db, err := bdutils.Conectar()

	   // Por si hay algun error
	   if err != nil {
		   panic(err.Error())
	   }

	   // Recibe datos de MySQL
	   results, err := db.Query("SELECT valueMessages, dateMessages FROM messages ORDER BY dateMessages DESC;")
	   if err != nil {
		   panic(err.Error())
	   }

	   // Recorre los datos de MySQL y los devuelve como una variable data
	   for results.Next() {

		   err = results.Scan(&message, &date)

		   result := wallPost{message, date}

		   Wall = append(Wall, result)

		   if err != nil {
			   panic(err.Error())
		   }
	   }

	   // Cuando termine la funcion Main, cierra la conexion
	   defer db.Close()

	   // ------------------------------------
	   // Usamos el template para enviar datos
	   // ------------------------------------

	   //fmt.Println(Wall)

       t, err := template.ParseFiles("index.html")

       if err != nil {
           log.Print("Hubo un eror mostrando esta pagina: ", err)
       }

       err = t.Execute(w, Wall)

       if err != nil {
           log.Printf("Problema al enviar a HTML ", err)
       }
   }

   /*func randomMessage (randomNumber int) string {
   		var randomMsg string
   		switch randomNumber {
			case 0 :
				randomMsg = "Diego, estás re puto 💞"
			case 1 :
				randomMsg = "Idea robada de Diego"
			case 2 :
				randomMsg = "Copyright Lucius Inc., una subsidiaria de Walt Disney"
		}

	   return randomMsg
   }*/
