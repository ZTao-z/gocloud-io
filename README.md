### 博客地址
本次作业的[博客地址](https://blog.csdn.net/think_A_lot/article/details/83992531)

### 一、概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

### 二、任务
编程 web 应用程序 cloudgo-io。

### 三、实现基本功能

#### 1、支持静态文件服务
访问静态文件目录

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112152229778.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3RoaW5rX0FfbG90,size_16,color_FFFFFF,t_70)

查看css/index.css

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112152451891.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3RoaW5rX0FfbG90,size_16,color_FFFFFF,t_70)

查看images/1.jpg

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112152612156.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3RoaW5rX0FfbG90,size_16,color_FFFFFF,t_70)
#### 2、支持简单 js 访问

访问/getInfo路径，获取并显示文本信息

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181113141130321.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3RoaW5rX0FfbG90,size_16,color_FFFFFF,t_70)
#### 3、提交表单，并输出一个表格

输入表单

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112153413885.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3RoaW5rX0FfbG90,size_16,color_FFFFFF,t_70)

输出表格

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112153609574.png)

#### 4、对 /unknown 给出开发中的提示，返回码 5xx

访问localhost:8080/unknown路径，返回501错误。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20181112153711386.png)

### 四、提高要求

#### 1、分析阅读`gzip`过滤器的源码
##### gzip简介
GZIP最早由Jean-loup Gailly和Mark Adler创建，用于UNⅨ系统的文件压缩。我们在Linux中经常会用到后缀为.gz的文件，它们就是GZIP格式的。现今已经成为Internet 上使用非常普遍的一种数据压缩格式，或者说一种文件格式。
HTTP协议上的GZIP编码是一种用来改进WEB应用程序性能的技术。大流量的WEB站点常常使用GZIP压缩技术来让用户感受更快的速度。这一般是指WWW服务器中安装的一个功能，当有人来访问这个服务器中的网站时，服务器中的这个功能就将网页内容压缩后传输到来访的电脑浏览器中显示出来。一般对纯文本内容可压缩到原大小的40%.这样传输就快了，效果就是你点击网址后会很快的显示出来。当然这也会增加服务器的负载. 一般服务器中都安装有这个功能模块的。
##### 源代码分析
由于这部分内容较多，因此写在了另一个博客上。这是新的[博客地址](https://blog.csdn.net/think_A_lot/article/details/83994987)
#### 2、编写中间件，使得用户可以使用 gb2312 或 gbk 字符编码的浏览器提交表单、显示网页。
这部分我调用了`github.com/axgle/mahonia`库中的函数。这部分的字符格式转换函数已写好，但是由于无法获取请求报文头部的关于字符编码格式的值，而且没有相应的测试环境，因此最终没有实现。
```go
// 位于converter.go
package service

import (
    "github.com/axgle/mahonia"
)
// 字符格式转换函数
func ConvertToString(src, srcCode, tarCode string) string {
    srcCoder := mahonia.NewDecoder(srcCode)
    srcResult := srcCoder.ConvertString(src)
    tarCoder := mahonia.NewDecoder(tarCode)
    _, cdata, _ := tarCoder.Translate([]byte(srcResult), true)
    result := string(cdata)
    return result
}
```
