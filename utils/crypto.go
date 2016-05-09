/*
* @Author: xuhuan
* @Date:   2016-05-09 15:04:15
* @Last Modified by:   xuhuan
* @Last Modified time: 2016-05-09 15:04:15
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
