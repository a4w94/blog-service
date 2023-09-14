//流程位置
1.api路由器--api帶入參數相關驗證綁定，api回應訊息
=>internal/routers/version
ex:
func(r Tag)Create(c *gin.Context){
    ...
    err := svc.CreateTag(&param)
    response.ToResponse(gin.H{})
}
2.api路由器業務邏輯封裝
=>internal/service
ex:
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}
3.資料庫操作封裝帶入參數
=>internal/dao
ex:
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return tag.Create(d.engine)
}
4.資料庫操作邏輯,sql敘述
=>internal/model
ex:
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}



//啟動Jaeger =>追蹤api相關資訊用
docker run -d --name jaeger \
-e COLLETCTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16

//auth 
{header}
app_key:terry
app_secret:blog-service

//go-bindata 包裝設定檔.資源檔進go程式
//-o 指定路徑 -pkg 指定包名
$ go-bindata -o configs/config.go -pkg=configs configs/config.yaml

//壓縮執行檔
$ upx {執行檔名稱}
//非必要壓縮方式
$ go build -ldflags="-w -s"
-w :去除DWARF偵錯資訊，panic拋出時，呼叫堆疊資訊沒有檔名，行號資訊
-s :去除符號資訊，無法使用gdb偵錯

//使用-ldflags 將編譯時間 版本編號 git hash寫入二進位執行檔
$ go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse origin`"
//執行 並賦值給isVersion(flag name:version)
$ ./blog_service -version test

//swagger 頁面
http://127.0.0.1:8000/swagger/index.html
go get -u github.com/swaggo/swag/cmd/swag
export PATH=$PATH:/Users/terry_hsiesh/go/bin
test
new 11
