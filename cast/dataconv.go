package cast

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	json "github.com/json-iterator/go"
)

type DataConv struct {
	isLittleEndian bool
}

func (dc *DataConv) StructToByte(s interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if dc.isLittleEndian {
		if err := binary.Write(buf, binary.LittleEndian, s); err != nil {
			fmt.Println("StructEncodeLittleEndian:", err)
			return nil, err
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, s); err != nil {
			fmt.Println("StructEncodeBigEndian:", err)
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (dc *DataConv) GetWord(pb [2]byte, nseq uint) [2]byte {
	var ca [2]byte
	if nseq != 12 { // 高字节在前
		ca[0] = pb[1]
		ca[1] = pb[0]
	} else {
		ca[0] = pb[0]
		ca[1] = pb[1]
	}
	return ca
}

func (dc *DataConv) GetDWord(pb [4]byte, nseq uint) [4]byte {
	var ca [4]byte
	if nseq == 1234 {
		ca[0] = pb[0]
		ca[1] = pb[1]
		ca[2] = pb[2]
		ca[3] = pb[3]
	} else if nseq == 2143 {
		ca[0] = pb[1]
		ca[1] = pb[0]
		ca[2] = pb[3]
		ca[3] = pb[2]
	} else if nseq == 4321 {
		ca[0] = pb[3]
		ca[1] = pb[2]
		ca[2] = pb[1]
		ca[3] = pb[0]
	} else {
		ca[0] = pb[2]
		ca[1] = pb[3]
		ca[2] = pb[0]
		ca[3] = pb[1]
	}
	return ca
}

func (dc *DataConv) GetBoolean(b [1]byte) bool {
	var ca = false
	bytebuff := bytes.NewBuffer(b[:])
	if dc.isLittleEndian {
		binary.Read(bytebuff, binary.LittleEndian, &ca)
	} else {
		binary.Read(bytebuff, binary.BigEndian, &ca)
	}
	return ca
}

func (dc *DataConv) SetBoolean(in bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, in); err != nil {
		fmt.Println("StructEncodeBigEndian:", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dc *DataConv) GetInt8(b [1]byte) int8 {
	return int8(b[0])
}

func (dc *DataConv) SetInt8(in int8) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, in); err != nil {
		fmt.Println("StructEncodeBigEndian:", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dc *DataConv) SetUint8(in uint8) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, in); err != nil {
		fmt.Println("StructEncodeBigEndian:", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dc *DataConv) GetInt16(b [2]byte, nseq uint) int16 {
	var ca int16 = 0
	dw := dc.GetWord(b, nseq)
	bytebuff := bytes.NewBuffer(dw[:])
	if dc.isLittleEndian {
		binary.Read(bytebuff, binary.LittleEndian, &ca)
	} else {
		binary.Read(bytebuff, binary.BigEndian, &ca)
	}
	return ca
}

func (dc *DataConv) SetInt16(in int16, nseq uint) ([]byte, error) {
	buf := new(bytes.Buffer)
	if dc.isLittleEndian {
		if err := binary.Write(buf, binary.LittleEndian, in); err != nil {
			fmt.Println("StructEncodeLittleEndian:", err)
			return nil, err
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, in); err != nil {
			fmt.Println("StructEncodeBigEndian:", err)
			return nil, err
		}
	}
	data := buf.Bytes()
	b := [2]byte{data[0], data[1]}
	result := dc.GetWord(b, nseq)
	return result[:], nil
}

func (dc *DataConv) GetInt32(b [4]byte, nseq uint) int32 {
	var ca int32 = 0
	dw := dc.GetDWord(b, nseq)
	bytebuff := bytes.NewBuffer(dw[:])
	if dc.isLittleEndian {
		binary.Read(bytebuff, binary.LittleEndian, &ca)
	} else {
		binary.Read(bytebuff, binary.BigEndian, &ca)
	}
	return ca
}

func (dc *DataConv) SetInt32(in int32, nseq uint) ([]byte, error) {
	buf := new(bytes.Buffer)
	if dc.isLittleEndian {
		if err := binary.Write(buf, binary.LittleEndian, in); err != nil {
			fmt.Println("StructEncodeLittleEndian:", err)
			return nil, err
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, in); err != nil {
			fmt.Println("StructEncodeBigEndian:", err)
			return nil, err
		}
	}
	data := buf.Bytes()
	b := [4]byte{data[0], data[1], data[2], data[3]}
	result := dc.GetDWord(b, nseq)
	return result[:], nil
}

func (dc *DataConv) GetFloat32(b [4]byte, nseq uint) float32 {
	var ca float32 = 0
	dw := dc.GetDWord(b, nseq)
	bytebuff := bytes.NewBuffer(dw[:])
	if dc.isLittleEndian {
		binary.Read(bytebuff, binary.LittleEndian, &ca)
	} else {
		binary.Read(bytebuff, binary.BigEndian, &ca)
	}
	return ca
}

func (dc *DataConv) SetFloat32(in float32, nseq uint) ([]byte, error) {
	buf := new(bytes.Buffer)
	if dc.isLittleEndian {
		if err := binary.Write(buf, binary.LittleEndian, in); err != nil {
			fmt.Println("StructEncodeLittleEndian:", err)
			return nil, err
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, in); err != nil {
			fmt.Println("StructEncodeBigEndian:", err)
			return nil, err
		}
	}
	data := buf.Bytes()
	b := [4]byte{data[0], data[1], data[2], data[3]}
	result := dc.GetDWord(b, nseq)
	return result[:], nil
}

func (dc *DataConv) GetFloat64(b [8]byte, nseq uint) (ret float64) {
	var ca [8]byte
	if nseq == 1234 {
		ca[0] = b[0]
		ca[1] = b[1]
		ca[2] = b[2]
		ca[3] = b[3]
		ca[4] = b[4]
		ca[5] = b[5]
		ca[6] = b[6]
		ca[7] = b[7]
	} else if nseq == 2143 {
		ca[0] = b[2]
		ca[1] = b[3]
		ca[2] = b[0]
		ca[3] = b[1]
		ca[4] = b[6]
		ca[5] = b[7]
		ca[6] = b[4]
		ca[7] = b[5]
	} else if nseq == 4321 {
		ca[0] = b[6]
		ca[1] = b[7]
		ca[2] = b[4]
		ca[3] = b[5]
		ca[4] = b[2]
		ca[5] = b[3]
		ca[6] = b[0]
		ca[7] = b[1]
	} else {
		ca[0] = b[4]
		ca[1] = b[5]
		ca[2] = b[6]
		ca[3] = b[7]
		ca[4] = b[0]
		ca[5] = b[1]
		ca[6] = b[2]
		ca[7] = b[3]
	}

	byteBuff := bytes.NewBuffer(ca[:])
	if dc.isLittleEndian {
		binary.Read(byteBuff, binary.LittleEndian, &ret)
	} else {
		binary.Read(byteBuff, binary.BigEndian, &ret)
	}
	return ret
}

func (dc *DataConv) SetFloat64(in float64, nseq uint) ([]byte, error) {
	ca := make([]byte, 8)
	buf := new(bytes.Buffer)
	if dc.isLittleEndian {
		if err := binary.Write(buf, binary.LittleEndian, in); err != nil {
			fmt.Println("StructEncodeLittleEndian:", err)
			return nil, err
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, in); err != nil {
			fmt.Println("StructEncodeBigEndian:", err)
			return nil, err
		}
	}
	data := buf.Bytes()

	if nseq == 1234 {
		ca[0] = data[0]
		ca[1] = data[1]
		ca[2] = data[2]
		ca[3] = data[3]
		ca[4] = data[4]
		ca[5] = data[5]
		ca[6] = data[6]
		ca[7] = data[7]
	} else if nseq == 2143 {
		ca[0] = data[2]
		ca[1] = data[3]
		ca[2] = data[0]
		ca[3] = data[1]
		ca[4] = data[6]
		ca[5] = data[7]
		ca[6] = data[4]
		ca[7] = data[5]
	} else if nseq == 4321 {
		ca[0] = data[6]
		ca[1] = data[7]
		ca[2] = data[4]
		ca[3] = data[5]
		ca[4] = data[2]
		ca[5] = data[3]
		ca[6] = data[0]
		ca[7] = data[1]
	} else {
		ca[0] = data[4]
		ca[1] = data[5]
		ca[2] = data[6]
		ca[3] = data[7]
		ca[4] = data[0]
		ca[5] = data[1]
		ca[6] = data[2]
		ca[7] = data[3]
	}
	return ca, nil
}

var DaCv DataConv

// Convert string to other types
func ConvertStringToType(valueType string, value string) (result interface{}, err error) {
	switch valueType {
	case "int":
		return strconv.ParseInt(value, 10, 64)
	case "float":
		return strconv.ParseFloat(value, 32)
	case "double":
		return strconv.ParseFloat(value, 64)
	case "boolean", "bool":
		return strconv.ParseBool(value)
	case "string":
		return value, nil
	default:
		return nil, errors.New("Convert failed")
	}
}

func ConvertInterfaceToString(value interface{}) (result string) {
	typeOfA := reflect.TypeOf(value)
	switch typeOfA.Kind().String() {
	case "int":
		return strconv.Itoa(value.(int))
	case "int64":
		return strconv.FormatInt(value.(int64), 10)
	case "string":
		return value.(string)
	case "float32":
		return strconv.FormatFloat(value.(float64), 'G', -1, 32)
	case "float64":
		return strconv.FormatFloat(value.(float64), 'G', -1, 64)
	default:
		return ""
	}
}

func ConvertByteToType(b []byte, valueType string, nseq uint) (result interface{}, err error) {
	if b == nil || len(b) < 1 {
		return nil, errors.New("Input Data Is nil ")
	}
	switch valueType {
	case "byte":
		return b[0], nil
	case "int":
		da := int(0)
		if len(b) < 4 {
			if len(b) == 1 {
				ca := [1]byte{b[0]}
				da = int(DaCv.GetInt8(ca))
			} else if len(b) == 2 {
				ca := [2]byte{b[0], b[1]}
				da = int(DaCv.GetInt16(ca, nseq/100))
			} else {
				return da, errors.New("Convert Int failed")
			}
		} else {
			ca := [4]byte{b[0], b[1], b[2], b[3]}
			da = int(DaCv.GetInt32(ca, nseq))
		}
		return da, nil
	case "float":
		da := float32(0)
		if len(b) < 4 {
			return da, errors.New("Convert Int failed")
		} else {
			ca := [4]byte{b[0], b[1], b[2], b[3]}
			da = float32(DaCv.GetFloat32(ca, nseq))
		}
		return da, nil
	case "double":
		da := float64(0)
		if len(b) < 8 {
			return da, errors.New("Convert Int failed")
		} else {
			ca := [8]byte{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]}
			da = DaCv.GetFloat64(ca, nseq)
		}
		return da, nil
	case "bool":
		da := false
		ca := [1]byte{b[0]}
		da = DaCv.GetBoolean(ca)
		return da, nil
	case "string":
		da := string(b)
		return da, nil
	default:
		return nil, errors.New("Convert failed")
	}
}

func ConvertTypeToByte(value interface{}, nseq uint) (result []byte, err error) {
	if value == nil {
		return nil, errors.New("input is nil")
	}

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		fmt.Println("StructEncodeBigEndian:", err)
		return nil, err
	}

	switch value := value.(type) {
	case bool:
		return DaCv.SetBoolean(value)
	case int8:
		return DaCv.SetInt8(value)
	case byte:
		return DaCv.SetUint8(value)
	case int16:
		return DaCv.SetInt16(value, nseq/100)
	case int32:
		return DaCv.SetInt32(value, nseq)
	case float32:
		return DaCv.SetFloat32(value, nseq)
	case float64:
		return DaCv.SetFloat64(value, nseq)
	default:
		return nil, errors.New("Unsupported Type")
	}
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	// for k, v := range m {
	//	fmt.Printf("%v: %v\n", k, v)
	// }

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

func ConvertInterfaceToInt(value interface{}) int {
	typeOfA := reflect.TypeOf(value)
	switch typeOfA.Kind().String() {
	case "int":
		return value.(int)
	case "int64":
		return int(value.(int64))
	case "string":
		v, _ := strconv.Atoi(value.(string))
		return v
	case "float32":
		return int(value.(float32))
	case "float64":
		return int(value.(float64))
	default:
		return 0
	}
}

func ConvertInterfaceToType(dataType string, data interface{}) (interface{}, error) {
	switch dataType {
	case "bool":
		return ToBoolE(data)
	case "string":
		return ToStringE(data)
	case "byte":
		r, err := ToUint8E(data)
		return byte(r), err
	case "int":
		return ToInt32E(data)
	case "float":
		return ToFloat32E(data)
	case "double":
		return ToFloat64E(data)
	}
	return nil, errors.New("Unsupported type")
}
