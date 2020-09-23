package repository

import (
	"database/sql"
	"enigma-mhs/models"
	"errors"
	"fmt"

)

type IMhsMatkulRepository interface {
	Insert(q models.MhsMatkul) (*models.MhsMatkul,error)
	DeleteByMHSID(q models.MhsMatkul) (*models.MhsMatkul,error)
	SelectByMHSId(MhsId string) ([]*models.MhsMatkul, error)
}

var (
	mhsmatkulQueries = map[string]string{
		"insertMhsMatkul":"insert into matkulmhs_db(id_mahasiswa,id_matkul) values(?,?)",
		"deleteMhsMatkulByMHSID":"delete from matkulmhs_db where id_mahasiswa=?",
		"detailMatkulMahasiswa":"select qty,id_mahasiswa,id_matkul from matkulmhs_db where id_mahasiswa=?",
	}
)

type MhsMatkulRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewMhsMatkulRepository(db *sql.DB) IMhsMatkulRepository {
	ps := make(map[string]*sql.Stmt,len(mhsmatkulQueries))
	for n,v := range mhsmatkulQueries {
		p,err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n]=p
	}
	return &MhsMatkulRepository{
		db:db,ps:ps,
	}
}

func (r *MhsMatkulRepository) Insert(q models.MhsMatkul) (*models.MhsMatkul,error){
	res,err := r.ps["insertMhsMatkul"].Exec(q.Id_Mahasiswa,q.Id_Matkul)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &q,nil
}

func (r *MhsMatkulRepository) DeleteByMHSID(q models.MhsMatkul) (*models.MhsMatkul,error) {
	res,err := r.ps["deleteMhsMatkulByMHSID"].Exec(q.Id_Mahasiswa)
	if err != nil {
		return nil,err
	}
	affectedNo,err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &q,nil
}

func (r *MhsMatkulRepository) SelectByMHSId(MhsId string) ([]*models.MhsMatkul, error) {
	rows, err := r.ps["detailMatkulMahasiswa"].Query(MhsId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []*models.MhsMatkul

	for rows.Next() {
		res := new(models.MhsMatkul)
		err = rows.Scan(&res.Qty,&res.Id_Mahasiswa, &res.Id_Matkul)
		if err != nil {
			panic(err)
		}
		result = append(result,res)
	}
	//log.Println(res,reflect.TypeOf(res))
	return result, nil
}

