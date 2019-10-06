package pathfile

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	s := `C:/Users/pzxy/Desktop/test`
	filePathDemo(s)
	fmt.Println("**************************")

	filePathDir(s)
}

func filePathDemo(testFilePath string) {
	//第二个入参相当于一个回调函数
	//Walk会将testFilePath路径下的所有目录和文件都传递给Walk函数
	//path是文件或者目录的路径,info是os.FileInfo对象,能获取FileInfo结构体中的属性(Name等)
	//我们自己要 然后判断是否是文件,再对文件进行处理
	err := filepath.Walk(testFilePath, func(path string, info os.FileInfo, err error) error {
		return handleTest(path, info, err)
	})
	if err != nil {
		panic(err)
	}
}

func handleTest(path string, info os.FileInfo, err error) error {
	fmt.Printf("path : %s \ninfo : %v \nerr : %v \n", path, info.Name(), err)
	return nil
}

//这是打印了 路径的所属目录
func filePathDir(s string) {
	fmt.Printf("filePath Dir : %v\n", filepath.Dir(s))
}
