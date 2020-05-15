package tmp

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ChargeRobots struct {
	chargeRobots []ChargeRobot
}

type ChargeRobot struct {
	DeviceId                int32
	RobotTaskTypeRelationId string
}

func TestSliceMarshal(t *testing.T) {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed, err:%v", err)
		return
	}

	fmt.Printf("data to marshal : %s\n", string(data))

	var unmarshalS []map[string]interface{}
	err2 := json.Unmarshal(data, &unmarshalS)
	if err2 != nil {
		fmt.Printf("json.marshal failed, err:%v", err)
		return
	}
	fmt.Printf("unmarshal data : %v\n", unmarshalS)
}

func TestMap(t *testing.T) {
	m := make(map[int]int, 5)
	m[1] = 1
	addInt(m)
	fmt.Println(m)
}

func addInt(m map[int]int) {
	m[2] = 2
}
