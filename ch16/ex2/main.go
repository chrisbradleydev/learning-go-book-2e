package main

import (
	"fmt"
	"unsafe"
)

// OrderInfo has a total size of 56 bytes
type OrderInfo struct {
	OrderCode   rune     // 4 bytes, plus 4 for padding
	Amount      int      // 8 bytes, no padding
	OrderNumber uint16   // 2 bytes, plus 6 for padding
	Items       []string // 24 bytes, no padding
	IsReady     bool     // 1 byte, plus 7 for padding
}

// SmallOrderInfo has a total size of 40 bytes
type SmallOrderInfo struct {
	IsReady     bool     // 1 byte, plus 1 byte of padding
	OrderNumber uint16   // 2 bytes, no padding
	OrderCode   rune     // 4 bytes, no padding
	Amount      int      // 8 bytes, no padding
	Items       []string // 24 bytes, no padding
}

func main() {
	// Print the size of the OrderInfo struct
	orderInfo := OrderInfo{}
	fmt.Println("Size of OrderInfo:", unsafe.Sizeof(orderInfo))

	// Print the offsets of each field in the OrderInfo struct
	fmt.Println("Offset of OrderCode:", unsafe.Offsetof(orderInfo.OrderCode))
	fmt.Println("Offset of Amount:", unsafe.Offsetof(orderInfo.Amount))
	fmt.Println("Offset of OrderNumber:", unsafe.Offsetof(orderInfo.OrderNumber))
	fmt.Println("Offset of Items:", unsafe.Offsetof(orderInfo.Items))
	fmt.Println("Offset of IsReady:", unsafe.Offsetof(orderInfo.IsReady))

	// Print the size of the OrderInfo struct
	smallOrderInfo := SmallOrderInfo{}
	fmt.Println("\nSize of SmallOrderInfo:", unsafe.Sizeof(smallOrderInfo))

	// Print the offsets of each field in the OrderInfo struct
	fmt.Println("Offset of OrderCode:", unsafe.Offsetof(smallOrderInfo.OrderCode))
	fmt.Println("Offset of Amount:", unsafe.Offsetof(smallOrderInfo.Amount))
	fmt.Println("Offset of OrderNumber:", unsafe.Offsetof(smallOrderInfo.OrderNumber))
	fmt.Println("Offset of Items:", unsafe.Offsetof(smallOrderInfo.Items))
	fmt.Println("Offset of IsReady:", unsafe.Offsetof(smallOrderInfo.IsReady))
}
