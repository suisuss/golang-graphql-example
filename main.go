package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"github.com/suisuss/golang-graphql-example/app/domain/repository/answer"
	"github.com/suisuss/golang-graphql-example/app/domain/repository/question"
	"github.com/suisuss/golang-graphql-example/app/domain/repository/question_option"
	"github.com/suisuss/golang-graphql-example/app/generated"
	"github.com/suisuss/golang-graphql-example/app/infrastructure/db"
	"github.com/suisuss/golang-graphql-example/app/infrastructure/persistence"
	"github.com/suisuss/golang-graphql-example/app/interfaces"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	var (
		defaultPort = "8080"
		database = os.Getenv("DATABASE")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn := db.OpenDB(database)

	var ansService answer.AnsService
	var questionService question.QuesService
	var questionOptService question_option.OptService

	ansService = persistence.NewAnswer(conn)
	questionService = persistence.NewQuestion(conn)
	questionOptService = persistence.NewQuestionOption(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService:            ansService,
		QuestionService:       questionService,
		QuestionOptionService: questionOptService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
