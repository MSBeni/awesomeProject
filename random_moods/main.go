package main


import(
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){
	args := os.Args[1:]
	moods := [...]string{
		"Happy", "Not-Bad", "Excellent", "Moddy",
	}
	if len(args)<1{
		fmt.Println("ERROR!!! enter the name")
	}else{
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods))
		fmt.Printf("Your mood seems to be %s\n", moods[n])
	}
}