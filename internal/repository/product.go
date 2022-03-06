package repository

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"database/sql"
	"inovasi-aktif-go/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ProductList() ([]*model.Product, error) {
	stmt, err := database.Db.Prepare("SELECT Product.ID, Product.Name, Product.Price, Product.Photo, Business.ID, Business.Name, Business.Photo FROM Product INNER JOIN Business ON Business.ID = Product.BusinessID")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*model.Product
	for rows.Next() {
		var ID, Price, Name, BusinessID, BusinessName string
		var Photo, BusinessPhoto sql.NullString
		err := rows.Scan(&ID, &Name, &Price, &Photo, &BusinessID, &BusinessName, &BusinessPhoto)
		if err != nil {
			return nil, err
		}

		product := &model.Product{
			ID:    ID,
			Name:  Name,
			Price: Price,
			Photo: Photo.String,
			Business: &model.Business {
				ID: BusinessID,
				Name: BusinessName,
				Photo: BusinessPhoto.String,
			},
		}
		
		products = append(products, product)
	}

	return products, nil
}

func ProductByBusinessID(businessID string) ([]*model.Product, error) {
	stmt, err := database.Db.Prepare("SELECT Product.ID, Product.Name, Product.Price, Product.Photo, Business.ID, Business.Name, Business.Photo FROM Product INNER JOIN Business ON Business.ID = Product.BusinessID WHERE Product.BusinessID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(businessID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*model.Product
	for rows.Next() {
		var ID, Price, Name, BusinessID, BusinessName string
		var Photo, BusinessPhoto sql.NullString
		err := rows.Scan(&ID, &Name, &Price, &Photo, &BusinessID, &BusinessName, &BusinessPhoto)
		if err != nil {
			return nil, err
		}

		product := &model.Product{
			ID:    ID,
			Name:  Name,
			Price: Price,
			Photo: Photo.String,
			Business: &model.Business {
				ID: BusinessID,
				Name: BusinessName,
				Photo: BusinessPhoto.String,
			},
		}
		
		products = append(products, product)
	}

	return products, nil
}

func ProductByID(id string) (*model.Product, error) {
	stmt, err := database.Db.Prepare("SELECT Product.ID, Product.Name, Product.Price, Product.Photo, Business.ID, Business.Name, Business.Photo FROM Product INNER JOIN Business ON Business.ID = Product.BusinessID WHERE Product.ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ID, Price, Name, BusinessID, BusinessName string
	var Photo, BusinessPhoto sql.NullString
	
	for rows.Next() {
		err := rows.Scan(&ID, &Name, &Price, &Photo, &BusinessID, &BusinessName, &BusinessPhoto)
		if err != nil {
			return nil, err
		}
	}

	if ID == "" {
		return nil, gqlerror.Errorf("Produk tidak ditemukan.")
	}

	product := &model.Product{
		ID:    ID,
		Name:  Name,
		Price: Price,
		Photo: Photo.String,
		Business: &model.Business {
			ID: BusinessID,
			Name: BusinessName,
			Photo: BusinessPhoto.String,
		},
	}

	return product, nil
}