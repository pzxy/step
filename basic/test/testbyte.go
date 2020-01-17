package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//测试按个读字节

func main() {
	demo6()
}

func demo1() {
	var sbyte [1024]byte
	var byte2 []byte
	var byte3 []byte
	byte2 = []byte("你好呀")
	fmt.Printf("byte2数据: %v\n", byte2)
	copy(sbyte[:], byte2[:])
	fmt.Printf("1024数据: %v\n", len(sbyte))
	for _, data := range sbyte {
		if data != 0 {
			byte3 = append(byte3, data)
		} else {
			break
		}
	}
	fmt.Printf("byte3数据: %v\n", byte3)
}

func demo2() {
	for i, v := range bytes.Split([]byte("1000-1"), []byte("-")) {
		if i == 1 && v != nil {
			fmt.Println(fmt.Sprintf("%s", v))
		}
	}
}
func demo3() {
	var s []string
	s = append(s, "我")
	s = append(s, "爱")
	s = append(s, "你")
	b, _ := json.Marshal(s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%d\n", b)
	fmt.Printf("%b\n", b)
}

//1.坐标coordinate_x，coordinate_y ，
func demo4() {
	var y = 1.2
	var x = 3.0
	var num = 1
	for i := 1; i <= 17; i++ {
		fmt.Printf("UPDATE `xinqiuboche_s300_test`.`map_details` SET  `map_base_id` = '1000',"+
			" `coordinate_x` = '%0.1f', `icon_num` = '71', `coordinate_y` = '%0.1f' WHERE serial_number = '1000-%d';\n", x, y, num)
		y += 2.3
		x += 5.4
		num += 1
	}
}

func demo5() {

	y := 0.0
	for i := 35; i <= 51; i++ {
		fmt.Printf("UPDATE `xinqiuboche_s300_second_test`.`map_details` SET `map_base_id` = '1004', `coordinate_x` = '15.3',"+
			" `icon_num` = '71', `coordinate_y` = '%0.2f', `turning_influence` = '1000-%d,1000-%d', `stall_level` = 0, `stall_priority_numbers` = 0,"+
			" `away_aisle` = 0, `north_id` = '1004-%d', `north_away` = 230, `south_id` = '1004-%d', `south_away` = 230, `west_id` = '1004-%d',"+
			" `west_away` = 540, `east_id` = '1004-%d', `east_away` = 540, `orientation` = 0, `extensing_one` = '', `extensing_two` = '',"+
			" `extensing_three` = '', `width` = 5.40, `height` = 2.30, `rota` = 90.00, `relation_point` = NULL, `id_after_rotate` = NULL,"+
			" `elevator_id` = '', `area_id` = NULL, `block_kind` = NULL WHERE  `serial_number` = '1004-%d';\n", y, i-1, i+1,i-1, i+1, i-17,i+17, i)
		y += 2.3
	}

}

func demo6() {

	for i := 18; i <= 34; i++ {
		fmt.Printf("update map_details set orientation = 0,stall_level=0,stall_priority_numbers=0,away_aisle =1 WHERE   serial_number ='1004-%d';\n", i)
	}
}
