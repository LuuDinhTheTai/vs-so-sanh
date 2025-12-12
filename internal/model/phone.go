package model

type Phone struct {
	Brand    string // Samsung
	Model    string // Galaxy Z Tri-Fold (Giả định dựa trên specs)
	Network  Network
	Launch   Launch
	Body     Body
	Display  Display
	Platform Platform
	Memory   Memory
	Camera   Camera
	Sound    Sound
	Comms    Comms
	Features Features
	Battery  Battery
	Misc     Misc
}

type Network struct {
	Technology string
	Bands2G    string // Tên biến không được bắt đầu bằng số
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
	IPRating           string // IP48
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

type Camera struct {
	MainModules    string
	MainFeatures   string
	MainVideo      string
	SelfieModules  string // Bao gồm cả camera bìa (cover camera)
	SelfieFeatures string
	SelfieVideo    string
}

type Sound struct {
	Loudspeaker string
	Jack35mm    string // Dùng string để lưu "No" hoặc chi tiết khác
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
	Other   string // DeX support
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
