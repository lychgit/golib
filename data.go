package golib

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	ROUNDING_MODE_FLOOR = 0
	ROUNDING_MODE_ROUND = 1
	ROUNDING_MODE_CELL = 2
)

//判断data是否为空
func Empty(data interface{}) bool {
	if data == nil || data == "" {
		return true
		//} else if  {
		//	// 复合数据类型
		//	return true
	}
	return false
}

func GetStructKey(data map[string]interface{}) string {
	for k, _ := range data {
		return k
	}
	return ""
}

type MapsSort struct {
	Key     string
	MapList []map[string]interface{}
}

func (m *MapsSort) Len() int {
	return len(m.MapList)
}

func (m *MapsSort) Less(i, j int) bool {
	key1 := GetStructKey(m.MapList[i])
	key2 := GetStructKey(m.MapList[j])
	return key1 < key2
}

func (m *MapsSort) Swap(i, j int) {
	m.MapList[i], m.MapList[j] = m.MapList[j], m.MapList[i]
}

// 根据map类型的切片的某个键值排序
func Sort(key string, maps []map[string]interface{}) []map[string]interface{} {
	mapsSort := MapsSort{}
	mapsSort.Key = key
	mapsSort.MapList = maps
	sort.Sort(&mapsSort)
	return mapsSort.MapList
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Double Md5 Encrypt
func DoubleMd5Encrypt(data string) string {
	return Md5(Md5(data))
}

// Sha256 Encrypt
func Sha256Encrypt(key string) string {
	id := fmt.Sprintf(key)
	hash := sha256.New()
	hash.Write([]byte(id))
	sum := sha256.Sum256([]byte(id))
	id = fmt.Sprintf("%x", sum)
	return id
}

// 判断数据是否在一个数组中
// dataType 是value的数据类型 如果value和array切片（或数组）中值的类型不同则函数会返回false
func InArray(array interface{}, value interface{}) bool {
	// 获取value的数据类型
	valueType := reflect.TypeOf(value).String()
	val := reflect.ValueOf(array)
	kind := val.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i:=0; i< val.Len(); i++ {
			_v := val.Index(i)
			// 判断数组中的值和value的值是否相同
			if _v.Kind().String() != valueType {
				break
			}
			if _v.Interface() == value {
				return true
			}
		}
	}
	return false
}

// 将数据的类型强制转换为string类型
func ToString(data interface{}) string{
	var str string
	switch data.(type) {
	case uint:
		str =  strconv.FormatUint(uint64(interface{}(data).(uint)), 10)
		break
	case uint8:
		str =  strconv.FormatUint(uint64(interface{}(data).(uint8)), 10)
		break
	case uint16:
		str =  strconv.FormatUint(uint64(interface{}(data).(uint16)), 10)
		break
	case uint32:
		str =  strconv.FormatUint(uint64(interface{}(data).(uint32)), 10)
		break
	case uint64:
		str =  strconv.FormatUint(uint64(interface{}(data).(uint8)), 10)
		break
	case int:
		str = strconv.Itoa(interface{}(data).(int))
		break
	case int8:
		str = strconv.Itoa(int(interface{}(data).(int8)))
		break
	case int32:
		str = strconv.Itoa(int(interface{}(data).(int32)))
		break
	case int64:
		//str = strconv.Itoa(int(interface{}(data).(int64)))
		str = strconv.FormatInt(interface{}(data).(int64),10)
		break
	case float32:
		str = strconv.FormatFloat(interface{}(data).(float64),'f',7, 32)
		break
	case float64:
		str = strconv.FormatFloat(interface{}(data).(float64),'f',15, 64)
		break
	case string:
		str = interface{}(data).(string)
		break
	default:
		str = interface{}(data).(string)
	}
	return str
}

/**
 * separator	必需。规定在哪里分割字符串。
 * str	必需。要分割的字符串。
 * limit 可选。规定所返回的数组元素的数目。 可能的值：
 * 大于 0 - 返回包含最多 limit 个元素的数组
 * 小于 0 - 返回包含除了最后的 -limit 个元素以外的所有元素的数组
 * 0 - 返回包含一个元素的数组
 */
