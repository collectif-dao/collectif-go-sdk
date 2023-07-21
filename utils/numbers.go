package utils

import "math/big"

func GetAttoFilFromFIL(amount int64) *big.Int {
	fil := big.NewInt(1000000000000000000)
	v := big.NewInt(int64(amount))

	return big.NewInt(1).Mul(v, fil)
}
