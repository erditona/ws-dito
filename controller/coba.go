package controller

import (
	cek "github.com/aiteung/presensi"
	"github.com/erditona/ws-dito/config"
	"github.com/gofiber/fiber/v2"

	"net/http"

	model "github.com/erditona/be_pmb/model"
	module "github.com/erditona/be_pmb/module"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/erditona/ws-dito",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

// func Homepage(c *fiber.Ctx) error {
// 	ipaddr := musik.GetIPaddress()
// 	return c.JSON(ipaddr)
// }

func GetPresensi(c *fiber.Ctx) error {
	n1 := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(n1)
}

func GetPendaftaran(c *fiber.Ctx) error {
	ps := module.GetPendaftaranFromKTP(320132321321,config.Ulbimongoconn,"pendaftaran_maba")
	return c.JSON(ps)
}

func GetCamaba(c *fiber.Ctx) error {
	ps := module.GetCamabaFromPhoneNumber("085725722483",config.Ulbimongoconn,"daftar_camaba")
	return c.JSON(ps)
}

func GetJurusan(c *fiber.Ctx) error {
	ps := module.GetJurusanFromKDJurusan("D3TI",config.Ulbimongoconn,"daftar_jurusan")
	return c.JSON(ps)
}

func GetSekolah(c *fiber.Ctx) error {
	ps := module.GetDaftarSekolahFromKDSekolah(4,config.Ulbimongoconn,"daftar_sekolah")
	return c.JSON(ps)
}



func InsertPendaftaran(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pendaftaran model.Pendaftaran
	if err := c.BodyParser(&pendaftaran); err != nil {
		return err
	}
	insertedID := module.InsertPendaftaran(db, "pendaftaran_maba",
		pendaftaran.KDPendaftar,
		pendaftaran.Biodata,
		pendaftaran.AsalSekolah,
		pendaftaran.Jurusan,
		pendaftaran.Jalur,
		pendaftaran.AlUlbi,
		pendaftaran.AlJurusan)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertCamaba(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var camaba model.Camaba
	if err := c.BodyParser(&camaba); err != nil {
		return err
	}
	insertedID := module.InsertDaftarCamaba(db, "daftar_camaba",
		camaba.Ktp,
		camaba.Nama,
		camaba.Phone_number,
		camaba.Address)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertJurusan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jurusan model.Jurusan
	if err := c.BodyParser(&jurusan); err != nil {
		return err
	}
	insertedID := module.InsertDaftarJurusan(db, "daftar_jurusan",
		jurusan.KDJurusan,
		jurusan.Nama,
		jurusan.Jenjang)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertSekolah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var sekolah model.DaftarSekolah
	if err := c.BodyParser(&sekolah); err != nil {
		return err
	}
	insertedID := module.InsertDaftarSekolah(db, "daftar_sekolah", 
		sekolah.KDSekolah,
		sekolah.Nama,
		sekolah.Address,
		sekolah.Phone_number)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

