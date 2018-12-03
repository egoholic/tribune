package handlers

import (
	"context"
	"net/http"
	"time"

	pr "github.com/egoholic/tribune/space/particles/publication/repository"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func RenderTheLatestPublished(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient("mongo://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database("blog").Collection("publications")
	repo := pr.Make(collection)
	repo.LatestPublished()
}

func RenderAllWithPagination(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient("mongo://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database("blog").Collection("publications")

	repo := pr.Make(collection)
	repo.LatestPublished10()
}

func RenderOneBySlug(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient("mongo://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database("blog").Collection("publications")

	repo := pr.Make(collection)
	repo.BySlug("slug") // TODO: add context and url/query string params to handler signatures.
}
