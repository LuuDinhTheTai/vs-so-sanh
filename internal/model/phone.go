package model

type Phone struct {
	BrandName      string
	ModelName      string
	ImageUrl       string
	Specifications Specifications
}

type Specifications struct {
	Network      Network
	Launch       Launch
	Body         Body
	Display      Display
	Platform     Platform
	Memory       Memory
	MainCamera   MainCamera
	SelfieCamera SelfieCamera
	Sound        Sound
	Comms        Comms
	Features     Features
	Battery      Battery
	Misc         Misc
}

type Network struct {
	Technology string
	Bands2G    string
	Bands3G    string
	Bands4G    string
	Bands5G    string
	Speed      string
}

type Launch struct {
	Announced string
	Status    string
}

type Body struct {
	DimensionsUnfolded string
	DimensionsFolded   string
	Weight             string
	Build              string
	SIM                string
	IPRating           string
}

type Display struct {
	MainType        string
	MainSize        string
	MainResolution  string
	CoverType       string
	CoverSize       string
	CoverResolution string
}

type Platform struct {
	OS      string
	Chipset string
	CPU     string
	GPU     string
}

type Memory struct {
	CardSlot string
	Internal string
}

type MainCamera struct {
	Single   string
	Features string
	Video    string
}

type SelfieCamera struct {
	Single string
	Video  string
}

type Sound struct {
	Loudspeaker string
	Jack35mm    string
}

type Comms struct {
	WLAN        string
	Bluetooth   string
	Positioning string
	NFC         string
	Radio       string
	USB         string
}

type Features struct {
	Sensors string
	Other   string
}

type Battery struct {
	Type     string
	Charging string
}

type Misc struct {
	Colors string
	Models string
	Price  string
}
