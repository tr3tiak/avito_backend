package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tr3tiak/avito_backend/internal/entity"
)

type myRepo struct {
	db *sql.DB
}

func NewRepo() *myRepo {
	conf := entity.NewConfig()
	fmt.Println(conf.NameDB, conf.PasswordDB, conf.UserDB)
	db, err := sql.Open("mysql", conf.UserDB+":"+conf.PasswordDB+"@/"+conf.NameDB)
	if err != nil {
		panic(err)
	}
	repo := myRepo{db: db}
	return &repo
}

func (repo *myRepo) Post(adv *entity.Adv) error {
	fmt.Println("post started")
	_, err := repo.db.Exec("INSERT ads(Name, Description) VALUES (?, ?)", adv.Name, adv.Description)
	if err != nil {
		return err
	}
	return nil
}

func (repo *myRepo) Get(id int, orderBy string) (*entity.Adv, error) {
	var row *sql.Rows
	var err error
	switch orderBy {
	case "asc":
		row, err = repo.db.Query("SELECT id, Name, Description FROM ads WHERE id = ?", id)
	case "desc":
		row, err = repo.db.Query("SELECT id, Name, Description FROM ads WHERE id = ?", id)
	}

	if err != nil {
		return nil, err
	}
	var Adv entity.Adv
	err = row.Scan(&Adv.Name, &Adv.Description)
	if err != nil {
		return nil, err
	}
	return &Adv, nil
}

func (repo *myRepo) GetPage(orderBy string) (*[]entity.Adv, error) {
	var rows *sql.Rows
	var err error
	var advList []entity.Adv
	switch orderBy {
	case "asc":
		rows, err = repo.db.Query("SELECT * FROM ads ORDER BY name ASC")
	case "desc":
		rows, err = repo.db.Query("SELECT * FROM ads ORDER BY name DESC")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for i := 0; i < 10 && rows.Next(); i++ {
		var Adv entity.Adv
		err := rows.Scan(&Adv.Name, &Adv.Description)
		if err != nil {
			return nil, err
		}
		advList = append(advList, Adv)
	}
	return &advList, nil
}
