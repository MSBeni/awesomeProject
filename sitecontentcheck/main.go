package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main(){
	args := os.Args[1]
	//wordtocheck := strings.ToLower(args)
	in := bufio.NewScanner(os.Stdin)

	rx := regexp.MustCompile(`[^a-z]+`)   // it is designed to check for the patterns

	in.Split(bufio.ScanWords)

	wordsmap := make(map[string] bool)

	for in.Scan(){
		word := strings.ToLower(in.Text())
		finalword := rx.ReplaceAllString(word,"")  /// check weather the word is in the rx structure or not
		if len(finalword)>2 {
			wordsmap[finalword] = true
		}
	}

	if wordsmap[args]{
		fmt.Printf("The word %s exist in this site \n", args)
	}else{
		fmt.Printf("OOOPPSSS, The word %s does not exist in this site \n", args)
	}

	// run curl -s https://www.w3.org/TR/PNG/iso_8859-1.txt | go run main.go fase   to check the code
}