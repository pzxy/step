和Java、PHP等语言不一样，sh的流程控制不可为空，如(以下为PHP流程控制写法)：
```java
<?php
if (isset($_GET["q"])) {
    search(q);
}
else {
    // 不做任何事情
}
```
在sh/bash里可不能这么写，如果else分支没有语句执行，就不要写这个else。