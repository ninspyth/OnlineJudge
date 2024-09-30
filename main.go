/*
API Endpoints: /submit -> for submitting the code

	/run -> for running the code with testcase

Example HTTP request when submitting

	```
	POST /submit
	HOST: localhost:8080

	{
		'user_id':'1',     integer
		'problem_id':'3',  integer
		'user_code':"#include<bits/stdc++.h> using namespace std; int main() { cout << "hello world" << endl; return 0;}",  string
		'input':"",        string || NULL
		'language':"C++"   string
	}
	```
*/
package main

import (
	"context"
	"github.com/ninspyth/OnlineJudge/handlers"
	"github.com/ninspyth/OnlineJudge/router"
	"github.com/ninspyth/OnlineJudge/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	GOROUTINES = 2
)

func StartHTTPServer(wg *sync.WaitGroup, shutdown <-chan struct{}) {
	//completion of a goroutine -> decrement the waitgroup counter
	defer wg.Done()

	//Initialize the gin router
	r := router.InitRouter()

	//Initalize http server
	srv := http.Server{
		Addr:    "192.168.29.42:8080",
		Handler: r,
	}

	//Start the server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error Creating the server")
		}
	}()
	//wait for shutdown signal
	<-shutdown
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//shutdown the Server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced to shut down server..")
	}
	log.Println("Server Exiting")
}

func main() {

	//Initalize RabbitMQ server
	err := handlers.InitAmqp()
	if err != nil {
		log.Fatalf("Error initializing Amqp")
	}
	defer handlers.StopAmqp()

	//create a waitgroup for goroutines
	shutdown := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(GOROUTINES)

	//Start Http Server
	go StartHTTPServer(&wg, shutdown)

	go func() {
		defer wg.Done()
		scheduler.StartMultipleWorkers()
		// defer scheduler.StopWorker()
	}()

	//create a channel to wait for signal to stop http server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//block the function till we get signal
	<-quit
	log.Println("Shutdown signal received...")

	//close the shutdown channel
	close(shutdown)

	//wait till all goroutines are done
	wg.Wait()
	log.Println("Application exited")
}
