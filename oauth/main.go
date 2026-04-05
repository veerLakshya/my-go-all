package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/openfga/go-sdk/oauth2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	tmpl           = template.Must(template.ParseGlob("templates/*.html"))
	dbMaria        *sql.DB
	mongoCol       *mongo.Collection
	keycloakConfig *oauth2.Config
)

// client secret = uVu1EniNrm4JUam1u3TsIhCHPm95vbsp

func init() {
	var err error

	// 1. Connect to Local MariaDB
	dbMaria, err = sql.Open("mysql", "root:root_password@tcp(127.0.0.1:3306)/oauth_lab")
	if err != nil {
		log.Fatal("error connecting to mariadb: ", err)
	}

	// 2. Connect to Local MongoDB & set up TTL Index
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error connecting to mongodb: ", err)
	}
	mongoCol = client.Database("oauth_lab").Collection("auth_sessions")

	//TTL Index: Auto-dleete session data after 10 minutes
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "create_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(600),
	}

	_, _ = mongoCol.Indexes().CreateOne(ctx, indexModel)

	// 3. Configure KeyClock OAuth2.0
	keycloakConfig = &oauth2.Config{
		ClientID:     "my-go-app",
		ClientSecret: "uVu1EniNrm4JUam1u3TsIhCHPm95vbsp",
		RedirectURL:  "https://localhost:8443/callback",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://localhost:8080/realms/protocol/openid-connect/auth",
			TokenURL:  "https://localhost:8080/realms/protocol/openid-connect/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}

func main() {
	http.HandleFunc("/", handleHome)

	err := http.ListenAndServeTLS(":8443", "certs/localhost+2.pem", "certs/localhost+2-key.pem", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
