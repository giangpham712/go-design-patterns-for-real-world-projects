package main

func div2(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return
}

func main() {

}
