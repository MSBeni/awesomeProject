package arrayofarray

import(
	"fmt"
)
// The median of the grades

func MedianOG(){
	var sum float64
	var median [2]float64
	studentsGrade := [2][3] float64{
		{9.5,10,8.75},
		{5.5,6.75,9},
	}
	for i, _:= range studentsGrade{
		for j,_:= range studentsGrade[i]{
			sum += studentsGrade[i][j]
		}
		median[i] = sum/float64(len(studentsGrade[i]))
	}
	//fmt.Printf("the median is %g\n", median)
	fmt.Println(median)
}