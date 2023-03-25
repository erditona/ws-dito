package url

import (
	"github.com/erditona/ws-dito/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	// page.Get("/presensi", controller.GetPresensi)
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/pendaftaran", controller.GetPendaftaran) 
	page.Get("/daftar-camaba", controller.GetCamaba) 
	page.Get("/daftar-jurusan", controller.GetJurusan) 
	page.Get("/daftar-sekolah", controller.GetSekolah) 
	page.Post("/insrt-jurusan", controller.InsertJurusan) 
	page.Post("/insrt-sekolah", controller.InsertSekolah) 
	page.Post("/insrt-camaba", controller.InsertCamaba) 
	page.Post("/insrt-pendaftaran", controller.InsertPendaftaran) 
}
