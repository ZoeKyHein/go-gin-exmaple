module go-gin-example

go 1.23.0

replace (
	github.com/ZoeKyHein/go-gin-example/middleware => /Users/wang/Desktop/CodeRepo/go-gin-example/middleware
	github.com/ZoeKyHein/pkg/setting => /Users/wang/Desktop/CodeRepo/go-gin-example/pkg/setting
	github.com/ZoekyHein/go-gin-example/conf => /Users/wang/Desktop/CodeRepo/go-gin-example/conf
	github.com/ZoekyHein/go-gin-example/models => /Users/wang/Desktop/CodeRepo/go-gin-example/models
	github.com/ZoekyHein/go-gin-example/routers => /Users/wang/Desktop/CodeRepo/go-gin-example/routers
)

require github.com/go-ini/ini v1.67.0 // indirect
