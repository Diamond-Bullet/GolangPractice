package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"github.com/bits-and-blooms/bitset"
	"math"
	"math/big"
)

type Filter struct {
	ElemNum     uint64
	BloomSize   uint64 //单位bit
	HashFuncNum uint64
	ErrRate     float64

	bitMap *bitset.BitSet
	keys   map[uint32]bool
}

func NewFilter(elemNum, bloomSize, hashFuncNum uint64, errRate float64) *Filter {
	return &Filter{ElemNum: elemNum, BloomSize: bloomSize, HashFuncNum: hashFuncNum, ErrRate: errRate}
}

// Init 初始化布隆过滤器
func (f *Filter) Init() {
	//分配布隆过滤器位图
	f.bitMap = bitset.New(uint(f.BloomSize))

	//初始化哈希函数
	//是否是类似HMAC-SHA256那种通过改变passphase值形成不同的哈希函数
	f.keys = make(map[uint32]bool)
	for uint64(len(f.keys)) < f.HashFuncNum {
		randNum, err := rand.Int(rand.Reader, new(big.Int).SetUint64(math.MaxUint32))
		if err != nil {
			panic(err)
		}
		f.keys[uint32(randNum.Uint64())] = true
	}
}

func (f *Filter) Add(elem []byte) {
	var buf [4]byte
	for k := range f.keys {
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(elem, buf[:]))
		result := hashResult.Mod(hashResult, big.NewInt(int64(f.BloomSize)))
		//把result对应的bit置1
		f.bitMap.Set(uint(result.Uint64()))
	}
}

// IsContain 判断元素是否在集合里面
func (f *Filter) IsContain(elem []byte) bool {
	var buf [4]byte
	for k := range f.keys {
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(elem, buf[:]))
		result := hashResult.Mod(hashResult, big.NewInt(int64(f.BloomSize)))
		//查询result对应的bit是否被置1
		if !f.bitMap.Test(uint(result.Uint64())) {
			return false
		}
	}
	return true
}

func HMACWithSHA128(seed []byte, key []byte) []byte {
	//hmac512 := hmac.New(sha512.New, key)
	hmac512 := hmac.New(sha1.New, key)
	hmac512.Write(seed)
	return hmac512.Sum(nil)
}
