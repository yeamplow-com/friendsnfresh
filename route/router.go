package route

import (
	"apifnf/api"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	// REGISTER, REGISTER, RESET
	e.POST("/v1/login", api.auth.login)
	e.POST("/v1/register", api.auth.register)
	e.POST("/v1/password/create", api.auth.createp)
	e.POST("/v1/password/reset", api.auth.resetp)

	// INFO USERS
	e.GET("/v1/users/me", api.auth.users.data_users)
	e.PUT("/v1/users/me", api.auth.users.update_users)
	e.PUT("/v1/users/password", api.auth.users.password_users)

	// HOME APP FRIENDS AND FRESH
	e.GET("/v1/banner", api.home.banner)
	e.GET("/v1/katagori/:id", api.home.katagori)
	e.GET("/v1/produk/all", api.home.produk_all)
	e.GET("/v1/produk/:id", api.home.produk)
	e.GET("/v1/produk/best", api.home.produk_best)
	e.GET("/search/:nama", api.home.search_produk)
	e.GET("/search/:id/:nama", api.home.search_produk_k)

	// KERANJANG USER (AUTH)
	e.GET("/v1/cart/me", api.auth.cart.data_cart)
	e.PUT("/v1/cart/:id", api.auth.cart.update_cart)
	e.DELETE("/v1/cart/:id", api.auth.cart.delete_cart)

	// TRANSAKSI
	e.GET("/v1/transaksi/me", api.auth.transaksi.data_transaksi)
	e.POST("/v1/transaksi/add", api.auth.transaksi.add_transaksi)

	// ADDRESS
	e.GET("/v1/address/me", api.auth.address.data_address)
	e.PUT("/v1/address/select/:id", api.auth.address.select_address)
	e.PUT("/v1/address/:id", api.auth.address.update_address)
	e.DELETE("/v1/address/:id", api.auth.address.delete_address)


	return e
}
