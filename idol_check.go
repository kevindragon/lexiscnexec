// 当新部署了IDOL后，使用这个程序来检查一些基本配置是否一致
//
// 1. 自定义分词
// 2. 停用记
// 3. 同义词
// 4. 百分号
// 5. 圆括号
// 6. 中括号
// 7. a股 A股
// 8. １９８９年1期
// 9. 5%
// 10 公司法  分词

package main

import (
	//"fmt"
	"github.com/kevindragon/lexiscnexec/idol"
)

func main() {
	idol.CustomDict()
}
