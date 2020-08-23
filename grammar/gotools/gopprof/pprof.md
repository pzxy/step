下载绘图工具

[http://www.graphviz.org/](http://www.graphviz.org/)

只能显示bench性能代码的图片

```bash
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
(pprof) web
```