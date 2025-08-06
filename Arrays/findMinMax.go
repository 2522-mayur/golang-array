package arrays

import (
	"fmt"
	"log"
)

func FindMinMax(arr []int)(int ,int,error){
	if len(arr)==0{
		return nil,nil,fmt.errorf("empty array")
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
	return min,max,nil
}
