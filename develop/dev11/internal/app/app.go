package app

import (
	"log"
	"net/http"

	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/server"
	service "github.com/Svcorp77777/Wildberries-L2/dev11/internal/service/calendar"
	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/transport"
	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/transport/rest"
)

func Run(port string) {
	router := http.NewServeMux()
	server := new(server.Server)

	// TODO! Init Service
	calendarService := service.ServiceCalendar()

	// TODO! Init Hadlers
	calculatorHandler := rest.HandlerCalendar(calendarService)

	// TODO! Call Handlers
	calculatorHandler.Init(router)

	// TODO! Check Middleware
	allowHosts := []string{}
	loggedRouter := transport.LoggingRequest(allowHosts, router)

	// TODO! Server Run
	if err := server.Run(port, loggedRouter); err != nil {
		log.Fatalf("ошибка при запуске сервера: %v", err)
		return
	}
}
