package clock

import (
	"fmt"
	"testing"
	"time"

	json "github.com/json-iterator/go"
)

type Order struct {
	OrderId    string    `json:"OrderId"`
	CreateTime MilliTime `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

func TestJsonMarshal(t *testing.T) {

	order := Order{
		OrderId:    "10001",
		CreateTime: MilliTime(time.Now()),
		UpdateTime: time.Now(),
	}

	orderBytes, err := json.Marshal(&order)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(string(orderBytes))
	}

}

func TestJsonUnMarshal(t *testing.T) {

	order := Order{}
	j := `{"OrderId":"10001","CreateTime":1625556709475,"UpdateTime":"2021-07-06 08:12:45"}`

	err := json.Unmarshal([]byte(j), &order)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(order)

	fmt.Println(order.CreateTime)
	fmt.Println(order.UpdateTime)
	fmt.Println(time.Time(order.CreateTime))
}
