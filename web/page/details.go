package page

import (
	"strings"
	"vs-so-sanh/internal/device/delivery/dto"

	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DetailsPage(device *dto.DeviceResponse) Node {
	specs := device.Specifications

	return shared.Page(device.ModelName,
		Div(Class("max-w-7xl mx-auto px-4 py-8"),

			// 1. Breadcrumb (Điều hướng phẳng)
			Nav(Class("flex items-center text-sm font-medium text-gray-500 mb-8"),
				A(Href("/"), Class("hover:text-indigo-600 transition-colors"), Text("Home")),
				Span(Class("mx-2 text-gray-300"), Text("/")),
				Span(Class("text-gray-900"), Text(device.ModelName)),
			),

			// 2. HERO SECTION: Chia 2 cột (Ảnh & Thông tin tóm tắt)
			Div(Class("grid grid-cols-1 lg:grid-cols-12 gap-8 mb-16"),

				// CỘT TRÁI: Ảnh sản phẩm (Bordered Box)
				Div(Class("lg:col-span-5 bg-white border border-gray-200 p-8 flex items-center justify-center relative"),
					// Label trạng thái (Ví dụ)
					Span(Class("absolute top-0 left-0 bg-slate-800 text-white text-xs font-bold px-3 py-1 uppercase tracking-wider"), Text("Official")),
					Img(Src(device.ImageUrl), Alt(device.ModelName), Class("max-h-96 w-auto object-contain")),
				),

				// CỘT PHẢI: Thông tin chính
				Div(Class("lg:col-span-7 flex flex-col justify-center"),
					H1(Class("text-4xl lg:text-5xl font-black text-gray-900 mb-2 uppercase tracking-tight leading-none"), Text(device.ModelName)),

					// Giá bán & Actions
					Div(Class("flex flex-wrap items-end gap-4 mb-8 border-b border-gray-100 pb-6"),
						P(Class("text-3xl font-bold text-blue-700 leading-none"), Text(specs.Misc.Price)),
						Span(Class("text-sm text-gray-400 font-medium mb-1"), Text("Estimated price")),
					),

					// Highlights Grid (4 ô thông số quan trọng)
					Div(Class("grid grid-cols-2 gap-4 mb-8"),
						// Sử dụng techBadge để hiển thị thông minh hơn
						flatHighlightBox("Chipset", specs.Platform.Chipset),
						flatHighlightBox("Display", specs.Display.MainSize),
						flatHighlightBox("Memory", specs.Memory.Internal),
						flatHighlightBox("Battery", specs.Battery.Type),
					),

					// Action Buttons (Flat Style - Vuông vức)
					Div(Class("flex gap-4"),
						Button(
							Class("bg-slate-900 text-white border-2 border-transparent "+
								"px-6 py-4 font-bold uppercase tracking-widest text-sm "+
								"transition-all duration-300 ease-in-out "+
								"hover:bg-white hover:text-slate-900 hover:border-slate-900"),
							I(Class("fas fa-plus mr-2")),
							Text("Add to Compare"),
						),
						Button(Class("flex-1 border-2 border-gray-900 text-gray-900 px-6 py-4 font-bold uppercase tracking-widest hover:bg-gray-900 hover:text-white transition-colors text-sm"),
							Text("Full Review"),
						),
					),
				),
			),

			// 3. DETAILED SPECS SECTION: Danh sách chi tiết
			Div(Class("border-t-4 border-gray-900 pt-10"),
				Div(Class("flex items-center justify-between mb-8"),
					H2(Class("text-2xl font-black text-gray-900 uppercase tracking-wide"), Text("Technical Specifications")),
					// Nút in hoặc share nhỏ
					Button(Class("text-gray-400 hover:text-indigo-600"), I(Class("fas fa-print"))),
				),

				Div(Class("grid md:grid-cols-2 gap-x-12 gap-y-12"),

					// CỘT 1
					Div(Class("space-y-10"),
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

					// CỘT 2
					Div(Class("space-y-10"),
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
						specSectionFlat("Camera System",
							specRowFlat("Main Camera", specs.MainCamera.Single),
							specRowFlat("Main Video", specs.MainCamera.Video),
							specRowFlat("Selfie Camera", specs.SelfieCamera.Single),
							specRowFlat("Selfie Video", specs.SelfieCamera.Video),
						),
						specSectionFlat("Battery & Charging",
							specRowFlat("Type", specs.Battery.Type),
							specRowFlat("Charging", specs.Battery.Charging),
						),
					),
				),
			),
		),
	)
}

// --- HELPER FUNCTIONS (FLAT DESIGN) ---

func flatHighlightBox(label, value string) Node {
	if value == "" {
		return nil
	}
	return Div(Class("border border-gray-200 p-4 hover:border-blue-600 transition-colors group cursor-default bg-white"),
		Span(Class("block text-[10px] text-gray-500 uppercase font-bold tracking-wider mb-2 group-hover:text-blue-600"), Text(label)),
		Span(Class("block text-sm font-bold text-gray-900 line-clamp-2 leading-snug"), Text(value)),
	)
}

// specSectionFlat: Nhóm thông số, tiêu đề đậm, không card wrapper
func specSectionFlat(title string, rows ...Node) Node {
	return Div(
		// Tiêu đề section: Chữ đậm, màu Indigo, gạch dưới mỏng
		H3(Class("text-lg font-extrabold text-slate-800 mb-4 uppercase border-b-2 border-slate-200 pb-2 inline-block"), Text(title)),
		Div(Class("space-y-0"), Group(rows)),
	)
}

// specRowFlat: Dòng thông số, border bottom mỏng, layout flex
func specRowFlat(label, value string) Node {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	// Layout: Label bên trái (nhỏ, xám), Value bên phải (đậm, đen)
	// Border dưới rất mờ (gray-100) để phân cách
	return Div(Class("grid grid-cols-1 sm:grid-cols-3 gap-2 py-3 border-b border-gray-100 last:border-0"),
		Span(Class("col-span-1 text-sm text-gray-500 font-medium"), Text(label)),
		Span(Class("col-span-2 text-sm text-gray-900 font-semibold leading-relaxed"), Text(value)),
	)
}
