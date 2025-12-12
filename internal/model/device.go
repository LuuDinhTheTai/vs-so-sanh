package model

type Device struct {
	ModelName      string         `bson:"model_name"`
	ImageUrl       string         `bson:"imageUrl"`
	Specifications Specifications `bson:"specifications"`
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
	Technology string `bson:"Technology"`
	Bands2G    string `bson:"2G bands"`
	Bands3G    string `bson:"3G bands"`
	Bands4G    string `bson:"4G bands"`
	Bands5G    string `bson:"5G bands"`
	Speed      string `bson:"Speed"`
}

type Launch struct {
	Announced string `bson:"Announced"`
	Status    string `bson:"Status"`
}

type Body struct {
	DimensionsUnfolded string `bson:"Dimensions Unfolded"`
	DimensionsFolded   string `bson:"Dimensions Folded"`
	Weight             string `bson:"Weight"`
	Build              string `bson:"Build"`
	SIM                string `bson:"SIM"`
	IPRating           string `bson:"IP Rating"`
}

type Display struct {
	MainType        string `bson:"Main Type"`
	MainSize        string `bson:"Main Size"`
	MainResolution  string `bson:"Main Resolution"`
	CoverType       string `bson:"Cover Type"`
	CoverSize       string `bson:"Cover Size"`
	CoverResolution string `bson:"Cover Resolution"`
}

type Platform struct {
	OS      string `bson:"OS"`
	Chipset string `bson:"Chipset"`
	CPU     string `bson:"CPU"`
	GPU     string `bson:"GPU"`
}

type Memory struct {
	CardSlot string `bson:"Card Slot"`
	Internal string `bson:"Internal"`
}

type MainCamera struct {
	Single   string `bson:"Single"`
	Features string `bson:"Features"`
	Video    string `bson:"Video"`
}

type SelfieCamera struct {
	Single string `bson:"Single"`
	Video  string `bson:"Video"`
}

type Sound struct {
	Loudspeaker string `bson:"Loudspeaker"`
	Jack35mm    string `bson:"Jack35Mm"`
}

type Comms struct {
	WLAN        string `bson:"WLAN"`
	Bluetooth   string `bson:"Bluetooth"`
	Positioning string `bson:"Positioning"`
	NFC         string `bson:"NFC"`
	Radio       string `bson:"Radio"`
	USB         string `bson:"USB"`
}

type Features struct {
	Sensors string `bson:"Sensors"`
	Other   string `bson:"Other"`
}

type Battery struct {
	Type     string `bson:"Type"`
	Charging string `bson:"Charging"`
}

type Misc struct {
	Colors string `bson:"Colors"`
	Models string `bson:"Models"`
	Price  string `bson:"Price"`
}
