package main

import "fmt"

type Vertex7 struct {
	Lat, Long float64
}

// まぁだいたいは型推論に任せれば良いしインテリジェーに任せればOK
var m3 = map[string]Vertex7{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m3)
}
