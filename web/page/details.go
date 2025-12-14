package page

import (
	"fmt"
	"strings"
	brandDto "vs-so-sanh/internal/brand/delivery/dto"
	"vs-so-sanh/internal/device/delivery/dto"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DetailsPage(device *dto.DeviceResponse, top20Brands []brandDto.BrandResponse) Node {
	specs := device.Specifications

	return shared.Page(device.ModelName,
		Div(Class("grid grid-cols-1 lg:grid-cols-10 gap-6"),
			// ASIDE (3/10)
			Aside(Class("lg:col-span-3 space-y-6"),
				shared.PhoneFinderWidget(top20Brands),
				shared.Top10Widget(),
			),

			// MAIN CONTENT (7/10)
			Main(Class("lg:col-span-7 space-y-10"),
				// Breadcrumb
				Nav(Class("flex items-center text-sm font-medium text-gray-500"),
					A(Href("/"), Class("hover:text-blue-600 transition-colors"), Text("Home")),
					Span(Class("mx-2 text-gray-400"), Text("/")),
					Span(Class("text-gray-700"), Text(device.ModelName)),
				),

				// HERO SECTION: Image + Main Info
				Div(Class("grid grid-cols-1 lg:grid-cols-12 gap-8"),
					// Left Column: Product Image
					Div(Class("lg:col-span-5 bg-gray-50 border border-gray-200 rounded-lg p-6 flex items-center justify-center"),
						Img(Src(device.ImageUrl), Alt(device.ModelName), Class("max-h-96 w-auto object-contain")),
					),

					// Right Column: Key Information & Actions
					Div(Class("lg:col-span-7 flex flex-col"),
						H1(Class("text-4xl font-bold text-gray-900 mb-2 leading-tight"), Text(device.ModelName)),
						P(Class("text-base text-gray-600 mb-6"),
							Text(fmt.Sprintf("Explore the full specifications, features, and price of the %s. See how it compares to the competition and decide if it's the right device for you.", device.ModelName)),
						),

						// Price
						Div(Class("flex flex-wrap items-center gap-x-4 gap-y-2 mb-6 border-y border-gray-200 py-4"),
							P(Class("text-3xl font-bold text-blue-600"), Text(specs.Misc.Price)),
							Span(Class("text-sm text-gray-500 font-medium"), Text("Estimated Price")),
						),

						// Spec Highlights
						Div(Class("grid grid-cols-2 gap-4 mb-8"),
							flatHighlightBox("Chipset", specs.Platform.Chipset),
							flatHighlightBox("Display", specs.Display.MainSize),
							flatHighlightBox("Memory", specs.Memory.Internal),
							flatHighlightBox("Battery", specs.Battery.Type),
						),

						// Action Buttons
						Div(Class("flex flex-col sm:flex-row gap-3"),
							Button(
								Class("w-full sm:w-auto flex-1 bg-blue-600 text-white font-bold py-3 px-6 rounded-lg text-center uppercase tracking-wider text-sm transition-colors hover:bg-blue-700 flex items-center justify-center"),
								I(Class("fas fa-plus mr-2")), // Assuming FontAwesome
								Text("Add to Compare"),
							),
							Button(
								Class("w-full sm:w-auto flex-1 bg-white text-gray-800 border border-gray-300 font-bold py-3 px-6 rounded-lg text-center uppercase tracking-wider text-sm transition-colors hover:bg-gray-100 flex items-center justify-center"),
								Text("Full Review"),
							),
						),
					),
				),

				// DETAILED SPECS SECTION
				Div(
					H2(Class("text-2xl font-bold text-gray-800 mb-6 border-b border-gray-200 pb-3"), Text("Technical Specifications")),
					Div(Class("grid md:grid-cols-2 gap-x-12 gap-y-10"),
						// Column 1
						Div(Class("space-y-8"),
							specSectionFlat("Network & Connectivity",
								specRowFlat("Technology", specs.Network.Technology),
								specRowFlat("Speed", specs.Network.Speed),
								specRowFlat("WLAN", specs.Comms.WLAN),
								specRowFlat("Bluetooth", specs.Comms.Bluetooth),
								specRowFlat("NFC", specs.Comms.NFC),
							),
							specSectionFlat("Body & Design",
								specRowFlat("Dimensions", specs.Body.DimensionsUnfolded),
								specRowFlat("Weight", specs.Body.Weight),
								specRowFlat("Build", specs.Body.Build),
								specRowFlat("SIM", specs.Body.SIM),
							),
							specSectionFlat("Display",
								specRowFlat("Type", specs.Display.MainType),
								specRowFlat("Size", specs.Display.MainSize),
								specRowFlat("Resolution", specs.Display.MainResolution),
							),
						),

						// Column 2
						Div(Class("space-y-8"),
							specSectionFlat("Platform",
								specRowFlat("OS", specs.Platform.OS),
								specRowFlat("Chipset", specs.Platform.Chipset),
								specRowFlat("CPU", specs.Platform.CPU),
								specRowFlat("GPU", specs.Platform.GPU),
							),
							specSectionFlat("Memory",
								specRowFlat("Card slot", specs.Memory.CardSlot),
								specRowFlat("Internal", specs.Memory.Internal),
							),
							specSectionFlat("Camera",
								specRowFlat("Main Camera", specs.MainCamera.Single),
								specRowFlat("Main Video", specs.MainCamera.Video),
								specRowFlat("Selfie Camera", specs.SelfieCamera.Single),
								specRowFlat("Selfie Video", specs.SelfieCamera.Video),
							),
							specSectionFlat("Battery",
								specRowFlat("Type", specs.Battery.Type),
								specRowFlat("Charging", specs.Battery.Charging),
							),
						),
					),
				),
			),
		),
	)
}

// --- HELPER COMPONENTS (FLAT DESIGN) ---

func flatHighlightBox(label, value string) Node {
	if value == "" {
		return nil
	}
	return Div(Class("border border-gray-200 rounded-lg p-4 hover:border-blue-500 transition-colors group cursor-default bg-white"),
		Span(Class("block text-xs text-gray-500 uppercase font-semibold tracking-wider mb-1 group-hover:text-blue-500"), Text(label)),
		Span(Class("block text-sm font-bold text-gray-800 line-clamp-2"), Text(value)),
	)
}

func specSectionFlat(title string, rows ...Node) Node {
	return Div(
		H3(Class("text-base font-bold text-gray-800 mb-4 uppercase tracking-wider"), Text(title)),
		Div(Class("space-y-3"), Group(rows)),
	)
}

func specRowFlat(label, value string) Node {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	return Div(Class("grid grid-cols-3 gap-4 py-3 border-b border-gray-100"),
		Span(Class("col-span-1 text-sm text-gray-600"), Text(label)),
		Span(Class("col-span-2 text-sm text-gray-800 font-medium"), Text(value)),
	)
}
