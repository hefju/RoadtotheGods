//学会func变量, 匿名函数, 字符串的前缀后缀, 文件路径的拼接. 文件夹遍历
package problem1

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
)

//使用filepath.Walk来遍历  filepath.WalkFunc
func Solve1231(srcpath string){
	FindFiles(srcpath,".jpg")
}

//rootpath 是开始文件夹路径  filetype是文件类型的后缀名, 例如".jpg"
func FindFiles(rootpath string,filetype string){
	var fwork filepath.WalkFunc //可以使用匿名函数来处理
	fwork=func(path string,info os.FileInfo,err error)error{
		if info.IsDir(){//不处理文件夹
			return  nil
		}else {//处理文件
			ok:=strings.HasSuffix(info.Name(),filetype)//找到jpg文件
			if ok{
				//url:=filepath.Join(path,info.Name())
				fmt.Println(path)
			}
		}
		return err
	}

	err:=filepath.Walk(rootpath,fwork)

	if err!=nil{
		fmt.Println(err)
	}
}
















