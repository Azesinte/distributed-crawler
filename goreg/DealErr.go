package DealErr

import (
	"fmt"
	"os"
)

func DealErr(err error, time string) {
	if err != nil {
		fmt.Println(time, err)
		os.Exit(1)
	}
}
