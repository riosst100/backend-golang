package repository

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"database/sql"
	"inovasi-aktif-go/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func BusinessList() ([]*model.Business, error) {
	stmt, err := database.Db.Prepare("SELECT Business.ID, Business.Name, Business.Photo, Business_Address.ID, Business_Address.Street, Address_Desa.ID, Address_Desa.Name, Address_Kecamatan.ID, Address_Kecamatan.Name FROM Business INNER JOIN Business_Address ON Business_Address.BusinessID = Business.ID INNER JOIN Address_Desa ON Address_Desa.ID = Business_Address.DesaID INNER JOIN Address_Kecamatan ON Address_Kecamatan.ID = Business_Address.KecamatanID")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var businesses []*model.Business
	for rows.Next() {
		var ID, Name, AddressID, AddressStreet, DesaID, DesaName, KecamatanID, KecamatanName string
		var Photo sql.NullString
		err := rows.Scan(&ID, &Name, &Photo, &AddressID, &AddressStreet, &DesaID, &DesaName, &KecamatanID, &KecamatanName)
		if err != nil {
			return nil, err
		}

		business := &model.Business{
			ID:    ID,
			Name:  Name,
			Photo: Photo.String,
			Address: &model.Address {
				ID: AddressID,
				Street: AddressStreet,
				Desa: &model.Desa {
					ID: DesaID,
					Name: DesaName,
				},
				Kecamatan: &model.Kecamatan {
					ID: KecamatanID,
					Name: KecamatanName,
				},
			},
		}
		
		businesses = append(businesses, business)
	}

	return businesses, nil
}

func BusinessByID(id string) (*model.Business, error) {
	stmt, err := database.Db.Prepare("SELECT Business.ID, Business.Name, Business.Photo, Business_Address.ID, Business_Address.Street, Address_Desa.ID, Address_Desa.Name, Address_Kecamatan.ID, Address_Kecamatan.Name, User.ID, User.Name, User.Photo FROM Business INNER JOIN Business_Address ON Business_Address.BusinessID = Business.ID INNER JOIN Address_Desa ON Address_Desa.ID = Business_Address.DesaID INNER JOIN Address_Kecamatan ON Address_Kecamatan.ID = Business_Address.KecamatanID INNER JOIN User ON User.ID = Business.UserID WHERE Business.ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ID, Name, AddressID, AddressStreet, DesaID, DesaName, KecamatanID, KecamatanName, UserID, UserName string
	var Photo, UserPhoto sql.NullString
	
	for rows.Next() {
		err := rows.Scan(&ID, &Name, &Photo, &AddressID, &AddressStreet, &DesaID, &DesaName, &KecamatanID, &KecamatanName, &UserID, &UserName, &UserPhoto)
		if err != nil {
			return nil, err
		}
	}

	if ID == "" {
		return nil, gqlerror.Errorf("Bisnis tidak ditemukan.")
	}

	business := &model.Business{
		ID:    ID,
		Name:  Name,
		Photo: Photo.String,
		Address: &model.Address {
			ID: AddressID,
			Street: AddressStreet,
			Desa: &model.Desa {
				ID: DesaID,
				Name: DesaName,
			},
			Kecamatan: &model.Kecamatan {
				ID: KecamatanID,
				Name: KecamatanName,
			},
		},
		User: &model.User {
			ID: UserID,
			Name: UserName,
			Photo: UserPhoto.String,
		},
	}

	return business, nil
}

func BusinessByUserID(userID string) ([]*model.Business, error) {
	stmt, err := database.Db.Prepare("SELECT Business.ID, Business.Name, Business.Photo, Business_Address.ID, Business_Address.Street, Address_Desa.ID, Address_Desa.Name, Address_Kecamatan.ID, Address_Kecamatan.Name FROM Business INNER JOIN Business_Address ON Business_Address.BusinessID = Business.ID INNER JOIN Address_Desa ON Address_Desa.ID = Business_Address.DesaID INNER JOIN Address_Kecamatan ON Address_Kecamatan.ID = Business_Address.KecamatanID WHERE Business.UserID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var businesses []*model.Business
	for rows.Next() {
		var ID, Name, AddressID, AddressStreet, DesaID, DesaName, KecamatanID, KecamatanName string
		var Photo sql.NullString
		err := rows.Scan(&ID, &Name, &Photo, &AddressID, &AddressStreet, &DesaID, &DesaName, &KecamatanID, &KecamatanName)
		if err != nil {
			return nil, err
		}

		business := &model.Business{
			ID:    ID,
			Name:  Name,
			Photo: Photo.String,
			Address: &model.Address {
				ID: AddressID,
				Street: AddressStreet,
				Desa: &model.Desa {
					ID: DesaID,
					Name: DesaName,
				},
				Kecamatan: &model.Kecamatan {
					ID: KecamatanID,
					Name: KecamatanName,
				},
			},
		}
		
		businesses = append(businesses, business)
	}

	return businesses, nil
}