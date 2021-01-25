package main

func avg(arg ...float64)(avg_ans float64){
	var sum float64
	if len(arg) == 0 {
		return 0
	}

	for i := 0; i< len(arg); i++ {
		sum += arg[i]
	}
	avg_ans = sum/float64(len(arg))
	return
}

func bubbleSort(arr []int){
	for i:=0; i<len(arr) - 1; i++ {
		for j:=0; j<len(arr) - i - 1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
func main(){
	//println(avg(2,3))
	 var ar []int = []int{5,3,4,1,2,10,11,0}
	bubbleSort(ar)
	for i:=0; i<len(ar); i++ {
		println(ar[i])
	}
}
