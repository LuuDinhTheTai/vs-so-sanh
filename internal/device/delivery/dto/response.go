package dto

import "vs-so-sanh/internal/model"

func From(device *model.Device) *DeviceResponse {
	return &DeviceResponse{
		ModelName: device.ModelName,
		ImageUrl:  device.ImageUrl,
		Specifications: SpecificationsResponse{
			Network: NetworkResponse{
				Technology: device.Specifications.Network.Technology,
				Bands2G:    device.Specifications.Network.Bands2G,
				Bands3G:    device.Specifications.Network.Bands3G,
				Bands4G:    device.Specifications.Network.Bands4G,
				Bands5G:    device.Specifications.Network.Bands5G,
				Speed:      device.Specifications.Network.Speed,
			},
			Launch: LaunchResponse{
				Announced: device.Specifications.Launch.Announced,
				Status:    device.Specifications.Launch.Status,
			},
			Body: BodyResponse{
				Weight:   device.Specifications.Body.Weight,
				Build:    device.Specifications.Body.Build,
				SIM:      device.Specifications.Body.SIM,
				IPRating: device.Specifications.Body.IPRating,
			},
			Display: DisplayResponse{
				MainType:       device.Specifications.Display.Type,
				MainSize:       device.Specifications.Display.Size,
				MainResolution: device.Specifications.Display.Resolution,
			},
			Platform: PlatformResponse{
				OS:      device.Specifications.Platform.OS,
				Chipset: device.Specifications.Platform.Chipset,
				CPU:     device.Specifications.Platform.CPU,
				GPU:     device.Specifications.Platform.GPU,
			},
			Memory: MemoryResponse{
				CardSlot: device.Specifications.Memory.CardSlot,
				Internal: device.Specifications.Memory.Internal,
			},
			MainCamera: MainCameraResponse{
				Single:   device.Specifications.MainCamera.Single,
				Features: device.Specifications.MainCamera.Features,
				Video:    device.Specifications.MainCamera.Video,
			},
			SelfieCamera: SelfieCameraResponse{
				Single: device.Specifications.SelfieCamera.Single,
				Video:  device.Specifications.SelfieCamera.Video,
			},
			Sound: SoundResponse{
				Loudspeaker: device.Specifications.Sound.Loudspeaker,
				Jack35mm:    device.Specifications.Sound.Jack35mm,
			},
			Comms: CommsResponse{
				WLAN:        device.Specifications.Comms.WLAN,
				Bluetooth:   device.Specifications.Comms.Bluetooth,
				Positioning: device.Specifications.Comms.Positioning,
				NFC:         device.Specifications.Comms.NFC,
				Radio:       device.Specifications.Comms.Radio,
				USB:         device.Specifications.Comms.USB,
			},
			Features: FeaturesResponse{
				Sensors: device.Specifications.Features.Sensors,
			},
			Battery: BatteryResponse{
				Type:     device.Specifications.Battery.Type,
				Charging: device.Specifications.Battery.Charging,
			},
			Misc: MiscResponse{
				Colors: device.Specifications.Misc.Colors,
				Models: device.Specifications.Misc.Models,
				Price:  device.Specifications.Misc.Price,
			},
		},
	}
}

var EmptyDeviceResponse = &DeviceResponse{
	ModelName: "Unknown Device",
	ImageUrl:  "",
	Specifications: SpecificationsResponse{
		Network: NetworkResponse{
			Technology: "",
			Bands2G:    "",
			Bands3G:    "",
			Bands4G:    "",
			Bands5G:    "",
			Speed:      "",
		},
		Launch: LaunchResponse{
			Announced: "",
			Status:    "",
		},
		Body: BodyResponse{
			DimensionsUnfolded: "",
			DimensionsFolded:   "",
			Weight:             "",
			Build:              "",
			SIM:                "",
			IPRating:           "",
		},
		Display: DisplayResponse{
			MainType:        "",
			MainSize:        "",
			MainResolution:  "",
			CoverType:       "",
			CoverSize:       "",
			CoverResolution: "",
		},
		Platform: PlatformResponse{
			OS:      "",
			Chipset: "",
			CPU:     "",
			GPU:     "",
		},
		Memory: MemoryResponse{
			CardSlot: "",
			Internal: "",
		},
		MainCamera: MainCameraResponse{
			Single:   "",
			Features: "",
			Video:    "",
		},
		SelfieCamera: SelfieCameraResponse{
			Single: "",
			Video:  "",
		},
		Sound: SoundResponse{
			Loudspeaker: "",
			Jack35mm:    "",
		},
		Comms: CommsResponse{
			WLAN:        "",
			Bluetooth:   "",
			Positioning: "",
			NFC:         "",
			Radio:       "",
			USB:         "",
		},
		Features: FeaturesResponse{
			Sensors: "",
			Other:   "",
		},
		Battery: BatteryResponse{
			Type:     "",
			Charging: "",
		},
		Misc: MiscResponse{
			Colors: "",
			Models: "",
			Price:  "",
		},
	},
}

type DeviceResponse struct {
	ModelName      string
	ImageUrl       string
	Specifications SpecificationsResponse
}

type SpecificationsResponse struct {
	Network      NetworkResponse
	Launch       LaunchResponse
	Body         BodyResponse
	Display      DisplayResponse
	Platform     PlatformResponse
	Memory       MemoryResponse
	MainCamera   MainCameraResponse
	SelfieCamera SelfieCameraResponse
	Sound        SoundResponse
	Comms        CommsResponse
	Features     FeaturesResponse
	Battery      BatteryResponse
	Misc         MiscResponse
}

type NetworkResponse struct {
	Technology string
	Bands2G    string
	Bands3G    string
	Bands4G    string
	Bands5G    string
	Speed      string
}

type LaunchResponse struct {
	Announced string
	Status    string
}

type BodyResponse struct {
	DimensionsUnfolded string
	DimensionsFolded   string
	Weight             string
	Build              string
	SIM                string
	IPRating           string
}

type DisplayResponse struct {
	MainType        string
	MainSize        string
	MainResolution  string
	CoverType       string
	CoverSize       string
	CoverResolution string
}

type PlatformResponse struct {
	OS      string
	Chipset string
	CPU     string
	GPU     string
}

type MemoryResponse struct {
	CardSlot string
	Internal string
}

type MainCameraResponse struct {
	Single   string
	Features string
	Video    string
}

type SelfieCameraResponse struct {
	Single string
	Video  string
}

type SoundResponse struct {
	Loudspeaker string
	Jack35mm    string
}

type CommsResponse struct {
	WLAN        string
	Bluetooth   string
	Positioning string
	NFC         string
	Radio       string
	USB         string
}

type FeaturesResponse struct {
	Sensors string
	Other   string
}

type BatteryResponse struct {
	Type     string
	Charging string
}

type MiscResponse struct {
	Colors string
	Models string
	Price  string
}
