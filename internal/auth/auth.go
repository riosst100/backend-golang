package auth

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"inovasi-aktif-go/graph/model"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO User(Name,Phone,Password) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Name, user.Phone, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func GetList() []model.User {
	stmt, err := database.Db.Prepare("SELECT User.ID, User.Name, User.Phone FROM User")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func (user *User) UserAuthenticate() bool {
	statement, err := database.Db.Prepare("select Password from User WHERE Phone = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Phone)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//GetUserIdByPhone check if a user exists in database by given phone
func GetUserIdByPhone(phone string) (int, error) {
	statement, err := database.Db.Prepare("SELECT ID FROM User WHERE Phone = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(phone)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

//GetUserByID check if a user exists in database and return the user object.
func GetPhoneById(userId string) (User, error) {
	statement, err := database.Db.Prepare("select Phone from User WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(userId)

	var phone string
	err = row.Scan(&phone)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return User{}, err
	}

	return User{ID: userId, Phone: phone}, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}