module github.com/yodstar/dueros-dpl

go 1.14

require (
	github.com/dueros/bot-sdk-go v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	github.com/zmrnet/dueros-dpl v0.0.0-20200906143408-736d3f135c37
)

replace (
	github.com/dueros/bot-sdk-go => github.com/yodstar/bot-sdk-go v0.0.0-20200905160449-28d9faad77a8
	github.com/yodstar/dueros-dpl => ./
)
