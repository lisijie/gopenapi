gopenapi
==========

腾讯OpenAPI的GO语言版本（未完成）

使用：

	api := gopenapi.NewOpenApi(你的appid, 你的appkey)

	user := gopenapi.NewUser(openid, openkey, platform)

	user.SetApi(api)

	user.GetInfo()
