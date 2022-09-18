package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func demo1() *DataResp {
	fmt.Println("----> demo1")
	resp, err := http.Get("https://solution.tuya.com/api/v2/market/searchList?pageNo=1&pageSize=100")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	dataResp := new(DataResp)
	json.NewDecoder(bytes.NewBuffer(body)).Decode(dataResp)
	return dataResp
}

func demo1Stub() *DataResp {
	fmt.Println("----> demo1Stub")
	return &DataResp{
		Success: true,
		Status:  "demo1Stub",
	}
}

type DataResp struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
}
