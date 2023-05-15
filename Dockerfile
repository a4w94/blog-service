# 指定基底image
FROM golang:latest
# 指定專案資料夾(根目錄)
WORKDIR /app
# 將專案內所有程式碼複製到docker 指定專案資料夾
COPY . /app
# 安裝專案所需套件
RUN go mod download

# 將專案編譯成可執行檔
RUN go build -o main
# 指定對外port 要和專案內的指定port一樣
EXPOSE 8000

ENV DB_HOST=mysql8019ㄦ
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=a4w941207
#執行已編譯好的2進位檔
CMD ["./main"]

#終端機指令
#build一個叫做user_name/project_name的image(記得要 .)
#$ docker build -t terry/blog_service .

# run一個container 並且將port號綁定本機的port號  
# --name NAMES -p 本機PORT:專案對外PORT terry/blog_service(IMAGE)
#$ docker run --name terry_blog --rm -d -p 5000:8000 host terry/blog_service

#本機環境
#$ docker run --name terry_blog --rm -d --network host terry/blog_service

#列出正在運行的container
#$ docker ps 

#將container停掉
#$ docker stop <CONTAINER ID>

#刪除image by id
#$ docker image rm -f IMAGE <image_id>

#更改現有容器的名稱 名稱唯一
#$ docker rename old-name new-name





