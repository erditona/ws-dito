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

	inimodellatihan "github.com/indrariksa/be_presensi/model"
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

// func InsertPendaftaran(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var pendaftaran model.Pendaftaran
// 	if err := c.BodyParser(&pendaftaran); err != nil {
// 		return err
// 	}
// 	insertedID := module.InsertPendaftaran(db, "pendaftaran_maba",
// 		pendaftaran.KDPendaftar,
// 		pendaftaran.Biodata,
// 		pendaftaran.AsalSekolah,
// 		pendaftaran.Jurusan,
// 		pendaftaran.Jalur,
// 		pendaftaran.AlUlbi,
// 		pendaftaran.AlJurusan)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }
func InsertPendaftaran(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pendaftaran model.Pendaftaran
	if err := c.BodyParser(&pendaftaran); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := module.InsertPendaftaran(db, "pendaftaran_maba",
		pendaftaran.KDPendaftar,
		pendaftaran.Biodata,
		pendaftaran.AsalSekolah,
		pendaftaran.Jurusan,
		pendaftaran.Jalur,
		pendaftaran.AlUlbi,
		pendaftaran.AlJurusan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// func InsertCamaba(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var camaba model.Camaba
// 	if err := c.BodyParser(&camaba); err != nil {
// 		return err
// 	}
// 	insertedID := module.InsertDaftarCamaba(db, "daftar_camaba",
// 		camaba.Ktp,
// 		camaba.Nama,
// 		camaba.Phone_number,
// 		camaba.Address)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }
func InsertCamaba(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var camaba model.Camaba
	if err := c.BodyParser(&camaba); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := module.InsertCamaba(db, "daftar_camaba",
		camaba.Ktp,
		camaba.Nama,
		camaba.Phone_number,
		camaba.Address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// func InsertJurusan(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var jurusan model.Jurusan
// 	if err := c.BodyParser(&jurusan); err != nil {
// 		return err
// 	}
// 	insertedID := module.InsertDaftarJurusan(db, "daftar_jurusan",
// 		jurusan.KDJurusan,
// 		jurusan.Nama,
// 		jurusan.Jenjang)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }
func InsertJurusan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jurusan model.Jurusan
	if err := c.BodyParser(&jurusan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := module.InsertJurusan(db, "daftar_jurusan",
	jurusan.KDJurusan,
	jurusan.Nama,
	jurusan.Jenjang)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// func InsertSekolah(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn
// 	var sekolah model.DaftarSekolah
// 	if err := c.BodyParser(&sekolah); err != nil {
// 		return err
// 	}
// 	insertedID := module.InsertDaftarSekolah(db, "daftar_sekolah", 
// 		sekolah.KDSekolah,
// 		sekolah.Nama,
// 		sekolah.Phone_number,
// 		sekolah.Address)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "data berhasil disimpan.",
// 		"inserted_id": insertedID,
// 	})
// }
func InsertSekolah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var sekolah model.DaftarSekolah
	if err := c.BodyParser(&sekolah); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := module.InsertSekolah(db, "daftar_sekolah", 
		sekolah.KDSekolah,
		sekolah.Nama,
		sekolah.Phone_number,
		sekolah.Address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
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

//LatihanWeek11 - Swegger
// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
//Latihan Week-6
//Fungsi get all data tanpa filter
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn2, "presensi")
	return c.JSON(ps)
}

// GetPresensiID - Swegger 
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
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

//InsertData - Swegger
// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
//week9
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

//Week10
//Function Update
// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

//Function Delete
// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
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

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn2, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}


//PMB
//GetPendaftaranID
func GetPendaftaranID(c *fiber.Ctx) error {
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
	ps, err := module.GetPendaftaranFromID(objID, config.Ulbimongoconn, "pendaftaran_maba")
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

//GetJurusanID
func GetJurusanID(c *fiber.Ctx) error {
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
	ps, err := module.GetJurusanFromID(objID, config.Ulbimongoconn, "daftar_jurusan")
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

//GetSekolahID
func GetSekolahID(c *fiber.Ctx) error {
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
	ps, err := module.GetSekolahFromID(objID, config.Ulbimongoconn, "daftar_sekolah")
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

//GetCamabaID
func GetCamabaID(c *fiber.Ctx) error {
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
	ps, err := module.GetCamabaFromID(objID, config.Ulbimongoconn, "daftar_camaba")
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

//Update-Delete

//Sekolah
func UpdateDataSekolah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a sekolah object
	var sekolah model.DaftarSekolah
	if err := c.BodyParser(&sekolah); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdateSekolah(db, "daftar_sekolah",
		objectID,
		sekolah.KDSekolah,
		sekolah.Nama,
		sekolah.Phone_number,
		sekolah.Address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}
func DeleteSekolahByID(c *fiber.Ctx) error {
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

	err = module.DeleteSekolahByID(objID, config.Ulbimongoconn, "daftar_sekolah")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

//Jurusan
func UpdateDataJurusan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a jurusan object
	var jurusan model.Jurusan
	if err := c.BodyParser(&jurusan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdateJurusan(db, "daftar_jurusan",
		objectID,
		jurusan.KDJurusan,
		jurusan.Nama,
		jurusan.Jenjang)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}
func DeleteJurusanByID(c *fiber.Ctx) error {
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

	err = module.DeleteJurusanByID(objID, config.Ulbimongoconn, "daftar_jurusan")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}