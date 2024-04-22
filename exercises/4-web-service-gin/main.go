package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Albun represente data about a record album
type album struct{
  ID     string  `json:"id"`
  Title  string  `json:"title"`
  Artist string  `json:"artist"`
  Price  float64 `json:"price"`
}

// Albun slice to seed record albin data.
var albums = []album{
  {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  router.GET("/albums/:id", getAlbumsByID)

  router.Run("localhost:8080")

}

// getAlbums responds with the list of all albums as JSON
func getAlbums(x *gin.Context){
  x.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an albun from JSON receiverd in the request
func postAlbums(c *gin.Context){

  var newAlbum album
  if err := c.BindJSON(&newAlbum); err != nil { // first declare err var then use it
    return
  }

  albums = append(albums, newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value  matches the id 
// parameter sent by the client, then returns that album as a reaponse.
func getAlbumsByID(c *gin.Context){
  id := c.Param("id")
  for _, a := range albums {
    if a.ID == id {
      c.IndentedJSON(http.StatusOK, a)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message":"algum not found"})
}
