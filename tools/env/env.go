package env

import (
	"os"
	"strconv"

	"github.com/eos/tools/errs"
)

func GetString(key,defalutValue string) string {
	if v,ok := os.LookupEnv(key);ok {
		return v
	}
	return defalutValue
}
func GetInt(key string,defaultValue int)(int ,error) {
	if v,ok := os.LookupEnv(key);ok {
		ans ,err := strconv.Atoi(v)
		if err != nil {
			return defaultValue,errs.WrapMsg(err,"parse int error")
		}
		return ans,nil
	}
	return defaultValue,nil 
}

func GetFloat64(key string,defaultValue float64) (float64,error) {
	if v,ok := os.LookupEnv(key);ok {
		ans,err := strconv.ParseFloat(v,64)
		if err != nil {
			return defaultValue,errs.WrapMsg(err,"parse float64 error")
		}
		return ans,nil
	}
	return defaultValue,nil
}
func GetBool(key string,defaultValue bool) (bool,error) {
	if v,ok := os.LookupEnv(key);ok {
		ans,err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue,errs.WrapMsg(err,"parse bool fail")
		}
		return ans,nil
	}
	return defaultValue,nil
}

