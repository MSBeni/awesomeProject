package scanbyword

import(
	"bufio"
	"fmt"
	"os"
)
// Scanning a file by its words
func Scanbyword(){
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan(){
		fmt.Println(in.Text())
	}
}