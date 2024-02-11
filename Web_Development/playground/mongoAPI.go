package main


import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// album represents data about a record album
type Suggest struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Sugg  string `json:"sugg"`
}


// suggest slice to seed record suggestion data.
var suggests = []Suggest{
    {ID: "1", Fname: "Charles", Lname: "France", Email: "CharlesF@gmail.com",Sugg: "I need another laptop"},
	{ID: "2", Fname: "Tim", Lname: "Turner", Email: "TT@gmail.com",Sugg: "I need another laptop"},
	{ID: "3", Fname: "Fred", Lname: "Charger", Email: "ThunderBolt23@gmail.com",Sugg: "Rotation of Coffee Flavors"},
}

// getAlbums responds with the list of all albums as Json.
func getSuggestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, suggests)
}

//PostSugg adds an suggestion from JSON recieved in the request body.
func postSuggestions(c *gin.Context) {
	var newSugg suggest

	//Call BinfJSON t bind the recieved JSON to NewALumbm
	if err := c.BindJSON(&newSugg); err != nil {
		return
	}
	//Add the new album to the slice.
	suggests = append(suggests, newSugg)
	c.IndentedJSON(http.StatusCreated, newSugg)
}
// getAlbumByID locates the album whose ID value matches the ID
//parameter sent by the client, then returns that album as a response.
func getSuggestionsByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range suggests {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "suggestion not found"})
}



func main(){
	router := gin.Default()
	router.GET("/api/suggests", getSuggestions)
	router.POST("/api/suggests", postSuggestions)
	router.GET("/api/suggests/:id", getSuggestionsByID)

	router.Run("localhost:5500")
}