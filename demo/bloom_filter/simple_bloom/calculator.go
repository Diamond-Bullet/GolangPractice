package main

import "math"

// CalBloomSize 计算布隆过滤器位图大小
// elemNum 元素个数
// errorRate 误判率
func CalBloomSize(elemNum uint64, errRate float64) uint64 {
	var bloomBitsSize = float64(elemNum) * math.Log(errRate) / (math.Log(2) * math.Log(2)) * (-1)
	return uint64(math.Ceil(bloomBitsSize))
}

// CalHashFuncNum 计算需要的哈希函数数量
// elemNum 元素个数
// bloomSize 布隆过滤器位图大小，单位bit
func CalHashFuncNum(elemNum, bloomSize uint64) uint64 {
	var k = math.Log(2) * float64(bloomSize) / float64(elemNum)
	return uint64(math.Ceil(k))
}

// CalErrRate 计算布隆过滤器误判率
// elemNum 元素个数
// bloomSize 布隆过滤器位图大小，单位bit
// hashFuncNum 哈希函数个数
func CalErrRate(elemNum, bloomSize, hashFuncNum uint64) float64 {
	var y = float64(elemNum) * float64(hashFuncNum) / float64(bloomSize)
	return math.Pow(float64(1)-math.Pow(math.E, y*float64(-1)), float64(hashFuncNum))
}
