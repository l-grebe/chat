# chat

Use ChatGPT to query information in the terminal!

**Note**: The `main` branch may be in an *unstable or even broken state* during development.
Please use [releases](https://github.com/l-grebe/chat/releases) instead of the `main` branch in order to get a stable
set of binaries.

## Features：

* It is written in `go language`
* It calls the `gpt-3.5-turbo` model of the `OpenAI` company, so you need to go
  to [openai.com](https://platform.openai.com/account/api-keys)
  Obtain `API secret key` from the website
* Based on [text translation service](https://ai.youdao.com/product-fanyi-text.s) provided by YouDaoZhiYun, it supports
  multi-language translation function
* In chat, it generates results based on the last 10 interactive conversations
* It will automatically save historical chat information, the information is saved in the `~/.chat.context` file, and it
  will automatically save 1000 interactive conversations
* The configuration file it depends on is `~/.chat.ini`, a sample configuration file: [.chat.ini](conf/default.chat.ini)

## Usage:

1. Download the corresponding binary file
2. Copy the file to the directory included in the system `$PATH`, such as: `sudo mv q /usr/local/bin/q`
3. Copy the default configuration file to the specified location, rename it, and modify the content of the file as
   needed
4. Run the program in terminal:

```shell
q Wake Up, Neo!
 ```

## Demos:

Translation turned off:

<img src="static/gif/demo-1.gif" width=800px>

Using the translation function (Simplified Chinese):

<img src="static/gif/demo-2.gif" width=800px>

Using the translation function (Japanese):

<img src="static/gif/demo-3.gif" width=800px>

Interactive mode:

<img src="static/gif/demo-4.gif" width=800px>

## Localization `README.md`'s

Language

* [Chinese](README-zh.md)    