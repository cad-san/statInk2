package statink2

type Stage string

const (
	Ama       Stage = "ama"
	Arowana   Stage = "arowana"
	Battera   Stage = "battera"
	BBass     Stage = "bbass"
	Chozame   Stage = "chozame"
	Devon     Stage = "devon"
	Engawa    Stage = "engawa"
	Fujitsubo Stage = "fujitsubo"
	Gangaze   Stage = "gangaze"
	Hakofugu  Stage = "hakofugu"
	Hokke     Stage = "hokke"
	Kombu     Stage = "kombu"
	Manta     Stage = "manta"
	Mozuku    Stage = "mozuku"
	Tachiuo   Stage = "tachiuo"
	Zatou     Stage = "zatou"
	Mystery   Stage = "mystery"
)

var StageList = map[Stage]string{
	Ama:       "海女美術大学",
	Arowana:   "アロワナモール",
	Battera:   "バッテラストリート",
	BBass:     "Bバスパーク",
	Chozame:   "チョウザメ造船",
	Devon:     "デボン海洋博物館",
	Engawa:    "エンガワ河川敷",
	Fujitsubo: "フジツボスポーツクラブ",
	Gangaze:   "ガンガゼ野外音楽堂",
	Hakofugu:  "ハコフグ倉庫",
	Hokke:     "ホッケふ頭",
	Kombu:     "コンブトラック",
	Manta:     "マンタマリア号",
	Mozuku:    "モズク農園",
	Tachiuo:   "タチウオパーキング",
	Zatou:     "ザトウマーケット",
	Mystery:   "ミステリーゾーン",
}
