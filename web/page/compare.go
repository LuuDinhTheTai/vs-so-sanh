package page

import (
	"fmt"
	model2 "vs-so-sanh/internal/model"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ComparePage(products []model2.Phone, allSpecKeys []string) Node {
	if len(products) == 0 {
		return shared.Page("So sánh",
			Div(Class("text-center py-20"),
				H2(Class("text-2xl font-bold text-gray-700"), Text("Chưa có sản phẩm nào được chọn")),
				A(Href("/"), Class("mt-4 inline-block text-indigo-600 hover:underline"), Text("Chọn sản phẩm ngay")),
			),
		)
	}

	return shared.Page("Bảng So Sánh",
		Div(Class("overflow-x-auto pb-10"),
			Table(Class("w-full min-w-[800px] border-collapse text-left"),
				// Header: Ảnh và Tên
				THead(
					Tr(
						Th(Class("p-4 border-b-2 border-gray-100 min-w-[200px]"), Text("Tiêu chí")),
						Maps(products, func(p model2.Phone) Node {
							return Th(Class("p-4 border-b-2 border-gray-100 w-1/4 align-bottom"),
								//Img(Src(p.Image), Class("w-24 h-24 object-contain mx-auto mb-4")),
								Div(Class("text-lg font-bold text-center"), Text(p.BrandName)),
								Div(Class("text-indigo-600 text-center font-normal"), Text(fmt.Sprintf("%v đ", p.BrandName))),
							)
						}),
					),
				),
				// Body: Các thông số
				TBody(
					Maps(allSpecKeys, func(key string) Node {
						return Tr(Class("hover:bg-gray-50 transition"),
							Td(Class("p-4 border-b border-gray-100 font-medium text-gray-500"), Text(key)),
							Maps(products, func(p model2.Phone) Node {
								// Giả sử bạn có hàm lấy value từ spec JSON, hoặc xử lý ở controller
								// Ở đây tôi giả định Product có method GetSpec(key)
								val := "N/A" // Placeholder
								return Td(Class("p-4 border-b border-gray-100 text-center text-gray-900"), Text(val))
							}),
						)
					}),
					// Nút hành động cuối bảng
					Tr(
						Td(Class("p-4"), Text("")),
						Maps(products, func(p model2.Phone) Node {
							return Td(Class("p-4 text-center"),
								A(Href(fmt.Sprintf("/product/%d", p.BrandName)), Class("text-indigo-600 hover:underline text-sm"), Text("Xem chi tiết")),
							)
						}),
					),
				),
			),
		),
	)
}
