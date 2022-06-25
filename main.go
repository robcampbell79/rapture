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
	var wArr []int32
	var toArr []int32
	var nxtArr []int

	for l := 0; l < len(gFrom); l++ {
		if gFrom[l] == 1 {
			gto := gTo[l]
			toArr = append(toArr, gto)
		}
	}

	for _, e := range toArr {
		fmt.Println("toArr ", e)
	}

	for m := 0; m < len(toArr); m++ {
		for n := 0; n < len(gFrom); n++ {
			if gFrom[n] == toArr[m] {
				nxtArr = append(nxtArr, n)
			}
		}
	}

	for _, e := range nxtArr {
		fmt.Println("nxtArr ", e)
	}

	// gFrom :=	[]int32{1, 2, 3, 4, 1, 3}
	// gTo :=	[]int32{2, 3, 4, 5, 3, 5}
	// gWeight :=	[]int32{30, 50, 70, 90, 70, 85}
	for i := 0; i < len(nxtArr); i++ {
		for k, e := range gTo {
			if e == gFrom[nxtArr[i]] && gFrom[k] == 1 {
				b = gTo[nxtArr[i]]
				w = gWeight[k]
			} else {
				continue
			}
		}
		MainLoop:
			for j = 0; j < gNodes; j++ {
				if b != gNodes {
					for k, e := range gFrom {
						if e == b {
							a = gFrom[k]
							b = gTo[k]
							w2 = gWeight[k] - w

							if w2 < 1 {
								w2 = 0
							}
		
							w = w + w2
						}
					}
				} else {
					
					wArr = append(wArr, w)

					break MainLoop

				}
			}
		
	}

	for _, e := range wArr {
		fmt.Println("wArr: ", e)
	}
}