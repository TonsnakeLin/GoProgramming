package stdtime

import "time"

type unsafeRecoveryController struct {
	timeout time.Time
}

func StdtimeTest1() {
	u := &unsafeRecoveryController{}
	println("u:", u.timeout.GoString())
	if time.Now().After(u.timeout) {
		println("Now is after u.timeout")
	} else {
		println("Now is before u.timeout")
	}
}
