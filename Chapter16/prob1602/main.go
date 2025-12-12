package main

import (
	"fmt"
	"unsafe"
)

// OrderInfoのサイズは56バイト
type OrderInfo struct {
	OrderCode   rune     // 4 bytes, plus 4 for padding  -> 0+8
	Amount      int      // 8 bytes, no padding  -> 8+8
	OrderNumber uint16   // 2 bytes, plus 6 for padding -> 16+8
	Items       []string // 24 bytes, no padding -> 24+24
	IsReady     bool     // 1 byte, plus 7 for padding -> 48+8 => 56
}

// SmallOrderInfo サイズ40バイト
type SmallOrderInfo struct {
	IsReady     bool     // 1 byte, plus 1 byte of padding -> 0+2
	OrderNumber uint16   // 2 bytes, no padding  -> 2+2
	OrderCode   rune     // 4 bytes, no padding  -> 4+4
	Amount      int      // 8 bytes, no padding  -> 8+8
	Items       []string // 24 bytes, no padding -> 16+24 => 40
}


func main() {
	// 構造体OrderInfoのサイズを出力
	info := OrderInfo{}
	fmt.Println("OrderInfoのサイズ:", unsafe.Sizeof(info))

	// 構造体OrderInfoの各フィールドのオフセットを出力
	fmt.Println(" OrderCodeのオフセット:", unsafe.Offsetof(info.OrderCode))
	fmt.Println(" Amountのオフセット:", unsafe.Offsetof(info.Amount))
	fmt.Println(" OrderNumberのオフセット:", unsafe.Offsetof(info.OrderNumber))
	fmt.Println(" Itemsのオフセット:", unsafe.Offsetof(info.Items))
	fmt.Println(" IsReadyのオフセット:", unsafe.Offsetof(info.IsReady))

	fmt.Println("--------")
	
	info2 := SmallOrderInfo{}
	fmt.Println("SmallOrderInfoのサイズ:", unsafe.Sizeof(info2))

	// 構造体SmallOrderInfoの各フィールドのオフセットを出力
	fmt.Println(" IsReadyのオフセット:", unsafe.Offsetof(info2.IsReady))
	fmt.Println(" OrderNumberのオフセット:", unsafe.Offsetof(info2.OrderNumber))
	fmt.Println(" OrderCodeのオフセット:", unsafe.Offsetof(info2.OrderCode))
	fmt.Println(" Amountのオフセット:", unsafe.Offsetof(info2.Amount))
	fmt.Println(" Itemsのオフセット:", unsafe.Offsetof(info2.Items))

}