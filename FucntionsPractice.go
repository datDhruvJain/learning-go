package main

func avg(lol ...float64)(avg_ans float64){
	var sum float64
	for i := 0; i< len(lol); i++ {
		sum += lol[i]
	}
	avg_ans = sum/float64(len(lol))
	return
}

func main(){
	println(avg(2,3))
}