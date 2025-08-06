package main

import (
	"fmt"
	//"log"
)

func FindMinMax(arr []int)(int ,int,string){
	if len(arr)==0{
		return 0,0,"empty ARRAY"
	}
	min,max :=arr[0],arr[0]
	for _,val:=range arr{
		if val<min{
			min=val
		}
		if val>max{
			max=val
		}
	}
	return min,max,""
}

func SortArray(arr []int)(bool,string){
	if len(arr)<1{
		fmt.Println("array is sorted")
	}
	issorted:=true

	for i:=1;i<len(arr);i++{
		if arr[i]<arr[i-1] {
			issorted= false
		}
	}
	return issorted,""
}

func reverseArr(arr []int)([]int,string){
	if len(arr)<=1{
		return arr,""
	}
	n:=len(arr)
	tempArr:=make([]int,n)
	
	for i:=0;i<len(arr);i++{
		tempArr[i]=arr[n-1-i]
	}
	return tempArr,""
}
func main() {

	//FindMinMax
	// array := []int{3, 5, 1, 8, -2, 7}
	// min, max, err :=FindMinMax(array)
	// if err != "" {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Minimum value: %d\n", min)
	// fmt.Printf("Maximum value: %d\n", max)

	//check array is sorted or not
    // sortArray:=[]int{1,2,3,6,4,5,6}
	// isSorted,err:=SortArray(sortArray)
	// if err!=""{
	// 	fmt.Println("error occured")
	// }
	// if isSorted{
	// 	fmt.Println("array is sorted")
	// }else{
	// 	fmt.Println("array is not sorted")
	// }

	//Reverse an Array
	Arr:=[]int{1,2,3,4,5}
	arr,err:=reverseArr(Arr)
	if  err!=""{
		fmt.Println("error")
	}
	fmt.Println(arr)
	
}