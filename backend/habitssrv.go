package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jazzopaul/habits/habits"
	habits_hello_world "github.com/jazzopaul/habits/habits/hello_world"
	"github.com/jazzopaul/habits/hello_world"
	"github.com/pseidemann/finish"
)

type app struct {
	ctx    context.Context
	cancel context.CancelFunc

	serverID string

	cl     *http.Client
	srv    *http.Server
	router chi.Router

	habitsSvc     *habits.Service
	helloWorldSvc *hello_world.Service

	helloWorldCtrl *habits_hello_world.Controller
}

func runWithCode() int {
	app := &app{}
	app.ctx, app.cancel = context.WithCancel(context.Background())
	defer app.cancel()

	app.cl = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 10,
	}

	app.srv = &http.Server{
		Addr: ":8080",

		ReadTimeout:  time.Minute * 2,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute * 2,
	}
	app.srv.BaseContext = func(_ net.Listener) context.Context {
		return app.ctx
	}

	app.router = chi.NewRouter()

	app.srv.Handler = app.router

	app.habitsSvc = habits.NewService()
	app.helloWorldSvc = hello_world.NewService()

	app.helloWorldCtrl = habits_hello_world.NewController(app.habitsSvc, app.helloWorldSvc)

	fin := &finish.Finisher{
		Timeout: time.Second * 10,
	}

	app.habitsSvc.RegisterPublicController()

	app.habitsSvc.RegisterProtectedController(app.helloWorldCtrl)

	err := app.habitsSvc.Mount(app.ctx, app.router, app.serverID)
	if err != nil {
		log.Println(err)
		return 1
	}

	go func() {
		log.Println("starting webserver")
		err := app.srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	fin.Wait()
	return 0
}
