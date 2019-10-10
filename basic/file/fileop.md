#文件的操作
## 文件属性
````
os.Stat()获取到fileInfo

包括 :

Name() string       // base name of the file
Size() int64        // length in bytes for regular files; system-dependent for others
Mode() FileMode     // file mode bits
ModTime() time.Time // modification time
IsDir() bool        // abbreviation for Mode().IsDir()
Sys() interface{} 

````
## 1.创建
````
1. os包创建文件
- os.Open 打开文件只读
- os.OpenFile 可以设置权限 读写|创建|每次打开清空 等操作用的比较多
2 . os包创建目录
- os.MkdirAll 设置目录名字和权限
````

## 2.读取复制和追加

### 2.1 读行
```go
var result []string
reader := bufio.NewReader(file) 
for {
    if a, _, c := reader.ReadLine(); c != io.EOF { 
        result = append(result, string(a))
        continue
    }
    break
}

```
####2.2 读字节

```go
 br.ReadByte()
```
####2,3 读字符串(读到换行符)
```go
br.ReadString('\n')
```

####2.4 写入
```go
var file *os.File 
file.Write([]byte(内容))
```
####2.5 复制和追加

````
1 复制是以写入的方式打开文件 
2 追加是追加的方式打开文件
然后执行 file.Write([]byte(内容))
````
## 3. 删除

````
1 os包删除目录
 os.RemoveAll(st.Name()) //st.Name() 相对路径 一般删除目录需要判断 st.isDir
2 os包删除文件
os.Remove()
3 os包清空文件
os.Truncate(*os.File, 0)
````

## 4 .路径拼接

```go
1 path包
filepath := path.Join("file1", fmt.Sprintf("%v", "test.txt"))
//filepath 为 file1\test.txt
```

## 5 .遍历获取所有文件

```go
//testFilePath 路径下的目录下所有目录和文件都放入回调函数中,且循环调用,直到没有目录
// handleTest 中处理所有文件,判断是目录还是文件等,或许一些其他操作
err := filepath.Walk(testFilePath, func(path string, info os.FileInfo, err error) error {
    return handleTest(path, info, err)
})
```
## 6 .抽象语法树

```go
//获取抽象语法树 parser包
astNode, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
if err != nil {
    log(err.Error())
    return
}
	
//遍历语法树 深度优先顺序遍历AST
//ast包
//func(Node) bool 中要进行处理
ast.Inspect(astNode,func(Node) bool)

```
```go
//遍历语法树例子
ast.Inspect(astNode,
    func(node ast.Node) bool {
        if ret, ok := node.(*ast.FuncDecl); ok {
            if strings.HasPrefix(ret.Name.String(), "Test") {
                result = append(result,
                    testItem{
                        dirPath:       filepath.Dir(fileName),
                        fileName:      fileName,
                        testFuncNames: ret.Name.String(),
                    },
                )
            }
            return true
        }
        return true
    })
```
