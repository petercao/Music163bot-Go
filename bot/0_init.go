package bot

import (
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"gorm.io/gorm"
)

var data utils.RequestData
var botName string
var cacheDir = "./cache"

var (
	aboutText = `*Music163bot-Go %s*
Github: https://github.com/XiaoMengXinX/Music163bot-Go

\[编译环境] %s
\[程序版本] %s
\[编译哈希] %s
\[编译日期] %s
\[编译系统] %s
\[运行环境] %s`
	musicInfo = `「%s」- %s
专辑: %s
#网易云音乐 #%s %.2fkpbs
via @%s`
	musicInfoMsg = `%s
专辑: %s
%s %sMB
`
	musicInfoNoBr = `%s
专辑: %s
`
	rmcacheReportAll = `清除全部数据库成功`
	rmcacheReport    = `清除 musicid : %s 缓存成功`
	searching        = `搜索中...`
	noResults        = `未找到结果`
	noCache          = `歌曲未缓存`
	tapToDownload    = `点击上方按钮缓存歌曲`
	tapMeToDown      = `点我缓存歌曲`
	unkonwnError     = `未知错误`
	alreadyStart     = `存在下载任务，请稍候...`
	hitCache         = `命中缓存，正在发送中...`
	sendMeTo         = `Send me to...`
	uploadFailed     = `发送失败 %s`
	getUrlFailed     = `获取下载链接失败`
	fetchInfo        = `正在获取歌曲信息...`
	fetchInfoFailed  = `获取歌曲信息失败`
	waitForDown      = `等待下载中...`
	downloading      = `下载中...`
	uploading        = `下载完成，发送中...`
)

type SongInfo struct {
	gorm.Model
	MusicID      int
	SongName     string
	SongArtists  string
	SongAlbum    string
	FileExt      string
	FileSize     string
	BitRate      int
	Duration     int
	FileID       string
	ThumbFileID  string
	FromUserID   int64
	FromUserName string
	FromChatID   int64
	FromChatName string
}
