# tcscript

#  自动化跑测试用例的脚本
## 使用方法

### 1. 修改filter.txt中过滤文件列表
 
 ```
 - Example:
 Windows:
    C:\mygo\src\pss\main\initsys\tests\inittests
    C:\mygo\src\pss\main\initsys\tests\testmain
 MacOS:
   /Volumes/data/mygo/src/pss/main/initsys/tests/inittests
   /Volumes/data/mygo/src/pss/main/initsys/tests/testmain
  ``` 
  
  ### 2.  修改run(.bat/.sh)脚本文件
  
 ````
  - Example:
  Windows:
    修改run.bat, 将文件中的此目录
    C:\mygo\src\pss\main\initsys\tests\parktasktests\
    替换为任意目录。
  MacOS:
    修改run.sh, 将文件中的此目录
    /Volumes/data/mygo/src/pss/main/initsys/tests/parktasktests
    替换为任意目录
 ````
 
 ### 3. 运行run(.bat/.sh)等待测试结束
 
 ````
    1. 没有通过的测试用例将写入到 failed_test_cases_outputs.txt
    2. 程序下的buf文件夹记载所有跑过的测试用例输出
 ````
 
 ## 代码覆盖率
  ### 1. html方式查看
  - 直接查看
  ```bash
   go tool cover -html=coverage.out 
  ```
- 生成html文件
 ```bash
  go tool cover -html=coverage.out -o coverage.html
 ```
### 2. 生成文本文件
  
```bash
 go tool cover -func=coverage.out -o coverage.txt
```
