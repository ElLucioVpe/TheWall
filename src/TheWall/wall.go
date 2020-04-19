   package main

   import (
      // local
      "bdutils"

      // go native
      "net/http"
      "log"
      "html/template"
      "fmt"
      "math/rand"

      // MySQL Driver
      _ "github.com/go-sql-driver/mysql"
   )

   // -------
   // CLASSES
   // -------

   type wallPost struct  {
       Message string
       Date string
       Photo template.HTML
   }

    type HtmlData struct {
        Posts []wallPost
        RandomFooter string
    }


   // ---------
   // Functions
   // ---------

   func main() {
	   fmt.Println("La wea esta pronta")

       bdutils.Conectar()

       http.HandleFunc("/", DisplayWebsite)
       http.HandleFunc("/jquery", SendJqueryJs)
	   http.HandleFunc("/js", SendJs)
       http.HandleFunc("/style", SendStyle)
       http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
       http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("upload"))))

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
        //r.ParseForm()

        bdutils.Insert(w,r)

        fmt.Println("-> Message succesfully written")
	   default:
		   fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	   }

	   // ------------------------
	   // Recibimos datos de MySQL
	   // ------------------------

	   var message string // Conjunto de variables de tipo wallPost
       var date string
       var photo string
	   var Wall []wallPost // Variable de tipo wallPost

	   // Abrimos conexion a MySQL
      db := bdutils.Conectar()

	   // Recibe datos de MySQL
	   results, err := db.Query("SELECT valueMessages, dateMessages, COALESCE(photoMessages, '') FROM messages ORDER BY dateMessages DESC;")
	   if err != nil {
		   panic(err.Error())
	   }

	   // Recorre los datos de MySQL y los devuelve como una variable data
	   for results.Next() {

           err = results.Scan(&message, &date, &photo)
           
		   result := wallPost{message, date, template.HTML(photo)}
           log.Print("Aaaaaaaaa: ", result.Photo)

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
       
       data := HtmlData{Wall, randomMessage()}

       t, err := template.ParseFiles("index.html")

       if err != nil {
           log.Print("Hubo un eror mostrando esta pagina: ", err)
       }

       err = t.Execute(w, data)

       if err != nil {
           log.Printf("Problema al enviar a HTML ", err)
       }
   }

   func randomMessage() string {
        min := 0
        max := 3
        randomNumber := rand.Intn(max - min) + min;
   		var randomMsg string
   		switch randomNumber {
			case 0 :
				randomMsg = "Diego, estás re puto 💞"
			case 1 :
				randomMsg = "Idea robada de Diego"
			case 2 :
                randomMsg = "Copyright Lucius Inc., una subsidiaria de Walt Disney"
            case 3 :
                randomMsg = "uwu"
            case 4 :
                randomMsg = "El código está en GitHub"
            case 5 :
                randomMsg = "Powered by Go, viejo"
            case 6 :
                randomMsg = "Si algo se rompe es culpa de Rigby 😞"
            case 7 :
                randomMsg = "🤠"
		}
	   return randomMsg
   }