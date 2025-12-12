package page

import (
	"fmt"
	model2 "vs-so-sanh/internal/model"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HomePage(products []model2.Phone) Node {
	return shared.Page("Danh sách sản phẩm",
		Div(Class("text-center mb-10"),
			H1(Class("text-4xl font-bold text-gray-900 mb-4"), Text("So sánh mọi thứ")),
			P(Class("text-lg text-gray-600"), Text("Chọn sản phẩm để thấy sự khác biệt")),
		),

		// Grid hiển thị sản phẩm
		Div(Class("grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"),
			Map(products, func(p model2.Phone) Node {
				return ProductCard(p)
			}),
		),
	)
}

func ProductCard(p model2.Phone) Node {
	return Div(Class("bg-white rounded-xl shadow-sm hover:shadow-md transition border border-gray-100 overflow-hidden flex flex-col"),
		// Ảnh sản phẩm
		Div(Class("h-48 bg-gray-200 w-full overflow-hidden")), //Img(Src(p.Image), Alt(p.Name), Class("w-full h-full object-cover transform hover:scale-105 transition duration-500")),

		// Nội dung
		Div(Class("p-5 flex flex-col flex-grow"),
			H3(Class("text-lg font-bold text-gray-900 mb-2"), Text(p.Brand)),
			P(Class("text-indigo-600 font-bold text-xl mb-4"), Text(fmt.Sprintf("%v đ", p.Brand))), // Format giá tuỳ ý

			Div(Class("mt-auto flex gap-2"),
				// Link chi tiết
				A(Href(fmt.Sprintf("/product/%d", p.Brand)),
					Class("flex-1 block text-center py-2 px-4 rounded border border-gray-300 text-gray-700 hover:bg-gray-50 text-sm font-medium"),
					Text("Chi tiết"),
				),
				// Nút HTMX thêm vào so sánh
				Button(
					Class("flex-1 py-2 px-4 rounded bg-indigo-600 text-white hover:bg-indigo-700 text-sm font-medium transition"),
					Attr("hx-post", fmt.Sprintf("/compare/toggle/%d", p.Brand)), // Gọi API
					Attr("hx-target", "#compare-tray"),                          // Cập nhật vào cái khay ở layout
					Attr("hx-swap", "innerHTML"),
					Text("So sánh +"),
				),
			),
		),
	)
}
