package repository

import (
	"database/sql"
	"enigma-mhs/models"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
)

type IMahasiswaRepository interface {
	Insert(mahasiswa models.Mahasiswa) (*models.Mahasiswa,error)
	Delete(mahasiswa models.Mahasiswa) (*models.Mahasiswa,error)
	SelectById(name string) (*models.Mahasiswa, error)
}

var (
	mhsQueries = map[string]string{
		"insertMahasiswa":"insert into mahasiswa_db(id,nama,kota_asal) values(?,?,?)",
		"deleteMahasiswa":"delete from mahasiswa_db where nama=?",
		"detailMahasiswaById":"select id,nama,kota_asal from mahasiswa_db where id=?",
	}
)

type MahasiswaRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewMahasiswaRepository(db *sql.DB) IMahasiswaRepository {
	ps := make(map[string]*sql.Stmt,len(mhsQueries))
	for n,v := range mhsQueries {
		p,err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n]=p
	}
	return &MahasiswaRepository{
		db:db,ps:ps,
	}
}

func (r *MahasiswaRepository) Insert(mahasiswa models.Mahasiswa) (*models.Mahasiswa,error){
	id := guuid.New()
	mahasiswa.Id_Mahasiswa = id.String()
	res,err := r.ps["insertMahasiswa"].Exec(mahasiswa.Id_Mahasiswa,mahasiswa.Nama_Mhs,mahasiswa.Kota_Asal)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &mahasiswa,nil
}

func (r *MahasiswaRepository) Delete(mahasiswa models.Mahasiswa) (*models.Mahasiswa,error) {
	res,err := r.ps["deleteMahasiswa"].Exec(mahasiswa.Nama_Mhs)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &mahasiswa,nil
}

func (r *MahasiswaRepository) SelectById(Id string) (*models.Mahasiswa, error) {
	rows, err := r.ps["detailMahasiswaById"].Query(Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := new(models.Mahasiswa)

	for rows.Next() {
		err = rows.Scan(&res.Id_Mahasiswa, &res.Nama_Mhs, &res.Kota_Asal)
	}
	if err != nil {
		panic(err)
	}
	//log.Println(res.Nama_Mhs,res.Kota_Asal)
	return res, nil
}