func Explode(separator string, str string)  map[interface{}]interface{}{
	attr := make(map[interface{}]interface{})
	for k,v := range strings.SplitN(str, separator, -1) {
		attr[k] = v
	}
	return attr
}

// @param str json字符串
// @param data 解析的数据结构 传入data的引用
func JsonDecode(str string, data interface{}) error {
	err := json.Unmarshal([]byte(str), data)
	return err
}

func JsonEncode(data interface{}) string {
	var content []byte
	var err error
	content, err = json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(content)
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}
	return fmt.Sprintf("%.2f %s", size, units[n])
}

func FloatScale(f float64, scale , roundMode int) float64 {
	var _f float64
	multiple := math.Pow10(scale)
	switch roundMode {
	case ROUNDING_MODE_FLOOR:
		// 舍弃
		_f = math.Floor(f * multiple) / multiple
	case ROUNDING_MODE_ROUND:
		// 进位
		_f = math.Round(f * multiple) / multiple
	default:
		_f = f
	}
	return _f
}

// 获取数组中某个字段的值的构成的数组指针
func ArrayColumn(array interface{}, columnKey string) ([] interface{}, error) {
	val := reflect.ValueOf(array)
	if val.Kind() == reflect.Slice {
		datas := make([]interface{}, 0)
		for i := 0; i < val.Len(); i++ {
			item := GetStructColumnValueByColumnKey(val.Index(i).Interface(), columnKey)
			datas = append(datas, item)
		}
		return datas, nil
	} else {
		return nil, errors.New("参数错误, 查找对象并非数组类型")
	}
}

// 获取数组中某个字段的值(string 类型)的构成的数组指针
func ArrayStringColumn(array interface{}, columnKey string) ([] string, error) {
	val := reflect.ValueOf(array)
	if val.Kind() == reflect.Slice {
		datas := make([]string, 0)
		for i := 0; i < val.Len(); i++ {
			item := GetStructColumnValueByColumnKey(val.Index(i).Interface(), columnKey)
			datas = append(datas, item.(string))
		}
		return datas, nil
	} else {
		return nil, errors.New("参数错误, 查找对象并非数组类型")
	}
}

// Get Struct column value by (struct column) key
// data should be a Struct or a (Struct) Ptr
func GetStructColumnValueByColumnKey(data interface{}, columnKey string) interface{} {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr || val.Kind() == reflect.Struct {
		elem := val.Elem() //获取结构体中个字段的值
		for i := 0; i < elem.NumField(); i++ {
			structField := elem.Type().Field(i) //结构体字段对应的值
			if structField.Tag.Get("json") == columnKey || structField.Name == columnKey {
				return elem.FieldByName(structField.Name).Interface()
			}
		}
	}
	return nil
}

// map数组中的值根据需要使用类型断言转换 columnKey需是唯一值不重复
func GetSliceColumnMapArray(data interface{}, columnKey string) (map[string]interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		dMap := make(map[string]interface{})
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			keyValue := GetStructColumnValueByColumnKey(v, columnKey)
			dMap[ToString(keyValue)] = v
		}
		return dMap, nil
	} else {
		return nil, errors.New("参数错误, 查找对象并非数组类型")
	}
}

// 根据数组中某个字段的值分组  返回的是多个数组映射
func GroupingMapDataByColumnKeyValue(data interface{}, columnKey string) (map[string][]interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Map {
		dMap := make(map[string][]interface{},0)
		mapKeys := val.MapKeys()
		for _, k := range mapKeys {
			v := val.MapIndex(k).Interface()
			keyValue := GetStructColumnValueByColumnKey(v, columnKey)
			mk := ToString(keyValue)
			if d, ok := dMap[mk]; ok {
				dMap[mk] = append(d, v)
			} else {
				dMap[mk] = []interface{}{v}
			}
		}
		return dMap, nil
	} else {
		return nil, errors.New("参数错误, 查找对象并非数组类型")
	}
}

func GetKindOfData(data interface{}) string {
	v := reflect.ValueOf(data)
	return v.Kind().String()
}