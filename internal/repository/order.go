package repository

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"inovasi-aktif-go/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func OrderByID(id string) (*model.Order, error) {
	stmt, err := database.Db.Prepare("SELECT Sales_Order.ID, User.ID, User.Name FROM Sales_Order INNER JOIN User ON User.ID = Sales_Order.UserID WHERE Sales_Order.ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ID, UserID, UserName string
	
	for rows.Next() {
		err := rows.Scan(&ID, &UserID, &UserName)
		if err != nil {
			return nil, err
		}
	}

	if ID == "" {
		return nil, gqlerror.Errorf("Order tidak ditemukan.")
	}

	order := &model.Order {
		ID: ID,
		User: &model.User {
			ID: UserID,
			Name: UserName,
		},
	}

	return order, nil
}