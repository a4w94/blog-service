package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 定義限流器介面所需的方法
type LimitIface interface {
	Key(c *gin.Context) string                        //取得對應的限流器的鍵值對名稱
	GetBucket(key string) (*ratelimit.Bucket, bool)   //取得權杖桶
	AddBuckets(rules ...LimiterBucketRule) LimitIface //新增多個權杖桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        //自定鍵值名稱
	FillInterval time.Duration //間隔多久時間放N個權杖
	Capacity     int64         //權杖桶容量
	Quantum      int64         //每次到達間隔時間後所放的實際權杖數量
}
