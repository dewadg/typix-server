package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dewadg/typix-server/internal"
	"github.com/dewadg/typix-server/internal/handlers/gql"
	"github.com/dewadg/typix-server/internal/repositories"
	"github.com/dewadg/typix-server/internal/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

func ServeHTTP(name string, version string, client *mongo.Client) *cobra.Command {
	mongoDB := client.Database(os.Getenv("MONGO_DATABASE"))

	mongoUserRepository := repositories.NewMongoUserRepository(mongoDB)
	mongoRaceRepository := repositories.NewMongoRaceRepository(mongoDB)

	userService := services.NewUserService(mongoUserRepository)
	raceService := services.NewRaceService(mongoRaceRepository, mongoUserRepository)

	container := internal.Container{
		MongoUserRepository: mongoUserRepository,
		MongoRaceRepository: mongoRaceRepository,
		UserService:         userService,
		RaceService:         raceService,
	}

	return &cobra.Command{
		Use:   "http:serve",
		Short: "Run HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			router := chi.NewRouter()

			cors := cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodHead},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
				AllowCredentials: true,
			})

			router.Use(middleware.RequestID)
			router.Use(middleware.RealIP)
			router.Use(middleware.Logger)
			router.Use(middleware.Recoverer)
			router.Use(render.SetContentType(render.ContentTypeJSON))
			router.Use(cors.Handler)

			router.Get("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				payload := map[string]interface{}{
					"name":    name,
					"version": version,
				}
				response, _ := json.Marshal(payload)
				writer.Write(response)
			}))

			gqlHandler, err := gql.NewGQLHandler(&container)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			router.Mount("/graphql", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				gqlHandler.ServeHTTP(writer, request)
			}))

			port := os.Getenv("APP_PORT")
			if port == "" {
				port = "8000"
			}

			host := ""
			if os.Getenv("APP_ENV") == "local" {
				host = "localhost"
			}

			fmt.Printf("App running on port %s\n", port)
			http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
		},
	}
}
