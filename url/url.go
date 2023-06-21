package url

import (
	"github.com/erditona/ws-dito/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	// page.Get("/presensi", controller.GetPresensi)
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                   

	page.Get("/daftar-camaba", controller.GetCamaba) 
	page.Get("/daftar-jurusan", controller.GetJurusan) 
	page.Get("/daftar-sekolah", controller.GetSekolah) 
	// page.Post("/insrt-jurusan", controller.InsertJurusan) 
	// page.Post("/insrt-sekolah", controller.InsertSekolah) 
	// page.Post("/insrt-camaba", controller.InsertCamaba) 
	// page.Post("/insrt-pendaftaran", controller.InsertPendaftaran) 
	
	//EndPoint PMB
	//GetAll
	page.Get("/pendaftaran", controller.GetAllPendaftaran) 
	page.Get("/jurusan", controller.GetAllJurusan) 
	page.Get("/sekolah", controller.GetAllSekolah) 
	page.Get("/camaba", controller.GetAllCamaba) 
	page.Get("/user", controller.GetAllUser) 
	//GetFromID
	page.Get("/pendaftaran/:id", controller.GetPendaftaranID) 
	page.Get("/pendaftaran/kd/:kdpendaftar", controller.GetPendaftaranKDPendaftar) 
	page.Get("/jurusan/:id", controller.GetJurusanID) 
	page.Get("/sekolah/:id", controller.GetSekolahID) 
	page.Get("/camaba/:id", controller.GetCamabaID) 
	page.Get("/user/:id", controller.GetUserID) 
	//Insert
	page.Post("/ins-pendaftaran", controller.InsertPendaftaran)
	page.Post("/ins-jurusan", controller.InsertJurusan)
	page.Post("/ins-camaba", controller.InsertCamaba)
	page.Post("/ins-sekolah", controller.InsertSekolah)
	//Update
	page.Put("/upd-status-pendaftaran/:id", controller.UpdateStatusPendaftaran)
	page.Put("/upd-pendaftaran/:id", controller.UpdateDataPendaftaran)
	page.Put("/upd-sekolah/:id", controller.UpdateDataSekolah)
	page.Put("/upd-jurusan/:id", controller.UpdateDataJurusan)
	//Delete
	page.Delete("/delete-pendaftaran/:id", controller.DeletePendaftaranByID)
	page.Delete("/delete-sekolah/:id", controller.DeleteSekolahByID)
	page.Delete("/delete-jurusan/:id", controller.DeleteJurusanByID)
	//Login-SignUp
	page.Post("/sign-up", controller.SignUp)
	page.Post("/sign-in", controller.SignIn)
	
	//latihanWeek6
	page.Get("/presensi", controller.GetAllPresensi) //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	//latihanWeek9
	page.Post("/ins", controller.InsertData)
	//LatihanWeek10
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	//LatihanWeek11-Swagger
	page.Get("/docs/*", swagger.HandlerDefault)
}
