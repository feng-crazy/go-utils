package hmap

import (
	"reflect"
	"testing"
)

func TestGetArrayMapValue(t *testing.T) {
	type args struct {
		mapArrayStr string
		key         string
	}
	arg := args{
		mapArrayStr: "[\n  {\n    \"value\": \"SNMPv2\",\n    \"name\": \"SNMPVersion\",\n    \"label\": \"SNMP Version\",\n    \"data_type\": \"array\",\n    \"desc\": \"\",\n    \"options\": [\n      {\n        \"label\": \"SNMPv1\",\n        \"value\": \"SNMPv1\"\n      },\n      {\n        \"label\": \"SNMPv2\",\n        \"value\": \"SNMPv2\"\n      },\n      {\n        \"label\": \"SNMPv3\",\n        \"value\": \"SNMPv3\"\n      }\n    ],\n    \"require\": true\n  },\n  {\n    \"value\": \"10.122.104.80\",\n    \"name\": \"ip\",\n    \"label\": \"IP\",\n    \"data_type\": \"string\",\n    \"desc\": \"\",\n    \"require\": true\n  },\n  {\n    \"value\": 161,\n    \"name\": \"port\",\n    \"label\": \"Port\",\n    \"data_type\": \"number\",\n    \"desc\": \"\",\n    \"require\": true\n  },\n  {\n    \"label\": \"connect\",\n    \"name\": \"connect\",\n    \"data_type\": \"string\",\n    \"value\": \"tcp\"\n  }\n]",
		key:         "connect",
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "test", args: arg, want: "tcp", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArrayMapValue(tt.args.mapArrayStr, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrayMapValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArrayMapValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
