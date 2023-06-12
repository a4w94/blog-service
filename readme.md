//啟動Jaeger 
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