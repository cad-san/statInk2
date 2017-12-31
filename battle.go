package statink2

const (
	Yes = "yes"
	No  = "no"

	// ResultWin means win for Battle Result.
	ResultWin = "win"
	// ResultLose means lose for Battle Result.
	ResultLose = "lose"
)

type Lobby string
type Mode string
type Rule string

const (
	Standard Lobby = "standard"
	Squad2   Lobby = "squad_2"
	Squad4   Lobby = "squad_4"
	Private  Lobby = "private"
)

const (
	Regular       Mode = "regular"
	Gachi         Mode = "gachi"
	PrivateBattle Mode = "private"
)

const (
	Nawabari Rule = "nawabari"
	Area     Rule = "area"
	Yagura   Rule = "yagura"
	Hoko     Rule = "hoko"
	Asari    Rule = "asari"
)
