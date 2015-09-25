package main

import (
	"fmt"
)
func totalboth(i int)int{
	if (i%3==0 && i%5==0 ){
		return 1
	}else{
		return 0
	}
}
func totalone(i int)int {
	if (i%3==0) {
		return 1

	}else if (i%5==0) {
		return 1
	}else {
		return 0
	}
}
func main(){
	var a[1000] int
	var option int
	var j,sum int =0,0
	fmt.Println("Enter 1 if you want to find the sum of numbers divisible by 3 or 5")
	fmt.Println("Enter 2 if you want to find the sum of numbers divisible by 3 and 5")
	fmt.Println("Enter the option of the command")
	fmt.Scanf("%d",&option)
	fmt.Println(" The option number is ",option)
	if (option == 1) {
		for i := 0; i<=1000; i++ {

			new := totalone(i)
			if (new == 1) {
				sum=sum+i
				a[j]=i
				j++
			}else {
				sum=sum+0
			}

		}
		fmt.Println(" the sum is", sum)
	}else if (option ==2) {
		for i := 0; i<=1000; i++ {

			new := totalboth(i)
			if (new == 1) {
				sum=sum+i
				a[j]=i
				j++
			}else {
				sum=sum+0
			}

		}
		fmt.Println(" the sum is", sum)
		for i:=0;i<j;i++  {
				fmt.Println(a[i])
			}
	}else{
				fmt.Println("Error. Wrong option.There is no such option")
		}
}

	//for i:=0;i<j;i++  {
	//	fmt.Println(a[i])
	//}
	//fmt.Println(" the sum is",sum)

