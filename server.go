package main

import (
	"context"
	"fmt"
	"leaning-graphql/graph"
	"leaning-graphql/models"
	"leaning-graphql/repository"
	"leaning-graphql/services"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := initDB(true)
	productRepo := repository.NewProductRepository(db)
	productSvc := services.NewProductServices(productRepo)
	// set variables in struct Pascal case because not use New instance
	resolver := graph.Resolver{
		ProductSvc: productSvc,
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))
	//e.Use(middleware.GraphQLMiddleware)
	e.POST("/query", echo.WrapHandler(srv))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%v", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
func initDB(IsAutoMigrate bool) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if IsAutoMigrate {
		db.AutoMigrate(models.Product{})
	}
	return db
}
