package entities

import "github.com/joaodrts/Drts.Ads.GO.API/data"

func Insert(ad Ad) (id int64, err error) {
	conn, err := data.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO ads (title, description, datestart, dateend, status, link) VALUES ($1, $2, $3, $4, $5, $6)`

	err = conn.QueryRow(sql, ad.Title, ad.Description, ad.DateStart, ad.DateEnd, ad.Status, ad.Link).Scan(&id)

	return
}

func GetById(id string) (ad Ad, err error) {
	conn, err := data.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM ads WHERE datestart >= current_date and dateend >= current_date and status=0 and id=$1`, id)

	err = row.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.DateStart, &ad.DateEnd, &ad.Status, &ad.Link)

	return
}

func Get(status string) (ads []Ad, err error) {
	conn, err := data.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM ads WHERE datestart >= current_date and dateend >= current_date and status=$1`, status)

	if err != nil {
		return
	}

	for rows.Next() {
		var ad Ad

		err = rows.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.DateStart, &ad.DateEnd, &ad.Status, &ad.Link)

		if err != nil {
			continue
		}

		ads = append(ads, ad)
	}

	return
}

func Update(id string, ad Ad) (idresp int64, err error) {

	conn, err := data.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql, err := conn.Exec(`UPDATE ads SET title=$2, description=$3, datestart=$4, dateend=$5, status=$6, link=$7 WHERE id =$1`,
		id, ad.Title, ad.Description, ad.DateStart, ad.DateEnd, ad.Status, ad.Link)

	if err != nil {
		return 0, err
	}

	return sql.RowsAffected()
}

func Delete(id string) (int64, error) {

	conn, err := data.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	// Status 0 -> ativo, 1 -> inativo, 2 -> finalizado
	sql, err := conn.Exec(`UPDATE ads SET status=1 WHERE id =$1`, id)

	if err != nil {
		return 0, err
	}

	return sql.RowsAffected()
}
