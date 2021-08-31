//导入包的时候可以直接导入标准库中的包
//也可以进行远程导入"import"github./spf13/viper/com""
//如果有两个包的名字相同时可以在导入时在导入包前面加入字段给包进行从命名
//导入的包不可以不使用
//可以在包前面加上"_"来导入这个包并且不使用，不过在执行main函数之前会调用导入包的所有init函数

package main

import (
	"code-master/chapter3/words"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text", count)

}
