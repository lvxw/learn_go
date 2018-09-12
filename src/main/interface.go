package main

import "fmt"

type  Phone interface {
	call()
}

type XiaomiPhone struct {
	name string
	price float64
}

type HuaweiPhone struct {
	name string
	price float64
}

func (x XiaomiPhone) call()  {
	fmt.Printf("%T\t%v is calling\n",x,x.name)
}

func (x HuaweiPhone) call()  {
	fmt.Printf("%T\t%v is calling...\n",x,x.name)
}

func main() {
	var xioami Phone = XiaomiPhone{"xioamiMAX5", 1599.9}
	xioami.call()

	var huawei Phone = HuaweiPhone{"huaweiPro10", 3588.9}
	huawei.call()
}