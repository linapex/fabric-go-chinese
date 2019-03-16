
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:24</date>
//</624456072739950592>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	cryptorand "crypto/rand"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/viper"
)

//等于返回A和B是否相同
type Equals func(a interface{}, b interface{}) bool

var viperLock sync.RWMutex

//包含返回给定切片a是否包含字符串s
func Contains(s string, a []string) bool {
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

//indexinslice返回数组中给定对象o的索引
func IndexInSlice(array interface{}, o interface{}, equals Equals) int {
	arr := reflect.ValueOf(array)
	for i := 0; i < arr.Len(); i++ {
		if equals(arr.Index(i).Interface(), o) {
			return i
		}
	}
	return -1
}

func numbericEqual(a interface{}, b interface{}) bool {
	return a.(int) == b.(int)
}

//GetRandomIndexs返回一段随机索引
//从0到给定的高位
func GetRandomIndices(indiceCount, highestIndex int) []int {
	if highestIndex+1 < indiceCount {
		return nil
	}

	indices := make([]int, 0)
	if highestIndex+1 == indiceCount {
		for i := 0; i < indiceCount; i++ {
			indices = append(indices, i)
		}
		return indices
	}

	for len(indices) < indiceCount {
		n := RandomInt(highestIndex + 1)
		if IndexInSlice(indices, n, numbericEqual) != -1 {
			continue
		}
		indices = append(indices, n)
	}
	return indices
}

//集合是通用的，线程安全的
//集合容器
type Set struct {
	items map[interface{}]struct{}
	lock  *sync.RWMutex
}

//新闻集返回新集
func NewSet() *Set {
	return &Set{lock: &sync.RWMutex{}, items: make(map[interface{}]struct{})}
}

//添加将给定项添加到集合
func (s *Set) Add(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items[item] = struct{}{}
}

//无论给定项是否在集合中，exists返回true
func (s *Set) Exists(item interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, exists := s.items[item]
	return exists
}

//SIZE返回集合的大小
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

//ToArray返回带有项的切片
//在调用该方法时
func (s *Set) ToArray() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	a := make([]interface{}, len(s.items))
	i := 0
	for item := range s.items {
		a[i] = item
		i++
	}
	return a
}

//清除从集合中移除所有元素
func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[interface{}]struct{})
}

//移除从集合中移除给定项
func (s *Set) Remove(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.items, item)
}

//printstacktrace打印到stdout
//所有虎尾鹦鹉
func PrintStackTrace() {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Printf("%s", buf)
}

//如果存在，则getIntorDefault返回config中的int值，否则返回默认值
func GetIntOrDefault(key string, defVal int) int {
	viperLock.RLock()
	defer viperLock.RUnlock()

	if val := viper.GetInt(key); val != 0 {
		return val
	}

	return defVal
}

//如果存在，则getfloat64ordefault从config返回float64值，否则返回默认值。
func GetFloat64OrDefault(key string, defVal float64) float64 {
	viperLock.RLock()
	defer viperLock.RUnlock()

	if val := viper.GetFloat64(key); val != 0 {
		return val
	}

	return defVal
}

//GetDurationOrDefault返回配置中的持续时间值（如果存在），否则返回默认值
func GetDurationOrDefault(key string, defVal time.Duration) time.Duration {
	viperLock.RLock()
	defer viperLock.RUnlock()

	if val := viper.GetDuration(key); val != 0 {
		return val
	}

	return defVal
}

//setval将键值存储到viper
func SetVal(key string, val interface{}) {
	viperLock.Lock()
	defer viperLock.Unlock()
	viper.Set(key, val)
}

//RandomNT以int形式返回[0，n]中的非负伪随机整数。
//如果n<=0，它会恐慌
func RandomInt(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("Got invalid (non positive) value: %d", n))
	}
	m := int(RandomUInt64()) % n
	if m < 0 {
		return n + m
	}
	return m
}

//randomunt64返回随机uint64
func RandomUInt64() uint64 {
	b := make([]byte, 8)
	_, err := io.ReadFull(cryptorand.Reader, b)
	if err == nil {
		n := new(big.Int)
		return n.SetBytes(b).Uint64()
	}
	rand.Seed(rand.Int63())
	return uint64(rand.Int63())
}

func BytesToStrings(bytes [][]byte) []string {
	strings := make([]string, len(bytes))
	for i, b := range bytes {
		strings[i] = string(b)
	}
	return strings
}

func StringsToBytes(strings []string) [][]byte {
	bytes := make([][]byte, len(strings))
	for i, str := range strings {
		bytes[i] = []byte(str)
	}
	return bytes
}

