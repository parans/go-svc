package ashanti

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {

    shortStr := ""
    sz := -1
    for _, str := range strs {
        if sz == -1 {
            sz = len(str)
            shortStr = str
        } else if len(str) < sz {
            sz = len(str)
            shortStr = str
        }
    }
    longestSubs, temp := "", ""
    shortStrBytes := []byte(shortStr)
    count := 0
    for i:=0 ; i<len(shortStrBytes); i++ {
        temp += string(shortStrBytes[i])
        for _, str := range strs {
            if strings.HasPrefix(str, temp) {
                count++
            } else {
                break
            }
        }
        if count == len(strs) {
            longestSubs = temp
        } else {
            break
        }
        count = 0
    }
    return longestSubs
}
