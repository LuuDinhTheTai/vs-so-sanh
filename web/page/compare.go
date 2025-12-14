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

// CheckEmpty is a helper function to replace empty strings with a dash.
func CheckEmpty(s string) string {
	if strings.TrimSpace(s) == "" {
		return "-"
	}
	return s
}

func ComparePage(d1 *dto.DeviceResponse, d2 *dto.DeviceResponse, top20Brands []brandDto.BrandResponse) Node {
	title := fmt.Sprintf("%s vs %s - Full Specs Comparison", d1.ModelName, d2.ModelName)

	return shared.Page(title,
		// Applying the 3/7 grid layout
		Div(Class("grid grid-cols-1 lg:grid-cols-10 gap-6"),

			// ASIDE (3/10)
			Aside(Class("lg:col-span-3 space-y-6"),
				shared.PhoneFinderWidget(top20Brands),
				shared.Top10Widget(),
			),

			// MAIN CONTENT (7/10)
			Main(Class("lg:col-span-7 space-y-6"),
				// Breadcrumb
				//shared.Breadcrumb("Compare", d1.ModelName+" vs "+d2.ModelName),

				// --- STICKY HEADER ---
				Div(Class("sticky top-16 z-40 bg-white/95 backdrop-blur-md border-b-2 border-slate-900 mb-8 pt-4"),
					Div(Class("grid grid-cols-2 gap-4 md:gap-8 relative pb-4"),
						// VS Badge
						Div(Class("absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 z-10 pointer-events-none"),
							Span(Class("bg-slate-900 text-white font-black text-[10px] md:text-xs p-2 rounded-full border-4 border-white shadow-sm"), Text("VS")),
						),
						// Device 1 Column
						compareHeaderItem(d1, 1, d2.ModelName),
						// Device 2 Column
						compareHeaderItem(d2, 2, d1.ModelName),
					),
				),

				// --- MAIN SPECS TABLE ---
				Div(Class("space-y-12 pb-20"),
					compareSection("Launch",
						compareRow("Announced", d1.Specifications.Launch.Announced, d2.Specifications.Launch.Announced),
						compareRow("Status", d1.Specifications.Launch.Status, d2.Specifications.Launch.Status),
					),
					compareSection("Network",
						compareRow("Technology", d1.Specifications.Network.Technology, d2.Specifications.Network.Technology),
						compareRow("2G bands", d1.Specifications.Network.Bands2G, d2.Specifications.Network.Bands2G),
						compareRow("3G bands", d1.Specifications.Network.Bands3G, d2.Specifications.Network.Bands3G),
						compareRow("4G bands", d1.Specifications.Network.Bands4G, d2.Specifications.Network.Bands4G),
						compareRow("5G bands", d1.Specifications.Network.Bands5G, d2.Specifications.Network.Bands5G),
						compareRow("Speed", d1.Specifications.Network.Speed, d2.Specifications.Network.Speed),
					),
					compareSection("Body",
						//compareRow("Dimensions", d1.Specifications.Body.Dimensions, d2.Specifications.Body.Dimensions),
						compareRow("Weight", d1.Specifications.Body.Weight, d2.Specifications.Body.Weight),
						compareRow("Build", d1.Specifications.Body.Build, d2.Specifications.Body.Build),
						compareRow("SIM", d1.Specifications.Body.SIM, d2.Specifications.Body.SIM),
						compareRow("IP Rating", d1.Specifications.Body.IPRating, d2.Specifications.Body.IPRating),
					),
					compareSection("Display",
						compareRow("Type", d1.Specifications.Display.MainType, d2.Specifications.Display.MainType),
						compareRow("Size", d1.Specifications.Display.MainSize, d2.Specifications.Display.MainSize),
						compareRow("Resolution", d1.Specifications.Display.MainResolution, d2.Specifications.Display.MainResolution),
					),
					compareSection("Platform",
						compareRow("OS", d1.Specifications.Platform.OS, d2.Specifications.Platform.OS),
						compareRow("Chipset", d1.Specifications.Platform.Chipset, d2.Specifications.Platform.Chipset),
						compareRow("CPU", d1.Specifications.Platform.CPU, d2.Specifications.Platform.CPU),
						compareRow("GPU", d1.Specifications.Platform.GPU, d2.Specifications.Platform.GPU),
					),
					compareSection("Memory",
						compareRow("Card Slot", d1.Specifications.Memory.CardSlot, d2.Specifications.Memory.CardSlot),
						compareRow("Internal", d1.Specifications.Memory.Internal, d2.Specifications.Memory.Internal),
					),
					compareSection("Main Camera",
						compareRow("Modules", d1.Specifications.MainCamera.Single, d2.Specifications.MainCamera.Single),
						compareRow("Features", d1.Specifications.MainCamera.Features, d2.Specifications.MainCamera.Features),
						compareRow("Video", d1.Specifications.MainCamera.Video, d2.Specifications.MainCamera.Video),
					),
					compareSection("Selfie Camera",
						compareRow("Modules", d1.Specifications.SelfieCamera.Single, d2.Specifications.SelfieCamera.Single),
						compareRow("Video", d1.Specifications.SelfieCamera.Video, d2.Specifications.SelfieCamera.Video),
					),
					compareSection("Sound",
						compareRow("Loudspeaker", d1.Specifications.Sound.Loudspeaker, d2.Specifications.Sound.Loudspeaker),
						compareRow("3.5mm jack", d1.Specifications.Sound.Jack35mm, d2.Specifications.Sound.Jack35mm),
					),
					compareSection("Comms",
						compareRow("WLAN", d1.Specifications.Comms.WLAN, d2.Specifications.Comms.WLAN),
						compareRow("Bluetooth", d1.Specifications.Comms.Bluetooth, d2.Specifications.Comms.Bluetooth),
						compareRow("Positioning", d1.Specifications.Comms.Positioning, d2.Specifications.Comms.Positioning),
						compareRow("NFC", d1.Specifications.Comms.NFC, d2.Specifications.Comms.NFC),
						compareRow("Radio", d1.Specifications.Comms.Radio, d2.Specifications.Comms.Radio),
						compareRow("USB", d1.Specifications.Comms.USB, d2.Specifications.Comms.USB),
					),
					compareSection("Features",
						compareRow("Sensors", d1.Specifications.Features.Sensors, d2.Specifications.Features.Sensors),
					),
					compareSection("Battery",
						compareRow("Type", d1.Specifications.Battery.Type, d2.Specifications.Battery.Type),
						compareRow("Charging", d1.Specifications.Battery.Charging, d2.Specifications.Battery.Charging),
					),
					compareSection("Misc",
						compareRow("Colors", d1.Specifications.Misc.Colors, d2.Specifications.Misc.Colors),
						compareRow("Models", d1.Specifications.Misc.Models, d2.Specifications.Misc.Models),
						compareRow("Price", d1.Specifications.Misc.Price, d2.Specifications.Misc.Price),
					),
				),
			),
		),
	)
}

