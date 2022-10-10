幻14 insert `[fn] + [Del]`
F12 insert `[fn] + [j]`

```sh
#! /bin/bash


```

## bash和sh的区别

`sh`遵循POSIX规范：“当某行代码出错时，不继续往下解释”。
`bash`就算出错，也会继续向下执行。

>POSIX表示可移植操作系统接口（Portable Operating System Interface of UNIX，缩写为 POSIX ）。POSIX标准意在期望获得源代码级别的软件可移植性。换句话说，为一个POSIX兼容的操作系统编写的程序，应该可以在任何其它的POSIX操作系统上编译执行。

sh 跟bash的区别，实际上是bash有没开启POSIX模式的区别。

简单说，sh是bash的一种特殊的模式，sh就是开启了POSIX标准的bash， /bin/sh 相当于 `/bin/bash --posix`。

在Linux系统上/bin/sh往往是指向/bin/bash的符号链接

```bash
ln -s /bin/bash /bin/sh
```

### 查看

```bash
cat /etc/shells # 查看系统可使用的shell类型

cat /etc/passwd # 查看当前默认设置，一般在第一行：

ll /bin/sh #查看当前sh状态
```
