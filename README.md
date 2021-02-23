# dingding-go

Dingtalk SDK

### 使用

安装 dingding-go

```shell
$ go get github.com/molizz/dingding-go
```

初始化

```go
cfg := dingding.NewConfig(agentID, appKey, appSecret)
atm := dingding.NewDefaultAccessTokenManager()

dd := dingding.New(cfg, atm)
accessToken, err := dd.AccessToken()

// 处理事件订阅
// hub内部已处理 check_url 事件
// 如果你需要处理自定义的事件，请实现 EventProcessor 接口，并调用下方的Registor() 方法注册事件即可。
// 具体也可以参考考 event_check_url.go 文件中处理 check_url 事件的实现。
//
eventHub := dd.EventHub()
eventHub.Register(EventProcessor)

// body := http.Request.Body()
eventHub.Do(body)
```
