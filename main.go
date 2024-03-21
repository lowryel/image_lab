package main

import (
	"fmt"
	"image"

	_ "image/png"

	"log"
	"os"

	"github.com/fogleman/gg"
	"crypto/rand"
	"encoding/binary"
)

func main(){
	file, err:= os.Open("./images/prophets.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	img, str, err :=image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	randNS := RandomIDGen(0)

	err = gg.SaveJPG(fmt.Sprintf("./images/%v.%s", randNS, str),img, 15) // jpg to jpg ( image size reduced by 73% and with better quality)
	if err != nil{
		log.Fatal(err)
	}
}

func RandomIDGen(id uint32) uint32 {
  binary.Read(rand.Reader, binary.LittleEndian, &id)

  return id
}

