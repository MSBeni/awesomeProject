package small_pro

import(
	"bufio"
	"fmt"
	"os"
)
// a function to check the input
func Small_pro(){
	var line int;
	in2 := bufio.NewScanner(os.Stdin)
	in2.Scan()
	if err:= in2.Err(); err!=nil{
		fmt.Println("There is something wrong with your file")
	}
	for in2.Scan() {
		fmt.Println("Scanned Text: ", in2.Text())
		line++
	}
	fmt.Println(line)
}
