# Introduction

这是XXXX开发文档。

## 关于编辑XXXX文档说明：

1. 如何编写[SUMMARY.md](SUMMARY.md)文件 ？

   首先在 SUMMARY.md 文档中添加文档名称，然后运行 `run.sh` 文档自动创建。

   参考如下例子：
    ```bash
   # SUMMARY.md 目录：
    * [1. tedge边缘网关](/tedge/README.md)
        * [1.1 准备工作](/tedge/ready.md)
        * [1.2 网关安装](/tedge/instal.md)
        * [1.3 网关激活](/tedge/actived.md)
        * [1.4 工作目录](/tedge/struct.md)
    * [2. sdk](/sdk/README.md)
        * [2.1 配置文件](/sdk/config.md)
        * [2.2 接口信息](/sdk/interface.md)
        * [2.3 开发步骤](/sdk/develop.md)
    ```
   
   ```bash
   # 对应的目录结构：
   项目目录
    ├── README.md
    ├── SUMMARY.md
    ├── tedge
    │   ├── README.md
    │   ├── actived.md
    │   ├── instal.md
    │   ├── ready.md
    │   └── struct.md
    └── sdk
        ├── README.md
        ├── config.md
        ├── develop.md
        └── interface.md
   ```

2. 如何在本地预览编辑 ？

   ```bash
   chmod +x run.sh
   ./run.sh 
   # 示例1 默认端口启动: ./run.sh
   # 示例2 指定端口启动: ./run.sh 9527
   ```

   访问地址 `localhost:port`，默认端口 4000。如果执行 `run.sh` 不成功，请多尝试几次。

3. 如何同步到 [tuya-gitbook](xx) ？

   将修改分支合并到`develop`分支，文档会自动部署。