// --- COMPONENTS ---

func compareHeaderItem(d *dto.DeviceResponse, slot int, otherDeviceSlug string) Node {
	searchResultID := fmt.Sprintf("search-results-%d", slot)
	otherDeviceSlug = strings.ReplaceAll(strings.ToLower(otherDeviceSlug), " ", "-") // Ensure slug is URL-friendly

	return Div(Class("flex flex-col items-center relative group"),
		// 1. Image & Change Button
		Div(Class("relative w-full flex justify-center mb-3"),
			Div(Class("h-24 md:h-32 p-2 bg-gray-50 border border-transparent group-hover:border-gray-200 transition-colors rounded-sm"),
				Img(Src(d.ImageUrl), Alt(d.ModelName), Class("max-h-full w-auto object-contain mix-blend-multiply")),
			),
			Button(
				Class("absolute top-0 right-0 md:right-10 bg-white border border-gray-300 text-gray-500 hover:text-blue-600 hover:border-blue-600 p-1.5 shadow-sm transition-all opacity-100 md:opacity-0 md:group-hover:opacity-100"),
				Attr("onclick", fmt.Sprintf("document.getElementById('change-box-%d').classList.toggle('hidden'); document.getElementById('info-box-%d').classList.toggle('hidden')", slot, slot)),
				I(Class("fas fa-exchange-alt text-xs")),
			),
		),

		// 2. Info Box (Default)
		Div(ID(fmt.Sprintf("info-box-%d", slot)), Class("text-center w-full"),
			H2(Class("text-sm md:text-xl font-black text-slate-900 uppercase leading-tight mb-1 truncate px-2"), Text(d.ModelName)),
			P(Class("text-blue-700 font-bold text-xs md:text-base"), Text(d.Specifications.Misc.Price)),
		),

		// 3. Change Device Box (Hidden by default)
		Div(ID(fmt.Sprintf("change-box-%d", slot)), Class("hidden w-full absolute top-24 left-0 z-50 bg-white p-2 shadow-xl border border-gray-200"),
			Input(
				Type("text"),
				Placeholder("Type to search..."),
				Class("w-full border border-gray-300 p-2 text-sm focus:border-blue-600 outline-none"),
				Attr("hx-get", fmt.Sprintf("/partials/compare-search?slot=%d&other=%s", slot, otherDeviceSlug)),
				Attr("hx-trigger", "keyup changed delay:200ms"),
				Attr("hx-target", "#"+searchResultID),
				Attr("autofocus", "true"),
			),
			Div(ID(searchResultID), Class("mt-1 max-h-60 overflow-y-auto bg-white border-t border-gray-100")),
			Button(
				Class("w-full mt-2 text-xs text-red-500 font-bold uppercase hover:bg-red-50 py-1"),
				Attr("onclick", fmt.Sprintf("document.getElementById('change-box-%d').classList.add('hidden'); document.getElementById('info-box-%d').classList.remove('hidden')", slot, slot)),
				Text("Cancel"),
			),
		),
	)
}

