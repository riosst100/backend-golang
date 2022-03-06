package repository

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"inovasi-aktif-go/graph/model"
)

func KecamatanList() ([]*model.Kecamatan, error) {
	stmt, err := database.Db.Prepare("SELECT Address_kecamatan.ID, Address_kecamatan.Name FROM Address_kecamatan")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var kecamatans []*model.Kecamatan
	for rows.Next() {
		var ID, Name string
		err := rows.Scan(&ID, &Name)
		if err != nil {
			return nil, err
		}

		kecamatan := &model.Kecamatan{
			ID:    ID,
			Name:  Name,
		}
		
		kecamatans = append(kecamatans, kecamatan)
	}

	return kecamatans, nil
}

func DesaByKecamatanID(KecamatanID string) ([]*model.Desa, error) {
	stmt, err := database.Db.Prepare("SELECT Address_Desa.ID, Address_Desa.Name FROM Address_Desa WHERE Address_Desa.KecamatanID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(KecamatanID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var desas []*model.Desa
	for rows.Next() {
		var ID, Name string
		err := rows.Scan(&ID, &Name)
		if err != nil {
			return nil, err
		}

		desa := &model.Desa{
			ID:    ID,
			Name:  Name,
		}
		
		desas = append(desas, desa)
	}

	return desas, nil
}