package page

import (
	"strings"

	model "vs-so-sanh/internal/model" // Giả sử alias là model
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DetailsPage(device *model.Device) Node {
	specs := device.Specifications

	return shared.Page(device.ModelName,
		Div(Class("max-w-6xl mx-auto px-4 py-8"),
			// 1. Breadcrumb
			Nav(Class("mb-6"),
				A(Href("/"), Class("text-indigo-600 hover:text-indigo-800 font-medium transition-colors"), Text("← Quay lại danh sách")),
			),

			Div(Class("bg-white rounded-2xl shadow-xl overflow-hidden"),
				// 2. Phần Header & Ảnh (Top Section)
				Div(Class("md:flex border-b border-gray-200"),
					// Cột trái: Ảnh
					Div(Class("md:w-1/3 bg-gray-50 p-8 flex items-center justify-center border-r border-gray-100"),
						Img(Src(device.ImageUrl), Alt(device.ModelName), Class("max-h-96 w-auto object-contain hover:scale-105 transition-transform duration-300")),
					),

					// Cột phải: Thông tin tóm tắt
					Div(Class("md:w-2/3 p-8 md:p-10"),
						H1(Class("text-4xl font-extrabold text-gray-900 mb-2"), Text(device.ModelName)),
						P(Class("text-2xl text-indigo-600 font-bold mb-6"), Text(specs.Misc.Price)),

						// Highlights (Tóm tắt nhanh)
						Div(Class("grid grid-cols-2 gap-4 mb-6"),
							highlightBox("Chipset", specs.Platform.Chipset),
							highlightBox("Màn hình", specs.Display.MainSize),
							highlightBox("RAM/Bộ nhớ", specs.Memory.Internal),
							highlightBox("Pin", specs.Battery.Type),
						),

						// Nút hành động (Ví dụ)
						Button(Class("bg-indigo-600 text-white px-6 py-3 rounded-lg font-bold hover:bg-indigo-700 transition"), Text("So sánh ngay")),
					),
				),

				// 3. Phần Chi tiết thông số (Detail Specs)
				Div(Class("p-8 bg-gray-50"),
					H2(Class("text-2xl font-bold text-gray-800 mb-6 text-center uppercase tracking-wide"), Text("Thông số kỹ thuật chi tiết")),

					Div(Class("grid md:grid-cols-2 gap-8"),
						// Cột 1
						Div(Class("space-y-8"),
							specSection("Mạng & Kết nối",
								specRow("Công nghệ", specs.Network.Technology),
								specRow("Tốc độ", specs.Network.Speed),
								specRow("WLAN", specs.Comms.WLAN),
								specRow("Bluetooth", specs.Comms.Bluetooth),
							),
							specSection("Thiết kế (Body)",
								specRow("Kích thước (Mở)", specs.Body.DimensionsUnfolded),
								specRow("Kích thước (Gập)", specs.Body.DimensionsFolded),
								specRow("Trọng lượng", specs.Body.Weight),
								specRow("Build", specs.Body.Build),
								specRow("SIM", specs.Body.SIM),
							),
							specSection("Màn hình",
								specRow("Loại", specs.Display.MainType),
								specRow("Kích thước", specs.Display.MainSize),
								specRow("Độ phân giải", specs.Display.MainResolution),
								specRow("Màn hình phụ", specs.Display.CoverSize),
							),
						),

						// Cột 2
						Div(Class("space-y-8"),
							specSection("Cấu hình (Platform)",
								specRow("Hệ điều hành", specs.Platform.OS),
								specRow("Chipset", specs.Platform.Chipset),
								specRow("CPU", specs.Platform.CPU),
								specRow("GPU", specs.Platform.GPU),
								specRow("Thẻ nhớ", specs.Memory.CardSlot),
								specRow("Bộ nhớ trong", specs.Memory.Internal),
							),
							specSection("Camera",
								specRow("Camera sau", specs.MainCamera.Single),
								specRow("Video sau", specs.MainCamera.Video),
								specRow("Camera trước", specs.SelfieCamera.Single),
								specRow("Video trước", specs.SelfieCamera.Video),
							),
							specSection("Pin & Sạc",
								specRow("Loại pin", specs.Battery.Type),
								specRow("Sạc", specs.Battery.Charging),
							),
							specSection("Tính năng khác",
								specRow("Cảm biến", specs.Features.Sensors),
								specRow("Màu sắc", specs.Misc.Colors),
							),
						),
					),
				),
			),
		),
	)
}

// --- Helper Functions ---

// specSection: Tạo một nhóm thông số (ví dụ: Mạng, Màn hình)
func specSection(title string, rows ...Node) Node {
	return Div(Class("bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden"),
		H3(Class("bg-gray-100 px-6 py-3 font-bold text-gray-700 border-b border-gray-200"), Text(title)),
		Div(Class("p-4 space-y-3"), Group(rows)),
	)
}

// specRow: Tạo một dòng thông số (Label: Value). Nếu Value rỗng sẽ ẩn đi.
func specRow(label, value string) Node {
	if strings.TrimSpace(value) == "" {
		return nil // Gomponents sẽ bỏ qua Node này nếu trả về nil
	}
	return Div(Class("grid grid-cols-3 gap-2 text-sm border-b border-gray-50 last:border-0 pb-2 last:pb-0"),
		Span(Class("col-span-1 text-gray-500 font-medium"), Text(label)),
		Span(Class("col-span-2 text-gray-900"), Text(value)),
	)
}

// highlightBox: Tạo hộp thông tin nổi bật ở phần header
func highlightBox(label, value string) Node {
	if value == "" {
		return nil
	}
	return Div(Class("bg-indigo-50 p-3 rounded-lg border border-indigo-100"),
		Span(Class("block text-xs text-indigo-500 uppercase font-semibold"), Text(label)),
		Span(Class("block text-sm font-bold text-gray-800 truncate"), Text(value)),
	)
}