func compareSection(title string, rows ...Node) Node {
	var validRows []Node
	for _, r := range rows {
		if r != nil {
			validRows = append(validRows, r)
		}
	}
	if len(validRows) == 0 {
		return nil
	}

	return Div(
		H3(Class("text-center text-xs md:text-sm font-extrabold text-white bg-slate-900 py-2 uppercase tracking-widest mb-0"), Text(title)),
		Div(Class("border-x border-b border-gray-200 divide-y divide-gray-100"), Group(validRows)),
	)
}

func compareRow(label, v1, v2 string) Node {
	if strings.TrimSpace(v1) == "" && strings.TrimSpace(v2) == "" {
		return nil
	}
	return Div(Class("grid grid-cols-2 hover:bg-blue-50 transition-colors group"),
		// Column 1
		Div(Class("p-3 md:p-4 border-r border-gray-100 relative overflow-hidden"),
			Span(Class("block text-[9px] md:text-[10px] text-slate-400 font-bold uppercase tracking-wider mb-1"), Text(label)),
			Span(Class("block text-xs md:text-sm text-slate-900 font-medium leading-relaxed line-clamp-3"), Text(CheckEmpty(v1))),
		),
		// Column 2
		Div(Class("p-3 md:p-4 relative overflow-hidden"),
			Span(Class("md:hidden block text-[9px] text-slate-400 font-bold uppercase tracking-wider mb-1"), Text(label)), // Label for mobile
			Span(Class("block text-xs md:text-sm text-slate-900 font-medium leading-relaxed line-clamp-3"), Text(CheckEmpty(v2))),
		),
	)
}
