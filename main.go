package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/taiimrr/gorouter/internal/database"
)
type apiConfig struct {
	DB *database.Queries
}


func main(){
	fmt.Println("hello world")
	godotenv.Load()
	portStr := os.Getenv("PORT")
	if portStr == ""{
		log.Fatal("No port found")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DBURL not found")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection did not work", err)
	}
	
	
	apiCfg := apiConfig{
		DB:database.New(conn),
	
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.handlerUsersGet)




	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portStr,
	}

	log.Printf("Server runnin on port %v", portStr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portStr)

}