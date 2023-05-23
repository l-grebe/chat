# chat

在终端中使用ChatGPT查询信息！

**注意**: 在开发过程中， `main` 分支可能处于*不稳定甚至断裂状态*。
请使用 [releases](https://github.com/l-grebe/chat/releases) 代替“主”分支，以获得一组稳定的二进制文件。

## 基本介绍：

* 它使用`go`编写
* 它调用的是`OpenAI`公司的`gpt-3.5-turbo`模型，因此需要去 [openai.com](https://platform.openai.com/account/api-keys)
  网站上获取`API secret key`
* 基于有道智云提供的 [文本翻译服务](https://ai.youdao.com/product-fanyi-text.s) ，它支持多种语言翻译功能
* 在交谈中，它会基于最近的10次交互对话产生结果
* 它会自动保存历史聊天信息，信息保存在`~/.chat.context`文件中，会自动保存1000次交互对话
* 它依赖的配置文件是`~/.chat.ini`，配置文件样例: [.chat.ini](conf/default.chat.ini)

## 使用：

1. 下载对应的二进制文件
2. 将文件拷贝到系统`$PATH`包含的目录下，如: `sudo mv q /usr/local/bin/q`
3. 拷贝默认配置文件到指定位置，重命名，并按需修改文件中的内容
4. 在终端中运行程序：

```shell
q Wake Up, Neo!
 ```

## 示例：

关闭翻译功能：

<img src="static/gif/demo-1.gif" width=800px>

使用翻译功能（中文）：

<img src="static/gif/demo-2.gif" width=800px>

使用翻译功能（日文）：

<img src="static/gif/demo-3.gif" width=800px>

交互模式：

<img src="static/gif/demo-4.gif" width=800px>
