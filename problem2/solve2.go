package problem2

import (
	"io/ioutil"

	"path/filepath"
	"fmt"
	//"sync"
	"os"
	//"github.com/hefju/RoadtotheGods/tool"
	//
	//"sort"
	"encoding/gob"
	"github.com/hefju/RoadtotheGods/tool"
	"sort"
)

type sortelem struct {
	path string
	size int64
}

//type byName []os.FileInfo

//func (f sortelem) Len() int           { return len(f) }
//func (f sortelem) Less(i, j int) bool { return f[i].size < f[j].size }
//func (f sortelem) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

//var wg sync.WaitGroup
func Solve1231(srcpath string){
	fileSize=make(map[string]int64)



	//wg.Add(1)
	WalkFolderOne(srcpath)
	keys:=sliceOfKeys(fileSize)
	for _,v:=range keys  {
		fmt.Println(v,fileSize[v])
	}
	return
fmt.Println(tool.ByteSize(111))
	//wg.Wait()
	WriteToFile()
	elemList:=make([]sortelem,0)
	for k,v:=range fileSize{
		elemList=append(elemList,sortelem{path:k,size:v})
	}
	//sort.Sort(elemList)
	//for k,v:=range elemList  {
	//	fmt.Println(k,tool.ByteSize(v))
	//}

	//index:=0
	// var allcount int64=0
	//for k,v:=range fileSize{
	////	fmt.Println(k,v/1024,"KB")
	//	index++
	//	allcount+=v
	//	fmt.Println(index,k,tool.ByteSize(v))
	//}
	//fmt.Println("合计:",srcpath,tool.ByteSize(allcount))

}
func sliceOfKeys(mymap map[string]int64)[]string{
	newmap:=make(map[int][]string)
	var keys []int
	for k,v := range mymap {
		i32value:=int(v)
		keys = append(keys, int(i32value))
			paths,ok:= newmap[i32value]
			if ok {
			paths=append(paths,k)
		}else {
			tmp:=	make([]string,1)
				tmp[0]=k
			newmap[i32value]=tmp
		}
	}
	sort.Ints(keys)
	list:=make([]string,0)
	for _,intkey:=range keys{
		strdata:=newmap[intkey]
		for _,v:=range strdata{
			list=append(list,v)
		}
	}
	return list
}

func WriteToFile(){

	var filename="text.txt"
	file,err:=os.Create(filename)
	defer file.Close()
	if err!=nil{
		fmt.Println(err)
	}
	enc:=gob.NewEncoder(file)
	err=enc.Encode(fileSize)
	if err!=nil{
		fmt.Println(err)
	}
}


//获取当前文件夹下面的文件, 不会递归子文件夹
func WalkFolderOne(path string) {

	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			folder:=filepath.Join(path,fi.Name())
			fileSize[folder]=0
			FindFiles(folder,folder)
		} else {
			//println(path + "/" + fi.Name())
		}
	}
	//wg.Done()

}
var fileSize map[string]int64

//遍历子文件夹
func FindFiles(rootpath string,countName string){
	var fwork filepath.WalkFunc //可以使用匿名函数来处理
	fwork=func(path string,info os.FileInfo,err error)error{
		if info.IsDir(){//不处理文件夹
			return  nil
		}else {//处理文件
			fileSize[countName]+=info.Size()
		}
		return err
	}

	err:=filepath.Walk(rootpath,fwork)

	if err!=nil{
		fmt.Println(err)
	}
}