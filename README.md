# hexo_editor
<img src="https://xiamo.oss-accelerate.aliyuncs.com/xiamo/WordPress/2021/12/20211220154336504.png" />
因为编辑Hexo博客的方式实在麻烦，并且后台插件不是很好看，所以决定自己写，于是也就开始了这个项目，顺便在其中学习一下Golang 

### 预览
测试博客地址： 🔗 [Click here](http://demo.hexo.xiamoqwq.com)

博客管理端地址： 🔗 [Click here](http://admin.hexo.xiamoqwq.com)

账号：xiamo

密码：xiamo

文章根目录路径：/data/hexo/source/_posts/

### 大概会有的功能
- 博客编辑 (一些基础功能，暂时已经完成了)
- 插画收集 (已完成)
- 表情 (这里有markdown-emoji和表情图)
- 内容管理 (存放一些静态资源，以及在博客中使用这些静态资源)
- 发布说说
- css/js配置 (因为我的博客中用到了大量的css和一部分js，所以可以配置一个全局css/js，或者单独导入部分css/js到某个博客中)

### 关于启动方式
首次运行项目会产生一个配置文件，补充完整后重新启动，再运行Vue即可

### 配置文件中的内容
`port`  后端运行端口

`hexo_root`  hexo的根目录（此处用于限制可访问路径，比如说填写：'/data/hexo'，那么就只能访问/编辑该目录下的文件

`username`  登陆用户名

`password`  登陆密码

<div class="test" style="white-space: pre;background-color: #4A4453;color: black;border-radius: 4px;color: white;padding-bottom: 20px">
    <span style="color: #AFA8BA;"># API Service</span>
    <span style="color: deepskyblue">server:</span>
      <span style="color: #AFA8BA;"># 后端运行端口</span>
      <span style="color: #C1ADE3;">port:</span> 7070
      <span style="color: #AFA8BA;"># hexo的根目录（此处用于限制可访问路径，比如说填写：'/data/hexo'，那么就只能访问/编辑该目录下的文件</span>
      <span style="color: #C1ADE3;">hexo_root:</span> E:/work/hexo-xiamo/xiamo
      <span style="color: #AFA8BA;"># 登陆用户名</span>
      <span style="color: #C1ADE3;">username:</span> xiamo
      <span style="color: #AFA8BA;"># 登陆密码</span>
      <span style="color: #C1ADE3;">password:</span> 17906193.
    <span style="color: #AFA8BA;"># 阿里云相关配置</span>
    <span style="color: deepskyblue">ali:</span>
      <span style="color: #AFA8BA;"># accessKeyId，例：LTAI5tQvrYdwG7nAy</span>
      accessKeyId: 
      <span style="color: #AFA8BA;"># accessKeySecret，例：54646cmvsF7NPv6BEGVZN</span>
      accessKeySecret: 
      <span style="color: #AFA8BA;"># bucket，例：xiamo</span>
      bucket: 
      <span style="color: #AFA8BA;"># region，例：oss-cn-shenzhen</span>
      region: 
      <span style="color: #AFA8BA;"># oss域名，例：https://alioss.xiamoqwq.com</span>
      ossHost: 
</div>


### 前端传送门
🔗 [Click here](https://github.com/xm17906193/hexo_editor_vue)

### 我的博客
🔗 [Click here](https://qwq.link)

### 关于说说
说说的数据来自我的另一个项目 "树" ,各位如果想要加上这么一个小模块可以到博客编辑器中自取

说说在这个目录下： `/data/hexo/source/shuoshuo/`

###### 数据来源
如果你有已经完成的说说可以直接修改js代码中的接口地址

如果没有的话则可以到 `https://tree.qwq.link` 中发布一条说说，然后从接口返回中取到你的userId

然后copy目录下index.md再将js代码中的userId替换掉即可💦

