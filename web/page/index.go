package page

import (
	"vs-so-sanh/internal/brand/delivery/dto"
	"vs-so-sanh/internal/model"
	"vs-so-sanh/web/page/shared"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func HomePage(top20Brands []dto.BrandResponse, top20Devices []model.Device) Node {
	return shared.Page(
		"Vs-SoSanh",

		Div(Class("grid grid-cols-1 lg:grid-cols-10 gap-6"),

			Aside(Class("lg:col-span-3 space-y-6"),
				shared.PhoneFinderWidget(top20Brands),
				shared.Top10Widget(),
				//shared.LatestDevicesWidget(),
			),

			Main(Class("lg:col-span-7 space-y-6"),
				shared.SearchBar(),
				//shared.FeaturedSection(),
				shared.HotDeviceSection(top20Devices),
			),
		),
	)
}
