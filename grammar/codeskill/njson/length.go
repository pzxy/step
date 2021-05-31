package main

import "fmt"

func main() {
	s := "{BaseRequest:{Versionable:{ApiVersion:v2} RequestId:e6ea12bd-6982-467e-a646-b898b72a8c9e} Event:{Versionable:{ApiVersion:v2} Id:d7b3ca19-db63-4c00-bf97-fbdcf808b768 DeviceName:7eecacf382224599938f310a98f4d6b0 ProfileName:pdcsdhlsw2yd0a7n Created:1619100299976 Origin:1619100299976 Readings:[{Versionable:{ApiVersion:v2} Id:3e129fec-6c36-414c-bc66-9f9ecf1e3649 Created:1619100299975 Origin:1619100299 DeviceName:7eecacf382224599938f310a98f4d6b0 ResourceName:15 ProfileName:pdcsdhlsw2yd0a7n ValueType:Int16 BinaryReading:{BinaryValue:[] MediaType:} SimpleReading:{Value:0}}] Tags:map[]}}"
	b := []byte(s)
	fmt.Println(len(b))
	fmt.Println(b)
}
