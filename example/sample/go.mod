module sample

go 1.14

require (
	github.com/dueros/bot-sdk-go v0.0.0-00010101000000-000000000000
	github.com/yodstar/dueros-dpl v0.0.0-20200906143408-736d3f135c37
	github.com/yodstar/goutil/database v0.0.0-20201021154449-3193bd8ba434
	github.com/yodstar/goutil/logger v0.0.0-20201021154449-3193bd8ba434
)

replace (
	github.com/dueros/bot-sdk-go => github.com/yodstar/bot-sdk-go v0.0.0-20200905160449-28d9faad77a8
	github.com/yodstar/dueros-dpl => ./../../
)
