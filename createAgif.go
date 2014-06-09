package main 

import(
	"fmt"
	"os"
	"image/jpeg"
	"image/gif"
)


func main(){
	op := gif.Options{256,nil,nil}
	gifWriter, err := os.Create("new.gif")
	if err != nil{
		fmt.Println("something went wrong")
	}
	image1, _ := os.Open("image1.jpeg")
	image2, _ := os.Open("image2.jpeg")
	jpeg1, _ := jpeg.Decode(image1)
	jpeg2, _ := jpeg.Decode(image2)
	gif.Encode(gifWriter, jpeg1, &op)
	gif.Encode(gifWriter, jpeg2, &op)

}

