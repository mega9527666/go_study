package random_util

import (
	// "math/rand"
	"math/rand/v2"
	"time"
)

// RandomUtil 可控随机工具
type RandomUtil struct {
	Seed uint64
	r    *rand.Rand
}

// NewRandomUtil 创建随机工具
// seed == 0 时使用当前时间
func NewRandomUtil(seed int64) *RandomUtil {
	var s uint64
	if seed == 0 {
		s = uint64(time.Now().UnixNano())
	} else {
		s = uint64(seed)
	}

	// PCG 需要两个参数（state, stream）
	src := rand.NewPCG(s, s^0x9e3779b97f4a7c15)

	return &RandomUtil{
		Seed: s,
		r:    rand.New(src),
	}
}
func NewRandomUtilAuto() *RandomUtil {
	return NewRandomUtil(0)
}

// NextInt [min, max]（包含 max）
func (ru *RandomUtil) NextInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return ru.r.IntN(max-min+1) + min
}

// NextNumber [min, max)
func (ru *RandomUtil) NextNumber(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + ru.r.Float64()*(max-min)
}

// NextBoolean 随机 true / false
func (ru *RandomUtil) NextBoolean() bool {
	return ru.r.IntN(2) == 1
}

// RandomArr 从数组中随机取 needNum 个（不重复）
func RandomArr[T any](ru *RandomUtil, arr []T, needNum int) []T {
	if needNum <= 0 || len(arr) == 0 {
		return nil
	}

	// 拷贝一份，避免修改原数组
	temp := make([]T, len(arr))
	copy(temp, arr)

	if needNum > len(temp) {
		needNum = len(temp)
	}

	result := make([]T, 0, needNum)

	for i := 0; i < needNum; i++ {
		idx := ru.NextInt(0, len(temp)-1)
		result = append(result, temp[idx])

		// 删除已选元素（保证不重复）
		temp = append(temp[:idx], temp[idx+1:]...)
	}

	return result
}

// RandomItem 随机取一个
func RandomItem[T any](ru *RandomUtil, arr []T) T {
	return RandomArr(ru, arr, 1)[0]
}

// RandomIndexByPercent 按权重随机
func (ru *RandomUtil) RandomIndexByPercent(rateList []int) int {
	if len(rateList) == 0 {
		return -1
	}

	total := 0
	for _, v := range rateList {
		total += v
	}

	r := ru.NextInt(1, total)
	sum := 0

	for i, v := range rateList {
		sum += v
		if r <= sum {
			return i
		}
	}

	return len(rateList) - 1
}
