
package idutil

import (
	"github.com/eos/tools/utils/encrypt"
	"github.com/eos/tools/utils/stringutil"
	"github.com/eos/tools/utils/timeutil"
	"math/rand"
	"strconv"
	"time"
)
func GetMsgIDByMD5(sendID string) string {
	t := stringutil.Int64ToString(timeutil.GetCurrentTimestampByNano())
	return encrypt.MD5Encrypt(t + sendID + stringutil.Int64ToString(rand.Int63n(timeutil.GetCurrentTimestampByNano())))
}

func OperationIDGenerator() string {
	return strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10)
}
