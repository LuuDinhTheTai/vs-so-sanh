package page

import (
	"fmt"
	"log/slog"
	"vs-so-sanh/internal/brand"
	"vs-so-sanh/internal/brand/delivery/dto"
	"vs-so-sanh/internal/device"
	"vs-so-sanh/internal/model"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HomePage(brandUseCase brand.UseCase, deviceUseCase device.UseCase) Node {
	brands, err := brandUseCase.FindTop20()
	if err != nil {
		slog.Error("Error fetching top 10 brands: ", err)
		brands = []dto.BrandResponse{
			{Name: "Brand 1"},
			{Name: "Brand 2"},
		}
	}

	devices, err := deviceUseCase.FindTop20()
	if err != nil {
		slog.Error("Error fetching top 10 devices: ", err)
		devices = []model.Device{}
	}

	return shared.Page(
		"Vs-SoSanh",

		Div(Class("grid grid-cols-1 lg:grid-cols-10 gap-6"),

			Aside(Class("lg:col-span-3 space-y-6"),
				PhoneFinderWidget(brands),
				Top10Widget(),
				//LatestDevicesWidget(),
			),

			Main(Class("lg:col-span-7 space-y-6"),
				SearchBar(),
				//FeaturedSection(),
				HotDeviceSection(devices),
			),
		),
	)
}

// --- COMPONENTS CON (UI WIDGETS) ---

func PhoneFinderWidget(brands []dto.BrandResponse) Node {
	// Xóa "shadow-sm", "rounded-lg". Dùng "border" và "rounded-sm"
	return Div(Class("bg-white p-5 border border-gray-200 rounded-sm"),
		// Tiêu đề phẳng, dùng chữ đậm để nhấn mạnh
		H3(Class("text-sm font-extrabold text-gray-900 mb-4 uppercase tracking-wider"),
			I(Class("fas fa-search mr-2 text-indigo-600")), Text("Phone Finder"),
		),
		Div(Class("grid grid-cols-2 gap-y-2 gap-x-4 text-sm"),
			Maps(brands, func(b dto.BrandResponse) Node {
				// Link phẳng, bỏ underline mặc định, chỉ đổi màu khi hover
				return A(Href("#"), Class("text-gray-600 hover:text-indigo-600 font-medium transition-colors"), Text(b.Name))
			}),
		),
		Div(Class("mt-5 pt-3 border-t border-gray-100 text-center"),
			// Nút dạng Text phẳng
			A(Href("#"), Class("text-xs font-bold text-indigo-600 uppercase hover:text-indigo-800 tracking-wide"),
				I(Class("fas fa-list mr-1")), Text("View All Brands"),
			),
		),
	)
}

func Top10Widget() Node {
	type PhoneRank struct {
		Rank int
		Name string
		Hits string
	}
	phones := []PhoneRank{
		{1, "Samsung Galaxy S25 Ultra", "48,291"},
		{2, "Redmi Note 14 Pro+", "35,102"},
		{3, "iPhone 16 Pro Max", "29,554"},
		{4, "Samsung Galaxy A55", "25,100"},
		{5, "Google Pixel 9 Pro", "21,040"},
		{6, "Xiaomi 15", "19,880"},
		{7, "Sony Xperia 1 VI", "15,200"},
		{8, "OnePlus 13", "14,300"},
		{9, "Honor Magic7 Pro", "12,100"},
		{10, "Poco F6", "10,500"},
	}

	return Div(Class("bg-white p-5 border border-gray-200 rounded-sm"),
		H3(Class("text-sm font-extrabold text-gray-900 mb-1 uppercase tracking-wider"), Text("Trending Now")),
		P(Class("text-xs text-gray-500 mb-4"), Text("Daily interest statistics")),

		Table(Class("w-full text-sm text-left"),
			TBody(
				Map(phones, func(p PhoneRank) Node {
					return Tr(Class("border-b border-gray-100 last:border-0 hover:bg-gray-50 transition-colors"),
						// Số thứ tự to, đậm, màu nhạt (Flat Typo)
						Td(Class("py-3 font-black text-gray-300 w-8 text-lg leading-none"),
							Text(fmt.Sprintf("%d", p.Rank)),
						),
						Td(Class("py-2"),
							A(Href("#"), Class("font-bold text-gray-800 hover:text-indigo-600 block"), Text(p.Name)),
							Span(Class("text-xs font-mono text-gray-400"), Text(p.Hits+" hits")),
						),
					)
				}),
			),
		),
	)
}

func SearchBar() Node {
	return Div(Class("flex shadow-sm"),
		Input(Type("text"), Placeholder("Find your device..."),
			Class("grow bg-white px-4 py-3 text-gray-900 border border-gray-300 focus:border-blue-500 outline-none placeholder-gray-400 font-medium"),
		),
		Button(Class("bg-slate-900 text-white px-6 py-3 font-bold uppercase tracking-wide hover:bg-slate-800 transition-colors"),
			Text("Search"),
		),
	)
}

func HotDeviceSection(devices []model.Device) Node {
	return Div(Class("bg-white p-5 border border-gray-200 rounded-sm mt-8"), // Bỏ shadow
		// Header style phẳng: Dùng border trái đậm
		H3(Class("text-xl font-black text-gray-900 mb-6 border-l-8 border-indigo-600 pl-4 uppercase"), Text("Hot Phones")),
		Div(Class("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"), // Tăng gap
			Maps(devices, func(n model.Device) Node {
				// Card sản phẩm phẳng
				return A(Href("/devices/"+n.ModelName), Class("group block border border-gray-200 p-4 hover:border-indigo-600 transition-colors bg-white"), // Hover đổi màu border thay vì nổi lên
					Div(Class("h-40 bg-white mb-4 flex items-center justify-center p-2"), // Nền trắng thay vì xám
						Img(
							Class("max-h-full max-w-full object-contain"),
							Src(n.ImageUrl),
							Alt(n.ModelName),
						),
					),
					H4(Class("font-bold text-base text-gray-900 group-hover:text-indigo-600 text-center"), Text(n.ModelName)),
				)
			}),
		),
	)
}

func LatestDevicesWidget() Node {
	devices := []struct{ Name, Img string }{
		{"ZTE Pad", "https://fdn2.gsmarena.com/vv/bigpic/zte-pad.jpg"},
		{"vivo S50", "https://fdn2.gsmarena.com/vv/bigpic/vivo-s20-pro.jpg"},
		{"Honor X8d", "https://fdn2.gsmarena.com/vv/bigpic/honor-x8b.jpg"},
		{"Oppo Reno 15c", "https://fdn2.gsmarena.com/vv/bigpic/oppo-reno11-f-5g.jpg"},
	}

	return Div(Class("bg-white p-4 shadow-sm rounded-lg border border-gray-200"),
		H3(Class("text-lg font-bold text-gray-700 mb-3 border-b pb-2 uppercase"), Text("Latest Devices")),
		Div(Class("grid grid-cols-2 gap-4"),
			Maps(devices, func(d struct{ Name, Img string }) Node {
				return Div(Class("text-center group cursor-pointer"),
					Img(Src(d.Img), Alt(d.Name), Class("mx-auto h-24 object-contain mb-1 transition-transform group-hover:scale-105")),
					Div(Class("text-xs font-medium text-gray-700 group-hover:text-indigo-600"), Text(d.Name)),
				)
			}),
		),
	)
}

func FeaturedSection() Node {
	return Div(Class("grid grid-cols-1 md:grid-cols-2 gap-4"),

		Div(Class("relative h-64 bg-gray-900 rounded-lg overflow-hidden group cursor-pointer col-span-1 md:col-span-2 lg:col-span-1"),
			Img(Src("https://fdn.gsmarena.com/imgroot/news/24/12/huawei-mate-x7-review-hands-on/-727/gsmarena_001.jpg"), Class("w-full h-full object-cover opacity-80 group-hover:opacity-100 transition-opacity")),
			Div(Class("absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black p-4"),
				H2(Class("text-white text-xl font-bold leading-tight"), Text("Huawei Mate X7 hands-on review")),
				P(Class("text-gray-300 text-sm mt-1"), Text("22 hours ago • 19 comments")),
			),
		),

		Div(Class("flex flex-col space-y-4"),
			FeaturedSmallCard("Honor Magic8 Lite review", "https://fdn.gsmarena.com/imgroot/reviews/24/honor-magic6-lite/-727/gsmarena_001.jpg"),
			FeaturedSmallCard("Nothing Phone (3a) Lite review", "https://fdn.gsmarena.com/imgroot/reviews/24/nothing-phone-2a/-727/gsmarena_001.jpg"),
		),
	)
}

func FeaturedSmallCard(title, imgUrl string) Node {
	return Div(Class("relative h-28 bg-gray-900 rounded-lg overflow-hidden group cursor-pointer flex-grow"),
		Img(Src(imgUrl), Class("w-full h-full object-cover opacity-80 group-hover:opacity-100 transition-opacity")),
		Div(Class("absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black p-2"),
			H4(Class("text-white font-bold text-sm"), Text(title)),
		),
	)
}

func Maps[T any](data []T, f func(T) Node) Node {
	nodes := make([]Node, len(data))
	for i, item := range data {
		nodes[i] = f(item)
	}
	return Group(nodes)
}
