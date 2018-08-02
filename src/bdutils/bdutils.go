package bdutils

import (
   "database/sql"
)

func Conectar() (*sql.DB, error) {
		// Abrimos conexion a MySQL
		db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/theWall")

		// Por si hay algun error
		if err != nil {
			panic(err.Error())
		}
		
		return db, err;
}

func NewNullString(s string) sql.NullString {
   if len(s) == 0 {
     return sql.NullString{}
   }
   return sql.NullString{
      String: s,
      Valid: true,
   }
}
