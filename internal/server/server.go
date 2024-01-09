package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/v2/middleware/cors"
)

var (
	addr = flag.String("addr", ":", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Use(logger.New())
	app.Use(cors.New())

	app.get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlersRoomChatwebsocket))
	app.Get("/room/uuid/view/websocket", websocket.New(handlers.RoomViewWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/ssuid/chat/websocket")
	app.Get("/stream/:ssuid/view/websocket")
}
