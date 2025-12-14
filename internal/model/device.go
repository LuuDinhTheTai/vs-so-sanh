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
	Dimensions string `bson:"Dimensions"`
	Weight     string `bson:"Weight"`
	Build      string `bson:"Build"`
	SIM        string `bson:"SIM"`
	IPRating   string `bson:"IP Rating"`
}

type Display struct {
	Type       string `bson:"Type"`
	Size       string `bson:"Size"`
	Resolution string `bson:"Resolution"`
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
	Triple   string `bson:"Triple"`
	Features string `bson:"Features"`
	Single   string `bson:"Single"`
	Video    string `bson:"Video"`
}

type SelfieCamera struct {
	Single string `bson:"Single"`
	Video  string `bson:"Video"`
}

type Sound struct {
	Loudspeaker string `bson:"Loudspeaker"`
	Jack35mm    string `bson:"3.5mm jack"`
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
