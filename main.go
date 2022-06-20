package main

import "fmt"

func main() {

	var gNodes int32 = 5
	gFrom :=	[]int32{1, 2, 3, 4, 1, 3}
	gTo :=		[]int32{2, 3, 4, 5, 3, 5}
	gWeight :=	[]int32{30, 50, 70, 90, 70, 85}

	travelTo(gNodes, gFrom, gTo, gWeight)

}

func travelTo(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) {
	var a int32 = 0
	var b int32 = 0
	var w int32 = 0
	var w2 int32 = 0
	var j int32 = 0
	tripNum := 0
	var strArr []int32
	var wArr []int32
	var toArr []int32
	var nxtArr []int

	for l := 0; l < len(gFrom); l++ {
		if gFrom[l] == 1 {
			gto := gTo[l]
			toArr = append(toArr, gto)
		}
	}

	for m := 0; m < len(toArr); m++ {
		for n := 0; n < len(gFrom); n++ {
			if gFrom[n] == toArr[m] {
				nxtArr = append(nxtArr, n)
			}
		}
	}

	for k, e := range gFrom {
		if e == 1 {
			a = int32(k)
			strArr = append(strArr, a)
		}
	}

	// gFrom :=	[]int32{1, 2, 3, 4, 1, 3}
	// gTo :=	[]int32{2, 3, 4, 5, 3, 5}
	// gWeight :=	[]int32{30, 50, 70, 90, 70, 85}
	for i := 0; i < len(nxtArr); i++ {
		for k, e := range gTo {
			if gFrom[nxtArr[i]] == e {
				if gFrom[k] == 1 {
					a = int32(k)
				}
			}
		}
		MainLoop:
			for j = 0; j < gNodes; j++ {
				fmt.Println("a in mainloop ", a)
				if tripNum < 1 && a != gNodes {
					w = gWeight[a]
					b = gTo[a]
					tripNum++
					fmt.Println("First if; a, b and w are: ", a, b, w)
					fmt.Println("First if; tripNum is: ", tripNum)
				} else if tripNum > 0 && b != gNodes {
					for k, e := range gFrom {
						if e == b {
							a = int32(k)
							break
						}
					}
					w2 = gWeight[a] - w

					if w2 < 1 {
						w2 = 0
					}

					w = w + w2

					b = gTo[a]

					tripNum++
					fmt.Println("Second if; a, b and w are: ", a, b, w)
					fmt.Println("Second if; tripNum is: ", tripNum)
				} else if tripNum > 0 && b == gNodes {
					w2 = gWeight[a] - w

					if w2 < 1 {
						w2 = 0
					}

					w = w + w2

					fmt.Println("gWeight: ", w)

					wArr = append(wArr, w)

					break MainLoop
				}
			}
		
		b = 0
		w = 0
		w2 = 0
		tripNum = 0
		
	}

	for _, e := range wArr {
		fmt.Println("wArr: ", e)
	}
}