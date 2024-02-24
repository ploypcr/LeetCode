package gopractice

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		gasTank := gas[i]
		circuit := 0
		for j := i; circuit != len(gas); {
			if j > len(gas)-1 {
				j = 0
			}
			if gasTank >= cost[j] {
				
				circuit++
                //fmt.Printf("%v\n", circuit)
				nextStation := j + 1
				if j == len(gas)-1 {
					nextStation = 0
				}
				gasTank = gasTank - cost[j] + gas[nextStation]
				j++

			} else {
                
				break
			}
		}
        fmt.Printf("%v\n", circuit)
		if circuit == len(gas) {
			return i
		}else{
            if(circuit == 0){
                i++
            }else{
			    i += circuit
            }
		}

	}
	return -1
}