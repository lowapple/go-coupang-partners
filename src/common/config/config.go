package config

var Opts struct {
	// 쿠팡 ID
	CoupangId string `long:"id" description:"coupang partners id" required:"true"`
	// 쿠팡 PASSWORD
	CoupangPw string `long:"pw" description:"coupang partners pw" required:"true"`
	// 쿠팡 검색키워드
	Keyword string `short:"k" long:"keyword" description:"coupang search keyword" require:"false"`
}
