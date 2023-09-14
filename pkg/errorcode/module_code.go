package errorcode

var (
	ErrorGetTagListFail  = NewError(20010001, "取得標籤列表失敗")
	ErrorCreateTagFail   = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail   = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail   = NewError(2001004, "刪除標籤失敗")
	ErrorCountTagFail    = NewError(20010005, "統計標籤失敗")
	ErrorTagAlreadyExist = NewError(20010006, "標籤已存在")

	ErrorUploadFileFail = NewError(20030001, "上傳檔案失敗")
)

var (
	ErrorGetArticleListFail = NewError(20020001, "取得文章列表失敗")
	ErrorCreateArticleFail  = NewError(20020002, "建立文章失敗")
	ErrorUpdateArticleFail  = NewError(20020003, "更新文章失敗")
	ErrorDeleteArticleFail  = NewError(20020004, "删除文章失敗")
	ErrorCountArticleFail   = NewError(20020005, "計算文章數量失敗")
	ErrorGetArticleFail     = NewError(20020006, "取得文章失敗")
)
