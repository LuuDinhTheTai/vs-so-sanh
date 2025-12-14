package page

import (
	"fmt"
	"vs-so-sanh/internal/device/delivery/dto"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ComparePage(device1 *dto.DeviceResponse, device2 *dto.DeviceResponse) Node {
	return shared.Page(fmt.Sprintf("%s vs %s", device1.ModelName, device2.ModelName),
		Div(Class("max-w-7xl mx-auto px-4 py-8"),
			H1(Class("text-3xl font-bold text-gray-900 mb-8"), Text(fmt.Sprintf("%s vs %s", device1.ModelName, device2.ModelName))),

			Div(Class("grid grid-cols-1 md:grid-cols-2 gap-8"),
				renderDeviceComparison(device1),
				renderDeviceComparison(device2),
			),

			Div(Class("mt-12"),
				H2(Class("text-2xl font-bold text-gray-900 mb-6"), Text("Detailed Comparison")),
				renderDetailedComparison(device1, device2),
			),
		),
	)
}

func renderDeviceComparison(device *dto.DeviceResponse) Node {
	return Div(Class("bg-white rounded-lg shadow-md p-6"),
		Img(Src(device.ImageUrl), Alt(device.ModelName), Class("w-48 h-48 object-contain mx-auto mb-4")),
		H3(Class("text-xl font-semibold text-gray-800 text-center mb-2"), Text(device.ModelName)),
		P(Class("text-lg font-bold text-blue-700 text-center"), Text(device.Specifications.Misc.Price)),
		Div(Class("mt-4 text-sm text-gray-600"),
			P(Text(fmt.Sprintf("OS: %s", device.Specifications.Platform.OS))),
			P(Text(fmt.Sprintf("Chipset: %s", device.Specifications.Platform.Chipset))),
			P(Text(fmt.Sprintf("Display: %s, %s", device.Specifications.Display.MainSize, device.Specifications.Display.MainResolution))),
			P(Text(fmt.Sprintf("Memory: %s", device.Specifications.Memory.Internal))),
			P(Text(fmt.Sprintf("Battery: %s", device.Specifications.Battery.Type))),
		),
	)
}

func renderDetailedComparison(device1 *dto.DeviceResponse, device2 *dto.DeviceResponse) Node {
	return Div(Class("overflow-x-auto"),
		Table(Class("table-auto w-full"),
			THead(
				Tr(
					Th(Class("px-4 py-2 text-left text-sm font-semibold text-gray-700 bg-gray-100"), Text("Feature")),
					Th(Class("px-4 py-2 text-left text-sm font-semibold text-gray-700 bg-gray-100"), Text(device1.ModelName)),
					Th(Class("px-4 py-2 text-left text-sm font-semibold text-gray-700 bg-gray-100"), Text(device2.ModelName)),
				),
			),
			TBody(
				renderComparisonRow("Price", device1.Specifications.Misc.Price, device2.Specifications.Misc.Price),
				renderComparisonRow("Announced", device1.Specifications.Launch.Announced, device2.Specifications.Launch.Announced),
				renderComparisonRow("Status", device1.Specifications.Launch.Status, device2.Specifications.Launch.Status),
				renderComparisonRow("Weight", device1.Specifications.Body.Weight, device2.Specifications.Body.Weight),
				renderComparisonRow("Build", device1.Specifications.Body.Build, device2.Specifications.Body.Build),
				renderComparisonRow("SIM", device1.Specifications.Body.SIM, device2.Specifications.Body.SIM),
				renderComparisonRow("IP Rating", device1.Specifications.Body.IPRating, device2.Specifications.Body.IPRating),
				renderComparisonRow("Display Type", device1.Specifications.Display.MainType, device2.Specifications.Display.MainType),
				renderComparisonRow("Display Size", device1.Specifications.Display.MainSize, device2.Specifications.Display.MainSize),
				renderComparisonRow("Display Resolution", device1.Specifications.Display.MainResolution, device2.Specifications.Display.MainResolution),
				renderComparisonRow("OS", device1.Specifications.Platform.OS, device2.Specifications.Platform.OS),
				renderComparisonRow("Chipset", device1.Specifications.Platform.Chipset, device2.Specifications.Platform.Chipset),
				renderComparisonRow("CPU", device1.Specifications.Platform.CPU, device2.Specifications.Platform.CPU),
				renderComparisonRow("GPU", device1.Specifications.Platform.GPU, device2.Specifications.Platform.GPU),
				renderComparisonRow("Card Slot", device1.Specifications.Memory.CardSlot, device2.Specifications.Memory.CardSlot),
				renderComparisonRow("Internal Memory", device1.Specifications.Memory.Internal, device2.Specifications.Memory.Internal),
				renderComparisonRow("Main Camera", device1.Specifications.MainCamera.Single, device2.Specifications.MainCamera.Single),
				renderComparisonRow("Main Camera Features", device1.Specifications.MainCamera.Features, device2.Specifications.MainCamera.Features),
				renderComparisonRow("Main Camera Video", device1.Specifications.MainCamera.Video, device2.Specifications.MainCamera.Video),
				renderComparisonRow("Selfie Camera", device1.Specifications.SelfieCamera.Single, device2.Specifications.SelfieCamera.Single),
				renderComparisonRow("Selfie Camera Video", device1.Specifications.SelfieCamera.Video, device2.Specifications.SelfieCamera.Video),
				renderComparisonRow("Loudspeaker", device1.Specifications.Sound.Loudspeaker, device2.Specifications.Sound.Loudspeaker),
				renderComparisonRow("3.5mm Jack", device1.Specifications.Sound.Jack35mm, device2.Specifications.Sound.Jack35mm),
				renderComparisonRow("WLAN", device1.Specifications.Comms.WLAN, device2.Specifications.Comms.WLAN),
				renderComparisonRow("Bluetooth", device1.Specifications.Comms.Bluetooth, device2.Specifications.Comms.Bluetooth),
				renderComparisonRow("Positioning", device1.Specifications.Comms.Positioning, device2.Specifications.Comms.Positioning),
				renderComparisonRow("NFC", device1.Specifications.Comms.NFC, device2.Specifications.Comms.NFC),
				renderComparisonRow("Radio", device1.Specifications.Comms.Radio, device2.Specifications.Comms.Radio),
				renderComparisonRow("USB", device1.Specifications.Comms.USB, device2.Specifications.Comms.USB),
				renderComparisonRow("Sensors", device1.Specifications.Features.Sensors, device2.Specifications.Features.Sensors),
				renderComparisonRow("Battery Type", device1.Specifications.Battery.Type, device2.Specifications.Battery.Type),
				renderComparisonRow("Charging", device1.Specifications.Battery.Charging, device2.Specifications.Battery.Charging),
				renderComparisonRow("Colors", device1.Specifications.Misc.Colors, device2.Specifications.Misc.Colors),
				renderComparisonRow("Models", device1.Specifications.Misc.Models, device2.Specifications.Misc.Models),
			),
		),
	)
}

func renderComparisonRow(feature, value1, value2 string) Node {
	return Tr(Class("border-b border-gray-200"),
		Td(Class("px-4 py-3 text-sm font-medium text-gray-800"), Text(feature)),
		Td(Class("px-4 py-3 text-sm text-gray-700"), Text(value1)),
		Td(Class("px-4 py-3 text-sm text-gray-700"), Text(value2)),
	)
}
