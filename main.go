package main

import "fmt"

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

type Point struct {
	X, Y int
}

func (p Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)
	// if err := db.InitDbConfig(); err != nil {
	// 	log.Fatal(err)
	// }

	// db.Connect()

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.Run(":9090")

}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// }
