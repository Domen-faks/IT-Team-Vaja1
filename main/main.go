package main

import (
	"IT-Team-Vaja1/API"
	"IT-Team-Vaja1/DB/MariaDB"
	"IT-Team-Vaja1/Logic"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//Kreiramo DB objekt, ga inicializiramo in z njim naredimo objekt Logic (Kreiramo MariaDB ampak ga v Logic vstavimo kot tip DB - interface)
	db := &MariaDB.MariaDB{User: getEnv("root", "root"), Pass: getEnv("root", "root"), IP: getEnv("DBIP", "127.0.0.1"), Port: 3307, Database: "vaja1"}
	err := db.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logic := Logic.NewController(db)

	//Kreiramo naš router objekt
	var router Router
	router.engine = gin.Default()
	router.api = API.NewController(logic)

	//Registriramo HTTP REST API povezave
	err = router.registerRoutes()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Naredimo 2 kanala in enega od njih povežemo na sistemski exit signal
	quit := make(chan os.Signal, 0)
	done := make(chan bool, 0)
	signal.Notify(quit, os.Interrupt)

	//Definiramo port, handler in timeoute za HTTP server
	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//V ločenem threadu čakamo na exit signal in nato izklopimo http server
	go func() {

		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println(err.Error())
		}
		close(done)

	}()

	//V ločenem threadu zaženemo HTTP server
	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
		os.Exit(1)

	}()

	//Čakamo na konec izvajanja. Vsi deli programa so sedaj zagnani v ločenih threadih.
	//Dalje od tod pridemo samo če program izklopimo ali se zruži
	<-done

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
