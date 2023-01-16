package questions

// PickPrimeN1 MarkIt 筛素数, 素数定义
func PickPrimeN1(n int) []int {
	var res []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// PickPrimeN2 MarkIt 筛素数，埃氏筛
func PickPrimeN2(n int) []int {
	var res []int

	var NonPrime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !NonPrime[i] {
			res = append(res, i)
			for j := 2; j*i <= n; j++ {
				NonPrime[j*i] = true
			}
		}
	}

	return res
}

// PickPrimeN3 MarkIt 筛素数，线性筛
func PickPrimeN3(n int) []int {
	var res []int

	var NonPrime = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !NonPrime[i] {
			res = append(res, i)
		}
		for _, p := range res {
			if p*i > n {
				break
			}
			NonPrime[p*i] = true
			if i%p == 0 {
				break
			}
		}
	}

	return res
}
