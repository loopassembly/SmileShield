package main

import "fmt"

func main() {
	var arry=[]int{1,2,3,4,5,6,7,8,9,10}
	var target int =7;
	var result int = binarySearch(arry,target)

  
}

func binarySearch(arr []int, target int )int {
	var left int =0;
	var right int =len(arr)-1;
	var mid int =0;
	for left<=right{
		mid = (left+right)/2;
		if arr[mid]==target{
			return mid;
		}else if arr[mid]<target{
			left=mid+1;
		}else{
			right=mid-1;
		}
	}
	return -1;

}