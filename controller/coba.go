package controller

import (
	"errors"
	"fmt"

	cek "github.com/aiteung/presensi"
	"github.com/erditona/ws-dito/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"

	inimodullatihan "github.com/indrariksa/be_presensi/module"

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

//GetAllFunction

func GetPresensi(c *fiber.Ctx) error {
	n1 := cek.GetPresensiCurrentMonth(config.Ulbimongoconn2)
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

//InsertFunction

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
		sekolah.Phone_number,
		sekolah.Address)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

//GetAllFunction

func GetAllPendaftaran(c *fiber.Ctx) error {
	ps := module.GetAllPendaftaran(config.Ulbimongoconn,"pendaftaran_maba")
	return c.JSON(ps)
}

func GetAllJurusan(c *fiber.Ctx) error {
	ps := module.GetAllJurusan(config.Ulbimongoconn,"daftar_jurusan")
	return c.JSON(ps)
}

func GetAllSekolah(c *fiber.Ctx) error {
	ps := module.GetAllSekolah(config.Ulbimongoconn,"daftar_sekolah")
	return c.JSON(ps)
}

func GetAllCamaba(c *fiber.Ctx) error {
	ps := module.GetAllCamaba(config.Ulbimongoconn,"daftar_camaba")
	return c.JSON(ps)
}



//Latihan Week-6
//Fungsi get all data tanpa filter
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn2, "presensi")
	return c.JSON(ps)
}

//Fungsi get by id
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn2, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}
