# 安装go语言开发环境

本次实验我所使用的是CentOS7系统

相关文件地址：

[Github项目地址](https://github.com/wywwwwei/ServiceComputingOnCloud/tree/master/HW2_GoConfiguration)

[CSDN博客地址](https://blog.csdn.net/WeiXiaoAssassin/article/details/100863538)

[TOC]

## 1. 安装 VSCode 编辑器

> Visual Studio Code 是一个轻量级但功能强大的源代码编辑器，可在 Windows，macOS 和 Linux 桌面上运行。它内置了对JavaScript，TypeScript和Node.js的支持，并为其他语言（如C ++，C＃，Java，Python，PHP，Go）和运行时（如.NET和Unity）提供了丰富的扩展生态系统。

更多详细安装过程可以通过 VSCode官方教程[Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)查看

我这里只列出CentOS 7上的安装方法

```shell
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'
yum check-update
sudo yum install code
```

安装完成之后，我们可以通过在终端输入 `code` 来打开VSCode软件

![VSCode](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/vscode.png)

## 2. 安装golang

### 2.1 安装

首先我们依照步骤输入`sudo yum install golang`，会发现这是行不通的

![InstallFailed](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/installerror.PNG)

这是是因为golang位于第三方的yum源里面，而不在centos官方yum源里面

所以我们需要先通过

```shell
sudo yum install epel-release
```

安装EPEL源 —— 从[循序0010的csdn博客](https://blog.csdn.net/u011341352/article/details/82943871)中引用

> EPEL (Extra Packages for Enterprise Linux)是基于Fedora的一个项目，为“红帽系”的操作系统提供额外的软件包，适用于RHEL、CentOS和Scientific Linux.
>
> 使用方法：需要安装一个叫”epel-**release**”的软件包，这个软件包会自动配置yum的软件仓库。
>
> 当然你也可以不安装这个包，自己配置软件仓库也是一样的。

所以除了用上面的命令行来配置epel之外，也可通过替换/etc/yum.repos.d/epel.repo实现

```shell
# 使用EPEL源的第二个方法，从这个和上面的命令选一个执行
# 先备份
mv /etc/yum.repos.d/epel.repo /etc/yum.repos.d/epel.repo.backup
mv /etc/yum.repos.d/epel-testing.repo /etc/yum.repos.d/epel-testing.repo.backup
# 下载新repo
wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
```

好了，现在就可以继续实验指导的步骤了

```shell
sudo yum install golang
```

![InstalledGo](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/installedGo.PNG)

然后通过输入`rpm -ql golang |more` 查看安装到哪个目录

![gowhichdir](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/gowhichdir.PNG)

终端输入 `go version` 测试安装

![goversion](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/goversion.PNG)

### 2.2 设置环境变量

1. 创建工作空间

   > [golang中国项目组文档中对工作空间的描述](https://go-zh.org/doc/code.html#工作空间)
   >
   > `go` 工具为公共代码仓库中维护的开源代码而设计。
   >
   > Go代码必须放在**工作空间**内。它其实就是一个目录，其中包含三个子目录：
   >
   > - `src` 目录包含Go的源文件，它们被组织成包（每个目录都对应一个包）
   > - `pkg` 目录包含包对象
   > - `bin` 目录包含可执行命令
   >
   > `go` 工具用于构建源码包，并将其生成的二进制文件安装到 `pkg` 和 `bin` 目录中。
   >
   > `src` 子目录通常包会含多种版本控制的代码仓库（例如Git或Mercurial）， 以此来跟踪一个或多个源码包的开发。

   `GOPATH` 环境变量指定了你的工作空间位置。它或许是你在开发Go代码时， 唯一需要设置的环境变量。

   创建一个工作空间目录

   ```shell
   mkdir $HOME/gowork
   ```

2. 配置的环境变量

   在`$HOME/.profile`或`/etc/profile`（二选一）中添加以下指令

   ```shell
   export GOPATH=$HOME/gowork
   export PATH=$PATH:$GOPATH/bin
   ```

   使用指令`source $HOME/.profile` / `source /etc/profile`执行这些配置

   两种修改方式

   >关于linux中先关profile文件的解释：
   >
   >/etc/profile:此文件为系统的每个用户设置环境信息,当用户第一次登录时,该文件被执行.并从/etc/profile.d目录的配置文件中搜集shell的设置，全局生效，使用　source profile 即可
   >
   >~/.bash_profile:每个用户都可使用该文件输入专用于自己使用的shell信息,当用户登录时,该文件仅仅执行一次!默认情况下,设置一些环境变量,执行用户的.bashrc文件.此文件类似于/etc/profile，也是需要需要重启才会生效，/etc/profile对所有用户生效，~/.bash_profile只对当前用户生效。

   值得注意，如果没有`/etc/profile`中配置全局变量，那么**当用户使用`sudo`运行指令时，它的环境变量是root的而不是当前用户的**。

   具体验证：

   - 终端输入`go env`

     ![goenv_wu](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/goenv_wu.PNG)

   - 终端输入`sudo go env`

     ![goenv_root](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/goenv_root.PNG)

3. 终端输入`go env`检查配置

   ![goenv](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/goenv.PNG)

   找到GOPATH和GOROOT，成功

### 2.3 创建hello world！

**退出当前用户，然后重新登陆**

- 创建源代码目录：

  ```shell
  sudo mkdir $GOPATH/src/github.com/github-user/hello -p
  ```

- 使用 vs code 创建 hello.go

  ```shell
  cd $GOPATH/src/github.com/github-user/hello
  #在该目录中创建名为 hello.go 的文件
  sudo touch hello.go
  #使用VSCode打开
  code hello.go
  ```

  ```go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Printf("hello, world\n")
  }
  ```

- 在终端运行

  ```shell
  go run hello.go
  ```

  ![result](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/result.PNG)

## 3. 安装必要的工具和插件

### 3.1 安装 Git 客户端

go 语言的插件主要在 Github 上，安装 git 客户端是首要工作。

```shell
sudo yum install git
```

![gitversion](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/gitversion.PNG)

### 3.2 安装 go 的一些工具

进入 vscode ，它提示要安装一些工作，但是安装会出错，跟着提示会输出

![extesionerror](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/extensionerror.PNG)

出现问题的可按以下操作进行

1. **下载源代码到本地**

   ```shell
   # 创建文件夹
   mkdir $GOPATH/src/golang.org/x/
   cd $GOPATH/src/golang.org/x/
   # 下载源码
   git clone https://github.com/golang/tools.git
   ```

2. **安装工具包**

   ```shell
   go install golang.org/x/tools/go/buildutil
   ```

   退出 vscode，再进入，按提示安装！

   我在这里安装继续失败，不过提示不是连接不上golang.org，而是permission denied

   所以我修改了一下读写权限

   ```shell
   sudo chmod -R 777 /home/wu/gowork/src
   ```

   ![success](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/success.PNG)

   这样的化还有一个`golint` 工具会安装失败，需要执行

   ```shell
   cd $GOPATH/src/golang.org/x/
   git clone https://github.com/golang/lint.git
   
   git install golang.org/x/lint/golint
   ```

3. **安装运行 hello world**

   ```shell
   go install github.com/github-user/hello
   ```

   运行结果：

   ![RunSucceed](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/runsucceed.PNG)

## 4. 安装与运行 go tour

```shell
go get -u github.com/Go-zh/tour/gotour
```

但是事实证明，现在这个指令也需要科学……

![removed](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/remove.PNG)

重新输入

```shell
#先手动安装依赖包
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
#在工作空间的 bin 目录中创建一个可离线执行的 tour 文件
go get -u github.com/Go-zh/tour/gotour
```

然后就可以通过 `tour` 执行打开`127.0.0.1:3999`

![gotour](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/gotour.PNG)

## 5. 按文档写第一个包，做第一个测试

>  仔细阅读官方文档 [如何使用Go编程](https://go-zh.org/doc/code.html) ，并按文档写第一个包，做第一次测试

1. 编写一个库，并让 `hello` 程序来使用它

   ```shell
   #选择包路径
   mkdir $GOPATH/src/github.com/user/stringutil
   #进入目录
   cd $GOPATH/src/github.com/user/stringutil
   #创建go文件
   touch reverse.go
   #使用VSCode编辑该文件
   code reverse.go
   ```

   ```go
   // Package stringutil reverse.go
   // stringutil 包含有用于处理字符串的工具函数。
   package stringutil
   
   // Reverse 将其实参字符串以符文为单位左右反转。
   func Reverse(s string) string {
   	r := []rune(s)
   	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
   		r[i], r[j] = r[j], r[i]
   	}
   	return string(r)
   }
   ```

   现在用 `go build`/`go build github.com/user/stringutil` 命令来测试该包的**编译**。

   这不会产生输出文件。想要输出的话，必须使用 `go install` 命令，它会将包的对象放到工作空间的 `pkg` 目录中。

   确认 `stringutil` 包构建完毕后，修改原来的 `hello.go` 文件（它位于 `$GOPATH/src/github.com/user/hello`）来，引入 `stringutil` 包，使用包内的函数

   ```shell
   #因为之前的hello文件实际是 github.com/github-user/hello 目录下的，所以需要新建
   mkdir $GOPATH/src/github.com/user
   #进入目录
   cd $GOPATH/src/github.com/user
   #创建go文件
   touch hello.go
   #使用VSCode编辑该文件
   code hello.go
   ```

   ```go
   // hello.go
   package main
   
   import (
   	"fmt"
   
   	"github.com/user/stringutil"
   )
   
   func main() {
   	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
   }
   ```

   无论是安装包还是二进制文件，`go` 工具都会安装它所依赖的任何东西。 因此当我们通过

   ```shell
   go install github.com/user/hello
   ```

   来安装 `hello` 程序时，`stringutil` 包也会被自动安装。

   运行结果

   ![runhello](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/runhello.PNG)

   由于我们使用的是fmt.Printf函数，所以不会自动换行。现在我们将其改成fmt.Println

   再次安装运行

   ![RunHelloAgain](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/runhelloagain.PNG)

   确定是自动换行

   它们更多的区别

   >**Println :可以打印出字符串，和变量** 
   >**Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形**

   此时的工作空间

   - src

     ![srcdir](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/srcdir.PNG)

   - pkg

     ![pkgdir](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/pkgdir.PNG)

   - bin

     ![bindir](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/bindir.PNG)

2. 测试

   Go拥有一个轻量级的测试框架，它由 `go test` 命令和 `testing` 包构成。

   你可以通过创建一个名字以 `_test.go` 结尾的，包含名为 `TestXXX` 且签名为 `func (t *testing.T)` 函数的文件来编写测试。

   测试框架会运行每一个这样的函数；若该函数调用了像 `t.Error` 或 `t.Fail` 这样表示失败的函数，此测试即表示失败。

   我们可通过创建文件 `$GOPATH/src/github.com/user/stringutil/reverse_test.go` 来为 `stringutil` 添加测试。

   ```shell
   #进入目录
   cd $GOPATH/src/github.com/user/stringutil
   #创建go文件
   touch reverse_test.go
   #使用VSCode编辑该文件
   code reverse_test.go
   ```

   ```go
   // reverse_test.go
   package stringutil
   
   import "testing"
   
   func TestReverse(t *testing.T) {
   	cases := []struct {
   		in, want string
   	}{
   		{"Hello, world", "dlrow ,olleH"},
   		{"Hello, 世界", "界世 ,olleH"},
   		{"", ""},
   	}
   	for _, c := range cases {
   		got := Reverse(c.in)
   		if got != c.want {
   			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
   		}
   	}
   }
   ```

   接着使用 `go test` / `go test github.com/user/stringutil`运行该测试：

   ![testpack](https://raw.githubusercontent.com/wywwwwei/ServiceComputingOnCloud/master/HW2_GoConfiguration/pics/testpack.PNG)

   若你在包目录下运行 `go` 工具，也可以忽略包路径

   ---

   本次实验结束。

