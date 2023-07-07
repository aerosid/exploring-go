package city

import "example/codepen/db"

type City struct {
	CityId int64  `db:"city_id"`
	Name   string `db:"name"`
	State  string `db:"state"`
}

func FindById(id int64) (c *City) {
	c = &City{}
	sql := "select * from city where city_id = ?"
	db, err := db.GetInstance()
	if err == nil {
		err := db.Get(c, sql, id)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return c
}

func FindByState(state string) (list *[]City) {
	list = &[]City{}
	sql := "select * from city where state = ?"
	db, err := db.GetInstance()
	if err == nil {
		err := db.Select(list, sql, state)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return list
}

func (c *City) Save() *City {
	if c.CityId == 0 {
		c.CityId = c.insert()
	} else {
		c.update()
	}
	return c
}

func (c *City) Delete() (count int64) {
	sql := "delete from city where city_id = :city_id"
	db, err := db.GetInstance()
	if err == nil {
		result, err := db.NamedExec(sql, c)
		if err == nil {
			rowsAffected, err := result.RowsAffected()
			if err == nil {
				count = rowsAffected
			} else {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return count
}

func (c *City) insert() (id int64) {
	sql := "insert into city (name, state) values (:name, :state)"
	db, err := db.GetInstance()
	if err == nil {
		result, err := db.NamedExec(sql, c)
		if err == nil {
			last, err := result.LastInsertId()
			if err == nil {
				id = last
			} else {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return id
}

func (c *City) update() (count int64) {
	sql := "update city set name = :name, state = :state where city_id = :city_id"
	db, err := db.GetInstance()
	if err == nil {
		result, err := db.NamedExec(sql, c)
		if err == nil {
			rowsAffected, err := result.RowsAffected()
			if err == nil {
				count = rowsAffected
			} else {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return count
}
