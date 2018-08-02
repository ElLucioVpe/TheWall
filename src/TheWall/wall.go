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
       Message []string
   }

   type pageVariables struct {
	   Title string
   }

   // ---------
   // Functions
   // ---------

   func main() {
       http.HandleFunc("/", DisplayWebsite)
       http.HandleFunc("/jquery", SendJqueryJs)
       log.Fatal(http.ListenAndServe(":80", nil))
       

       fmt.Println("La wea esta pronta")
       //http.HandleFunc("/submit", Submit)
   }

   func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
       http.ServeFile(w, r, "js/jquery.js")
   }

   func DisplayWebsite(w http.ResponseWriter, r *http.Request) {
	   switch r.Method {
	   case "GET":

	   case "POST":
		   r.ParseForm()
		   message := r.FormValue("txtMensaje")
		   date := r.FormValue("txtDate")

         db, err := conectarDB.Conectar()

		   fmt.Println("it works :P")
	   
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

	   var Data []string // Conjunto de variables de tipo wallPost
	   var Wall wallPost // Variable de tipo wallPost

	   // Abrimos conexion a MySQL
      db, err := conectarDB.Conectar()

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
