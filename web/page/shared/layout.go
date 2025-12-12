package shared

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func Page(title string, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
			// 1. Thêm HTMX
			Script(Src("https://unpkg.com/htmx.org@1.9.10")),
		},
		Body: []Node{
			Class("bg-gray-50 min-h-screen flex flex-col"), // Sửa lại background cho sáng sủa hơn
			PageHeader(),
			Div(Class("grow py-8"),
				Container(false, // Tắt padding mặc định để tự control
					Group(children),
				),
			),
			PageFooter(),
			// 2. Nơi chứa thanh so sánh (Floating Bar) sẽ được HTMX cập nhật
			Div(ID("compare-tray")),
		},
	})
}
