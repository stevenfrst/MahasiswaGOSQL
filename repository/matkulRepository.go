package repository

import (
	"database/sql"
	"enigma-mhs/models"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
)

type IMatkulRepository interface {
	Insert(matkul models.Matkul) (*models.Matkul,error)
	DeleteByName(matkul models.Matkul) (*models.Matkul,error)
	SelectById(Id string) (*models.Matkul, error)
}

type MatkulRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

var (
	matkulQueries = map[string]string{
		"insertMatkul":"insert into matkul_db(id_matkul,nama_matkul,nama_dosen) values(?,?,?)",
		"deleteMatkulByName":"delete from matkul_db where nama_matkul=?",
		"detailMatkulById":"select id_matkul,nama_matkul,nama_dosen from matkul_db where id_matkul=?",
	}
)

func NewMatkulRepository(db *sql.DB) IMatkulRepository {
	ps := make(map[string]*sql.Stmt,len(matkulQueries))
	for n,v := range matkulQueries {
		p,err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n]=p
	}
	return &MatkulRepository{
		db:db,ps:ps,
	}
}

func (m *MatkulRepository) Insert(matkul models.Matkul) (*models.Matkul,error){
	id := guuid.New()
	matkul.Id_Matkul = id.String()
	res,err := m.ps["insertMatkul"].Exec(matkul.Id_Matkul,matkul.Nama_Matkul,matkul.Nama_Dosen)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &matkul,nil
}

func (m *MatkulRepository) DeleteByName(matkul models.Matkul) (*models.Matkul,error) {
	res,err := m.ps["deleteMatkulByName"].Exec(matkul.Nama_Matkul)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Delete failed", err))
	}
	return &matkul,nil
}

func (m *MatkulRepository) SelectById(Id string) (*models.Matkul, error) {
	rows, err := m.ps["detailMatkulById"].Query(Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := new(models.Matkul)

	for rows.Next() {
		err = rows.Scan(&res.Id_Matkul, &res.Nama_Matkul, &res.Nama_Dosen)
	}
	if err != nil {
		panic(err)
	}
	return res, nil
}