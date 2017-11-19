package statink2

type Stage string

const (
	Ama       Stage = "ama"
	Battera   Stage = "battera"
	BBass     Stage = "bbass"
	Chozame   Stage = "chozame"
	Engawa    Stage = "engawa"
	Fujitsubo Stage = "fujitsubo"
	Gangaze   Stage = "gangaze"
	Hokke     Stage = "hokke"
	Kombu     Stage = "kombu"
	Manta     Stage = "manta"
	Mozuku    Stage = "mozuku"
	Tachiuo   Stage = "tachiuo"
	Mystery   Stage = "mystery"
)

var StageList = map[Stage]string{
	Ama:       "海女美術大学",
	Battera:   "バッテラストリート",
	BBass:     "Bバスパーク",
	Chozame:   "チョウザメ造船",
	Engawa:    "エンガワ河川敷",
	Fujitsubo: "フジツボスポーツクラブ",
	Gangaze:   "ガンガゼ野外音楽堂",
	Hokke:     "ホッケふ頭",
	Kombu:     "コンブトラック",
	Manta:     "マンタマリア号",
	Mozuku:    "モズク農園",
	Tachiuo:   "タチウオパーキング",
	Mystery:   "ミステリーゾーン",
}
