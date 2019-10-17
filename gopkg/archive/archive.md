

>Go标准库中的tar包实现了tar格式压缩文件的存取。本包目标是覆盖大多数tar的变种，
包括GNU和BSD生成的tar文件。前面一句是go标准库中所说，其实就是可以使用这个包，
然后可以在windows，linux，mac，unix系统上运行，
达到不需要去知道不同操作系统的tar文件底层实现而可以跨平台使用。

```go
const (  
    // Types  
    TypeReg           = '0'    // 普通文件  
    TypeRegA          = '\x00' // 普通文件  
    TypeLink          = '1'    // 硬连接  
    TypeSymlink       = '2'    // 符号连接，软连接  
    TypeChar          = '3'    // 字符设备节点  
    TypeBlock         = '4'    // 块设备节点  
    TypeDir           = '5'    // 目录  
    TypeFifo          = '6'    // fifo node  
    TypeCont          = '7'    // 保留项  
    TypeXHeader       = 'x'    // 可扩展头部  
    TypeXGlobalHeader = 'g'    // 全局扩展头  
    TypeGNULongName   = 'L'    // Next file has a long name  
    TypeGNULongLink   = 'K'    // Next file symlinks to a file w/ a long name  
    TypeGNUSparse     = 'S'    // 稀疏文件  
)  
```
[硬链接和软连接](https://www.ibm.com/developerworks/cn/linux/l-cn-hardandsymb-links/)

字符设备：就是例如键盘，打印机等这种设备

块设备：就是这种以块单位存储数据，如U盘，SD卡，硬盘等这些设备

[设备节点](https://blog.csdn.net/u013904227/article/details/50483662)

[块设备](https://baike.baidu.com/item/%E5%9D%97%E8%AE%BE%E5%A4%87/2413231?fr=aladdin)

[字符设备](https://baike.baidu.com/item/%E5%AD%97%E7%AC%A6%E8%AE%BE%E5%A4%87)

[稀疏文件](https://blog.csdn.net/cymm_liu/article/details/8760033)


```go
type Header struct {  
   Name       string    // 文件名称  
   Mode       int64     // 文件的权限和模式位  
   Uid        int       // 文件所有者的用户 ID  
   Gid        int       // 文件所有者的组 ID  
   Size       int64     // 文件的字节长度  
   ModTime    time.Time // 文件的修改时间  
   Typeflag   byte      // 文件的类型  
   Linkname   string    // 链接文件的目标名称  
   Uname      string    // 文件所有者的用户名  
   Gname      string    // 文件所有者的组名  
   Devmajor   int64     // 字符设备或块设备的主设备号  
   Devminor   int64     // 字符设备或块设备的次设备号  
   AccessTime time.Time // 文件的访问时间  
   ChangeTime time.Time // 文件的状态更改时间  
}  
```
