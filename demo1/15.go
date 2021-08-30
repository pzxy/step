package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	t := &AddRequest{
		BaseConfig: BaseConfig{
			Id: "1",
		},
		GwConfig: GwConfig{
			Name: "太难了",
		},
	}
	b, _ := json.Marshal(t)
	target := &AddRequest{}
	err := json.Unmarshal(b, target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*target)
}

type BaseConfig struct {
	Id string `json:"id"`
}
type GwConfig struct {
	Name string `json:"name"`
}

type AddRequest struct {
	BaseConfig BaseConfig `json:"baseConfig"`
	GwConfig   GwConfig   `json:"gwConfig"`
}

// Validate satisfies the Validator interface
func (a AddRequest) Validate() error {
	return nil
}

//UnmarshalJSON implements the Unmarshaler interface for the AddEdgeHubConfigRequest type
func (a *AddRequest) UnmarshalJSON(b []byte) error {
	var ar struct {
		BaseConfig BaseConfig
		GwConfig   GwConfig
	}
	if err := json.Unmarshal(b, &ar); err != nil {
		return err
	}
	*a = AddRequest(ar)

	// validate AddEventRequest DTO
	if err := a.Validate(); err != nil {
		return err
	}
	return nil
}
