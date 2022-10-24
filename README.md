# golang-url-shortener

> Build a URL Shortener in Go with Redis

## Preivew

```json
{
  "long_url":"https://github.com/szy0syz",
  "short_url":"https://t.yna.app/dyg5s"
}
```

## Notes

```bash
go mod init golang-url-shortener

go get github.com/labstack/echo/v4
```

![image](https://user-images.githubusercontent.com/10555820/197456971-93c107f9-3886-40f5-8949-8123d21c8bb3.png)
### Algorithm For Generating a Short Link 🧮

- 使用BASE58 (一般似乎见过64)

> 发自心内：为啥要用BASE58呢？🤔🤔🤔 这是因为：

- BASE58对人阅读转换后的字符更加友好！
- 它会避免出现：`0` (zero), `O` (capital o), `I` (capital i), `l` ( lower L), `+` (plus), and `/` (slash)
- 原来如此 👍🏻
- 想想也是，真是Get到点子上了

- os.Exit(1)

> 动不动就是知识点回顾 😅
>
> 记得之前写next.js的容器时就用了 signal = 9

| 信号名  | 信号编号 | 产生原因     | 默认处理方式         |
|---------|----------|--------------|----------------------|
| SIGHUP  | 1        | 关闭终端     | 终止                 |
| SIGINT  | 2        | ctrl + c     | 终止                 |
| SIGQUIT | 3        | ctrl + \     | 终止+产生1个转储文件 |
| SIGABRT | 6        | abort()函数  | 终止+转储            |
| SIGPE   | 8        | 算术错误     | 终止                 |
| SIGKILL | 9        | kill -9 pid  | 终止，不可捕获/忽略  |
| SIGUSR1 | 10       | 自定义       | 忽略                 |
| SIGSEGV | 11       | 段错误       | 终止+产生1个转储文件 |
| SIGUSR2 | 12       | 自定义       | 忽略                 |
| SIGALRM | 14       | alarm ()函数 | 终止                 |
| SIGTERM | 15       | kill pid     | 终止（可以被忽略）   |
| SIGCHLD | 17       | 子状态变化   | 忽略                 |
| SIGSTOP | 19       | ctrl + z     | 暂停，不可捕获/忽略  |
