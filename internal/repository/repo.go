package repository

import (
	"database/sql"
	"entity"

	"github.com/go-sql-driver/mysql"
)

type myRepo struct {
	db *mysql.DB
}

func NewRepo() *myRepo {
	conf := entity.NewConfig()
	db, err := sql.Open("mysql", conf.UserDB+":"+conf.PasswordDB+"@/"+conf.NameDB)
	if err != nil {
		panic(err)
	}
	repo := myRepo{db: db}
	return &repo
}

func (repo *myRepo) Post(adv *entity.Adv) error {
	_, err := repo.db.Exec("INSERT INTO adv (name, description) VALUES (?, ?)", adv.Name, adv.Description)
	if err != nil {
		return err
	}
	return nil
}

func (repo *myRepo) Get(id int) (*entity.Adv, error) {
	row := repo.db.Query("SELECT name, description FROM ads WHERE id = ?", id)
	var Adv entity.Adv
	err := row.Scan(&Adv.Name, &Adv.Description)
	if err != nil {
		return nil, err
	}
	return Adv, nil
}

func (repo *myRepo) GetPage(orderBy string) (*[]entity.Adv, error) {
	var rows sql.Rows
	var advList []entity.Adv
	switch orderBy {
	case "asc":
		rows = repo.db.Query("SELECT * FROM adv ORDER BY name ASC")
	case "desc":
		rows = repo.db.Query("SELECT * FROM adv ORDER BY name DESC")
	}
	defer rows.Close()
	for i := 0; i < 10 && rows.Next(); i++ {
		var Adv entity.Adv
		err := rows.Scan(&Adv.Name, &Adv.description)
		if err != nil {
			return nil, err
		}
		advList = append(advList, Adv)
	}
	return &advList, nil
}
