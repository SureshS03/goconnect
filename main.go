package main
//TODO add more function
//redis the cicd in github action
import (
	"fmt"
	"net/http"
	//"os"
)

func main() {
	//s := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("PASSWORD"), os.Getenv("SSL_MODE"))
	//fmt.Println(s)
	db := NewDB("postgres", "user=postgres dbname=goconnect password=arya sslmode=disable")
	defer db.Close()
	router := NewRouter(db)
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	fmt.Println("server on 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(" main err", err)
	}
}