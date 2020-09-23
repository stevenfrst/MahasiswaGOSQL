package delivery

import (
	"database/sql"
	"enigma-mhs/models"
	"enigma-mhs/repository"
	usecase2 "enigma-mhs/usecase"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type AppDelivery struct {
	db *sql.DB
}

func NewAppDelivery(c *sql.DB) *AppDelivery {
	return &AppDelivery{
		db:c,
	}
}

func (bd AppDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"1": "Menu Mahasiswa",
		"2": "Menu Matkul",
		"3": "Menu Mahasiswa Dan Matkulnya",
		"q":  "Quit App",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Warehouse Application")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _,menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}


func (bd AppDelivery) mahasiswaForm() {
	var mahasiswaMenu = map[string]string{
		"1": "tambah mahasiswa",
		"2": "hapus mahasiswa",
		"3": "detail mahasiswa",
		"q":  "back to mainmenu",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Child Menu Mahasiswa")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _,menuCode := range bd.menuChoiceOrdered(mahasiswaMenu) {
		fmt.Printf("%s. %s\n", menuCode, mahasiswaMenu[menuCode])
	}
}

func (bd AppDelivery) matkulForm() {
	var matkulMenu = map[string]string{
		"1": "tambah matkul",
		"2": "hapus matkul",
		"3": "detail matkul",
		"q":  "back to mainmenu",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Child Menu Matkul")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _,menuCode := range bd.menuChoiceOrdered(matkulMenu) {
		fmt.Printf("%s. %s\n", menuCode, matkulMenu[menuCode])
	}
}

func (bd AppDelivery) dataMHSMATKULForm() {
	var matkulMenu = map[string]string{
		"1": "Mahasiswa Tambah Matkul",
		"2": "Mahasiswa Menghapus Matkul",
		"3": "Cari Matkul Yang Mahasiswa Ambil",
		"q":  "back to mainmenu",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Child Menu mahasiswa tambah matkul")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _,menuCode := range bd.menuChoiceOrdered(matkulMenu) {
		fmt.Printf("%s. %s\n", menuCode, matkulMenu[menuCode])
	}
}



func (bd *AppDelivery) Create() {
	var isExist = false
	var userChoice string
	var repoMhs = repository.NewMahasiswaRepository(bd.db)
	var repoMatkul = repository.NewMatkulRepository(bd.db)
	var repoMhsMatkul = repository.NewMhsMatkulRepository(bd.db)
	bd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "1":
			fmt.Println("Mahasiswa Menu")
			var isExist2 = false
			var userChoice2 string
			var name string
			var kota string
			var Id string
			var err error
			for isExist2 == false {
				bd.mahasiswaForm()
				fmt.Printf("\n%s", "Your Choice: ")
				fmt.Scan(&userChoice2)
				switch {
				case userChoice2 == "1":
					fmt.Println("tambah mahasiswa")
					fmt.Printf("\n%s", "Nama Mahasiswa: ");fmt.Scan(&name)
					if len(name) <= 3 {
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					fmt.Printf("\n%s", "Kota Asal: ");fmt.Scan(&kota)
					if len(kota) <= 3{
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					_,err = repoMhs.Insert(models.Mahasiswa{
						Nama_Mhs: name,
						Kota_Asal: kota,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "2":
					fmt.Println("hapus mahasiswa")
					fmt.Printf("\n%s", "Nama Mahasiswa: ");fmt.Scan(&name)
					if len(name) <= 3 {
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					_,err = repoMhs.Delete(models.Mahasiswa{
						Nama_Mhs:name,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "3":
					fmt.Println("detail mahasiswa By ID")
					fmt.Printf("\n%s", "ID Mahasiswa: ");fmt.Scan(&Id)
					if len(Id) != 36 {
						err = errors.New("ID Tidak 36")
						panic(err)
					}
					//_,err := repo.SelectById
					usecase := usecase2.NewMahasiswaUseCase(repoMhs)
					result,err := usecase.GetMahasiswaById(Id)
					if err != nil {
						panic(err)
					}
					fmt.Println(result)
				case userChoice2 == "q":
					isExist2 = true
					bd.mainMenuForm()
				default:
					fmt.Println("Unknown Code")
				}
			}
		case userChoice == "2":
			fmt.Println("Matkul Menu")
			var isExist2 = false
			var userChoice2 string
			var Id string
			var dosen string
			var name string
			var err error
			for isExist2 == false {
				bd.matkulForm()
				fmt.Printf("\n%s", "Your Choice: ")
				fmt.Scan(&userChoice2)
				switch {
				case userChoice2 == "1":
					fmt.Println("tambah matkul")
					fmt.Printf("\n%s", "Nama Matkul: ");fmt.Scan(&name)
					if len(name) <= 3 {
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					fmt.Printf("\n%s", "Nama Dosen: ");fmt.Scan(&dosen)
					if len(dosen) <= 3 {
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					_,err = repoMatkul.Insert(models.Matkul{
						Nama_Matkul: name,
						Nama_Dosen:  dosen,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "2":
					fmt.Println("hapus matkul")
					fmt.Printf("\n%s", "Nama Matkul: ");fmt.Scan(&name)
					if len(name) <= 3 {
						err = errors.New("Harus Lebih Dari 3")
						panic(err)
					}
					_,err = repoMatkul.DeleteByName(models.Matkul{
						Nama_Matkul: name,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "3":
					fmt.Println("detail mahasiswa By ID")
					fmt.Printf("\n%s", "ID Matkul: ");fmt.Scan(&Id)
					if len(Id) != 36 {
						err = errors.New("ID Tidak 36")
						panic(err)
					}
					//_,err := repo.SelectById
					usecase := usecase2.NewMatkulUseCase(repoMatkul)
					result,err := usecase.GetMatkulById(Id)
					if err != nil {
						panic(err)
					}
					fmt.Println(result)
				case userChoice2 == "q":
					isExist2 = true
					bd.mainMenuForm()
				default:
					fmt.Println("Unknown Code")
				}
			}
		case userChoice == "3":
			fmt.Println("Data Mahasiswa Dan Matkul")
			var isExist2 = false
			var userChoice2 string
			var err error
			var idMhs string
			var idMatkul string
			for isExist2 == false {
				bd.dataMHSMATKULForm()
				fmt.Printf("\n%s", "Your Choice: ")
				fmt.Scan(&userChoice2)
				switch {
				case userChoice2 == "1":
					fmt.Println("tambah data mahasiswa dan matkulnya")
					fmt.Printf("\n%s", "ID Mahasiswa: ");fmt.Scan(&idMhs)
					if len(idMhs) != 36 {
						err = errors.New("ID Tidak 36")
						panic(err)
					}
					fmt.Printf("\n%s", "ID Matkul: ");fmt.Scan(&idMatkul)
					if len(idMatkul) != 36 {
						err = errors.New("ID Tidak 36")
						panic(err)
					}
					_, err = repoMhsMatkul.Insert(models.MhsMatkul{
						Id_Mahasiswa: idMhs,
						Id_Matkul:    idMatkul,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "2":
					fmt.Println("hapus data Bedasarkan Id Mahasiswa")
					fmt.Printf("\n%s", "ID Mahasiswa: ");fmt.Scan(&idMhs)
					if len(idMatkul) != 36 {
						err = errors.New("ID Tidak 36")
						panic(err)
					}
					_,err = repoMhsMatkul.DeleteByMHSID(models.MhsMatkul{
						Id_Mahasiswa: idMhs,
					})
					if err != nil {
						panic(err)
					}
				case userChoice2 == "3":
					fmt.Println("detail mahasiswa dan matkulnya")
					fmt.Printf("\n%s", "ID Mahasiswa: ");fmt.Scan(&idMhs)
					usecase := usecase2.NewMhsMatkulUseCase(repoMhsMatkul)
					result,err := usecase.GetAllMhsById(idMhs)
					if err != nil {
						panic(err)
					}
					for _,y := range(result) {
						fmt.Println(y)
					}

				case userChoice2 == "q":
					isExist2 = true
					bd.mainMenuForm()
				default:
					fmt.Println("Unknown Code")
				}
			}
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")
		}

	}
}

func (bd *AppDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd AppDelivery) Run() {
	bd.Create()
}