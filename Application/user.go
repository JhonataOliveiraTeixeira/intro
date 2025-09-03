package application

import (
	"errors"
	"intro/db"
	"intro/domain"
	"log/slog"
)

func Create(name, email, password string, db db.DB) error {
	
	for _, v:= range db{
		if v.Email == email{
			return errors.New("email already exists")
		}
	}
	
	index := len(db)
	user := domain.User{
	 Name: name,
	 Email: email,
	 Password: password,
 }
 slog.Info("Create user", "user" ,user, "index", index )
	db[index] = user
	return nil
}

func Update(id int, name, email string, db db.DB)error{
	_, ok := db[id] 
	if !ok{
		return errors.New("user not found")
	}

	db[id] = domain.User{
		Name: name,
		Email: email,
	}
	return nil
}

func FindUnique(id int, db db.DB)(domain.User, error){
	slog.Info("FindUnique", "id", id)
	u, ok := db[id]
	if !ok{
		return domain.User{}, errors.New("user not found")
	}
	return u, nil
}

func FindAll(page int, db db.DB)[]domain.User{
	var result =  make([]domain.User, page)

	for i, v := range db{
		if i == page-1{
			break
		}
		result = append(result, v)
	}
	return result
}

func Delete(id int, db db.DB)error{
	_,ok:= db[id]
	if !ok{
		return errors.New("id not found")
	
	}
	delete(db, id)
	return  nil
}