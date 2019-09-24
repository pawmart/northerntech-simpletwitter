package server

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/go-openapi/loads"
	"github.com/pawmart/northerntech-simpletwitter/config"
	"github.com/pawmart/northerntech-simpletwitter/internal/storage/mongo"
	"github.com/pawmart/northerntech-simpletwitter/internal/handler"

	"github.com/pawmart/northerntech-simpletwitter/internal/restapi"
	o "github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
)

type cliArgs struct {
	Port int `arg:"-p,help:port to listen to"`
}

var (
	args = &cliArgs{
		Port: 6543,
	}
)

func GetAPIServer() *restapi.Server {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	cfg := config.NewConfig()
	storageCnf := cfg.GetDbConfig()

	s := mongo.NewStorage(storageCnf)
	h := handler.NewHandler(s)

	api := o.NewNortherntechSimpletwitterAPI(swaggerSpec)
	api.GetHealthHandler = o.GetHealthHandlerFunc(h.GetHealth)
	api.GetTweetsHandler = o.GetTweetsHandlerFunc(h.GetTweets)
	api.GetTweetsIDHandler = o.GetTweetsIDHandlerFunc(h.GetTweet)
	api.PostTweetsHandler = o.PostTweetsHandlerFunc(h.CreateTweet)
	api.PatchTweetsHandler = o.PatchTweetsHandlerFunc(h.UpdateTweet)
	api.DeleteTweetsIDHandler = o.DeleteTweetsIDHandlerFunc(h.DeleteTweet)

	server := restapi.NewServer(api)
	server.ConfigureAPI()

	return server
}

func Start() {
	arg.MustParse(args)

	server := GetAPIServer()
	defer server.Shutdown()

	server.Port = args.Port

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

