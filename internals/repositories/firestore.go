package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewFirestoreClient() (*firestore.Client, context.Context) {
	ctx := context.Background()
	option := option.WithCredentialsFile("../../serviceAccount.json")
	client, err := firestore.NewClient(ctx, "projectName", option)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return client, ctx
}
