package stdprint

import "fmt"

func TestPrintInt() {
	regionID := (uint64)(11234568)
	fmt.Println("string:", fmt.Sprintf("%20d0", regionID), "[]byte:", fmt.Sprintf("%20d0", regionID))
}
