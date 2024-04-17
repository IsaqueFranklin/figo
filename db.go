//Here is the place where you can add your database.
package main

import (
  //"time"
  "fmt"
  "log"
  "os"
  "context"
  "github.com/joho/godotenv"
  //"github.com/gofiber/template/html/v2"
  //"github.com/gofiber/fiber/v2"
  //"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  //"go.mongodb.org/mongo-driver/bson/primitive"
)

func db() {
  
  //Pegando URI do arquivo .env
  if err := godotenv.Load(); err != nil {
		log.Println("Não encontrado aquivo .env no sistema.")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Você precisa criar sua variável 'MONGODB_URI' no arquivo .env.")
	}

  //Conectando ao MongoDB com MongoDriver para Go.
  clientOptions := options.Client().ApplyURI(uri)

  //Context do mongo
  ctx := context.Background()
  
  //Conectando ao cliente mongo.
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
  }

  //Verificando a conexão com o banco de dados e ctx 
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Conectado ao MongoDB.")

  //Definindo database e collection
  collection := client.Database("test").Collection("posts")
  fmt.Println(collection)
}
