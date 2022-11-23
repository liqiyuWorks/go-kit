package base

/**
 * @description: 检查一个字符串是否存在于一个数组中
 * @param {string} target
 * @param {[]string} strList
 * @return {*}
 * @author: liqiyuWorks
 */
func In(target string, strList []string) bool {
	for _, element := range strList {
		if target == element {
			return true
		}
	}
	return false
}
