pacakge main

import (
    "context"
    "log"
    "strconv"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Post struct {
    ID    string `json:"id,omitempty" bson:"_id,omitempty"`
    Title string `json:"title,omitempty" bson:"title,omitempty"`
    Body  string `json:"body,omitempty" bson:"body,omitempty"` 
}

func CreatePost(ctx *fiber.Ctx) error {
  post := new(Post)
  if err := ctx.BodyParser(post); err != nil {
    return err
  }

  collection := client.Database("test").Collection("posts")
  _, err := collection.InsertOne(context.Background(), post)

  if err != nil {
    return err
  }

  return ctx.JSON(post)
}

func GetPosts(ctx *fiber.Ctx) error {
  var posts []Post
  collection := client.Database("test").Collection("posts")
  cursor, err := collection.Find(context.Background(), bson.M{})
  
  if err != nil {
    return err
  }

  defer cursor.Close(context.Background())

  for cursor.Next(context.Background()) {
    var post Post
    cursor.Decode(&post)
    posts = append(posts, post)
  }
  return ctx.JSON(posts)
}

func main() {
  app := fiber.New()

  //MongoDB connection

  clientOptions := options.Client().ApplyURI()

  var err error
  client, err = mongo.Connect(context.Background(), clientOptions)
  if err != nil {
    log.Fatal(err)
  }

  //Defer closing the client
  defer client.Disconnect(context.Background())

  //Routes
  app.Post("/posts", CreatePost)
  app.Get("/posts", GetPosts)

  //Serve static files
  app.Static("/", "./public")

  //Start the server
  log.Fatal(":8080")
}

