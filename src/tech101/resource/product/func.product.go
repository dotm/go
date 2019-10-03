package product

import (
	"fmt"
	"tech101/db"
	"tech101/model"
)

func Get(id int64) (product model.Product, err error) {
	query := fmt.Sprintf(`
		SELECT
			id, product_name, product_description, create_time, update_time
		FROM ws_product
		WHERE id = %d
	`, id)

	err = db.Client.QueryRow(query).Scan(&product.ID, &product.Name, &product.Description, &product.CreateTime, &product.UpdateTime)
	if err != nil {
		return
	}

	return
}

func GetAll(limit, offset int64) (products []model.Product, err error) {
	stmt, err := db.Client.Prepare(`
		SELECT
			id, product_name, product_description, create_time, update_time
		FROM ws_product
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
	`)
	if err != nil {
		return
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := model.Product{}
		errScan := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreateTime, &p.UpdateTime)
		if errScan != nil {
			continue
		}
		products = append(products, p)
	}
	return
}

func Create(name, description string) (product model.Product, err error) {
	query := fmt.Sprintf(`
		INSERT INTO ws_product (product_name, product_description, create_time, update_time)
		VALUES (
			'%s', '%s', NOW(), NOW()
		)
		RETURNING id, product_name, product_description, create_time, update_time
	`, name, description)

	err = db.Client.QueryRow(query).Scan(&product.ID, &product.Name, &product.Description, &product.CreateTime, &product.UpdateTime)
	if err != nil {
		return
	}

	return
}
