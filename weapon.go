package statink2

type Weapon string

const (
	Dot52Gal            Weapon = "52gal"
	Dot96Gal            Weapon = "96gal"
	Bold                Weapon = "bold"
	HeroShooterReplica  Weapon = "heroshooter_replica"
	JetSweeper          Weapon = "jetsweeper"
	JetSweeperCustom    Weapon = "jetsweeper_custom"
	Momiji              Weapon = "momiji"
	NZAP85              Weapon = "nzap85"
	Prime               Weapon = "prime"
	PrimeCollabo        Weapon = "prime_collabo"
	PromodelerMG        Weapon = "promodeler_mg"
	PromodelerRG        Weapon = "promodeler_rg"
	Sharp               Weapon = "sharp"
	SShooter            Weapon = "sshooter"
	SShooterCollabo     Weapon = "sshooter_collabo"
	Wakaba              Weapon = "wakaba"
	ClashBlaster        Weapon = "clashblaster"
	HeroBlasterReplica  Weapon = "heroblaster_replica"
	HotBlaster          Weapon = "hotblaster"
	HotBlasterCustom    Weapon = "hotblaster_custom"
	LongBlaster         Weapon = "longblaster"
	Nova                Weapon = "nova"
	Rapid               Weapon = "rapid"
	RapidElite          Weapon = "rapid_elite"
	H3ReelGun           Weapon = "h3reelgun"
	L3ReelGun           Weapon = "l3reelgun"
	DualSweeper         Weapon = "dualsweeper"
	HeroManeuverReplica Weapon = "heromaneuver_replica"
	Maneuver            Weapon = "maneuver"
	ManeuverCollabo     Weapon = "maneuver_collabo"
	Sputtery            Weapon = "sputtery"
	Carbon              Weapon = "carbon"
	Dynamo              Weapon = "dynamo"
	DynamoTesla         Weapon = "dynamo_tesla"
	HeroRollerReplica   Weapon = "heroroller_replica"
	Splatroller         Weapon = "splatroller"
	SplatrollerCollabo  Weapon = "splatroller_collabo"
	VariableRoller      Weapon = "variableroller"
	HeroBrushReplica    Weapon = "herobrush_replica"
	Hokusai             Weapon = "hokusai"
	Pablo               Weapon = "pablo"
	PabloHue            Weapon = "pablo_hue"
	Bamboo14Mk1         Weapon = "bamboo14mk1"
	HeroChargerReplica  Weapon = "herocharger_replica"
	Liter4K             Weapon = "liter4k"
	Liter4KCustom       Weapon = "liter4k_custom"
	Liter4KScope        Weapon = "liter4k_scope"
	Liter4KScopeCustom  Weapon = "liter4k_scope_custom"
	Soytuber            Weapon = "soytuber"
	SplatCharger        Weapon = "splatcharger"
	SplatChargerCollabo Weapon = "splatcharger_collabo"
	SplatScope          Weapon = "splatscope"
	SplatScopeCollabo   Weapon = "splatscope_collabo"
	SquicleanA          Weapon = "squiclean_a"
	BucketSlosher       Weapon = "bucketslosher"
	HeroSlosherReplica  Weapon = "heroslosher_replica"
	Hissen              Weapon = "hissen"
	ScrewSlosher        Weapon = "screwslosher"
	BarrelSpinner       Weapon = "barrelspinner"
	BarrelSpinnerDeco   Weapon = "barrelspinner_deco"
	HeroSpinnerReplica  Weapon = "herospinner_replica"
	SplatSpinner        Weapon = "splatspinner"
	CampingShelter      Weapon = "campingshelter"
	HeroShelterReplica  Weapon = "heroshelter_replica"
	Parashelter         Weapon = "parashelter"
)

var WeaponList = map[Weapon]string{
	Dot52Gal:            ".52ガロン",
	Dot96Gal:            ".96ガロン",
	Bold:                "ボールドマーカー",
	HeroShooterReplica:  "ヒーローシューター レプリカ",
	JetSweeper:          "ジェットスイーパー",
	JetSweeperCustom:    "ジェットスイーパーカスタム",
	Momiji:              "もみじシューター",
	NZAP85:              "N-ZAP85",
	Prime:               "プライムシューター",
	PrimeCollabo:        "プライムシューターコラボ",
	PromodelerMG:        "プロモデラーMG",
	PromodelerRG:        "プロモデラーRG",
	Sharp:               "シャープマーカー",
	SShooter:            "スプラシューター",
	SShooterCollabo:     "スプラシューターコラボ",
	Wakaba:              "わかばシューター",
	ClashBlaster:        "クラッシュブラスター",
	HeroBlasterReplica:  "ヒーローブラスター レプリカ",
	HotBlaster:          "ホットブラスター",
	HotBlasterCustom:    "ホットブラスターカスタム",
	LongBlaster:         "ロングブラスター",
	Nova:                "ノヴァブラスター",
	Rapid:               "ラピッドブラスター",
	RapidElite:          "Rブラスターエリート",
	H3ReelGun:           "H3リールガン",
	L3ReelGun:           "L3リールガン",
	DualSweeper:         "デュアルスイーパー",
	HeroManeuverReplica: "ヒーローマニューバー レプリカ",
	Maneuver:            "スプラマニューバー",
	ManeuverCollabo:     "スプラマニューバーコラボ",
	Sputtery:            "スパッタリー",
	Carbon:              "カーボンローラー",
	Dynamo:              "ダイナモローラー",
	DynamoTesla:         "ダイナモローラーテスラ",
	HeroRollerReplica:   "ヒーローローラーレプリカ",
	Splatroller:         "スプラローラー",
	SplatrollerCollabo:  "スプラローラーコラボ",
	VariableRoller:      "ヴァリアブルローラー",
	HeroBrushReplica:    "ヒーローブラシレプリカ",
	Hokusai:             "ホクサイ",
	Pablo:               "パブロ",
	PabloHue:            "パブロ・ヒュー",
	Bamboo14Mk1:         "14式竹筒銃・甲",
	HeroChargerReplica:  "ヒーローチャージャー レプリカ",
	Liter4K:             "リッター4K",
	Liter4KCustom:       "リッター4Kカスタム",
	Liter4KScope:        "4Kスコープ",
	Liter4KScopeCustom:  "4Kスコープカスタム",
	Soytuber:            "ソイチューバー",
	SplatCharger:        "スプラチャージャー",
	SplatChargerCollabo: "スプラチャージャーコラボ",
	SplatScope:          "スプラスコープ",
	SplatScopeCollabo:   "スプラスコープコラボ",
	SquicleanA:          "スクイックリンα",
	BucketSlosher:       "バケットスロッシャー",
	HeroSlosherReplica:  "ヒーロースロッシャー レプリカ",
	Hissen:              "ヒッセン",
	ScrewSlosher:        "スクリュースロッシャー",
	BarrelSpinner:       "バレルスピナー",
	BarrelSpinnerDeco:   "バレルスピナーデコ",
	HeroSpinnerReplica:  "ヒーロースピナー レプリカ",
	SplatSpinner:        "スプラスピナー",
	CampingShelter:      "キャンピングシェルター",
	HeroShelterReplica:  "ヒーローシェルター レプリカ",
	Parashelter:         "パラシェルター",
}
