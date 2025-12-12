package page

import (
	"fmt"
	model2 "vs-so-sanh/internal/model"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DetailsPage(p model2.Phone, specs map[string]string) Node {
	return shared.Page(p.BrandName,
		Div(Class("max-w-5xl mx-auto"),
			// Breadcrumb
			A(Href("/"), Class("text-indigo-600 hover:underline mb-6 inline-block"), Text("← Quay lại danh sách")),

			Div(Class("bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden"),
				Div(Class("md:flex"),
					// Ảnh lớn bên trái
					Div(Class("md:w-1/2 bg-gray-50 p-8 flex items-center justify-center")), //Img(Src(p.Image), Class("max-h-96 w-auto object-contain mix-blend-multiply")),

					// Thông tin bên phải
					Div(Class("md:w-1/2 p-8 md:p-12"),
						H1(Class("text-3xl font-bold text-gray-900 mb-2"), Text(p.BrandName)),
						P(Class("text-2xl text-indigo-600 font-bold mb-6"), Text(fmt.Sprintf("%v đ", p.BrandName))),

						H3(Class("text-sm font-semibold text-gray-500 uppercase tracking-wider mb-4"), Text("Thông số kỹ thuật")),

						// Bảng thông số mini
						Div(Class("space-y-3"),
							Maps(getKeys(specs), func(k string) Node {
								return Div(Class("flex justify-between border-b border-gray-100 pb-2"),
									Span(Class("text-gray-600"), Text(k)),
									Span(Class("font-medium text-gray-900"), Text(specs[k])),
								)
							}),
						),
					),
				),
			),
		),
	)
}

// Helper lấy key từ map (vì map không có thứ tự)
func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
