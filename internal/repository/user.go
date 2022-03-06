package repository

import (
	"inovasi-aktif-go/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"inovasi-aktif-go/graph/model"
	"inovasi-aktif-go/pkg/jwt"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func CreateUser(input model.NewUser) (string, error) {
	UserID, err := InputUser(input)
	if err != nil {
		return "", err
	}

	// Input user address
	statement, err := database.Db.Prepare("INSERT INTO User_Address(DesaID,KecamatanID,Street,UserID) VALUES(?,?,?,?)")
	if err != nil {
		return "", err
	}
	_, err = statement.Exec(input.Address.DesaID, input.Address.KecamatanID, input.Address.Street, UserID)
	if err != nil {
		return "", err
	}

	// Generate token
	token, err := jwt.GenerateToken(input.Phone)
	if err != nil {
		return "", err
	}

	return token, nil
}

func InputUser(input model.NewUser) (int64, error) {
	// Check input value
	if input.Name == "" {
		return 0, gqlerror.Errorf("Nama tidak boleh kosong.")
	}
	if input.Phone == "" {
		return 0, gqlerror.Errorf("Nomor HP tidak boleh kosong.")
	}
	if input.Password == "" {
		return 0, gqlerror.Errorf("Kata Sandi tidak boleh kosong.")
	}

	// Check user exist
	id, err := GetUserIdByPhone(input.Phone)
	if id != 0 {
		return 0, gqlerror.Errorf("Pengguna dengan nomor HP tersebut sudah ada.")
	}

	// Input user
	stmt, err := database.Db.Prepare("INSERT INTO User(Name,Phone,Password) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}

	// Encrypt password
	hashedPassword, err := HashPassword(input.Password)
	res, err := stmt.Exec(input.Name, input.Phone, hashedPassword)
	if err != nil {
		return 0, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return userID, nil
}

func AuthenticateUser(input model.LoginUser) (string, error) {
	// Check input value
	if input.Phone == "" {
		return "", gqlerror.Errorf("Nomor HP tidak boleh kosong.")
	}
	if input.Password == "" {
		return "", gqlerror.Errorf("Kata Sandi tidak boleh kosong.")
	}

	// Check user
	statement, err := database.Db.Prepare("SELECT Password FROM User WHERE Phone = ?")
	if err != nil {
		return "", gqlerror.Errorf("Akun dengan Nomor HP tersebut tidak ditemukan.")
	}
	row := statement.QueryRow(input.Phone)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", gqlerror.Errorf("Akun dengan Nomor HP tersebut tidak ditemukan.")
		} else {
			return "", err
		}
	}

	// Check password
	correct := CheckPasswordHash(input.Password, hashedPassword)
	if !correct {
		return "", gqlerror.Errorf("Kata Sandi yang dimasukan salah.")
	}

	// Generate token
	token, err := jwt.GenerateToken(input.Phone)
	if err != nil {
		return "", err
	}

	return token, nil
}

func UserByID(id string) (*model.User, error) {
	stmt, err := database.Db.Prepare("SELECT User.ID, User.Name, User.Phone, User.Photo, User_Address.ID, User_Address.Street, Address_Desa.ID, Address_Desa.Name, Address_Kecamatan.ID, Address_Kecamatan.Name FROM User INNER JOIN User_Address ON User_Address.UserID = User.ID INNER JOIN Address_Desa ON Address_Desa.ID = User_Address.DesaID INNER JOIN Address_Kecamatan ON Address_Kecamatan.ID = User_Address.KecamatanID WHERE User.ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ID, Name, Phone, AddressID, AddressStreet, DesaID, DesaName, KecamatanID, KecamatanName string
	var Photo sql.NullString
	
	for rows.Next() {
		err := rows.Scan(&ID, &Name, &Phone, &Photo, &AddressID, &AddressStreet, &DesaID, &DesaName, &KecamatanID, &KecamatanName)
		if err != nil {
			return nil, err
		}
	}

	if ID == "" {
		return nil, gqlerror.Errorf("Pengguna tidak ditemukan.")
	}

	user := &model.User{
		ID:    ID,
		Name:  Name,
		Phone: Phone,
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

	return user, nil
}

func UserList() ([]*model.User, error) {
	stmt, err := database.Db.Prepare("SELECT User.ID, User.Name, User.Phone, User.Photo, User_Address.ID, User_Address.Street, Address_Desa.ID, Address_Desa.Name, Address_Kecamatan.ID, Address_Kecamatan.Name FROM User INNER JOIN User_Address ON User_Address.UserID = User.ID INNER JOIN Address_Desa ON Address_Desa.ID = User_Address.DesaID INNER JOIN Address_Kecamatan ON Address_Kecamatan.ID = User_Address.KecamatanID")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		var ID, Name, Phone, AddressID, AddressStreet, DesaID, DesaName, KecamatanID, KecamatanName string
		var Photo sql.NullString
		err := rows.Scan(&ID, &Name, &Phone, &Photo, &AddressID, &AddressStreet, &DesaID, &DesaName, &KecamatanID, &KecamatanName)
		if err != nil {
			return nil, err
		}

		user := &model.User{
			ID:    ID,
			Name:  Name,
			Phone: Phone,
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
		
		users = append(users, user)
	}

	return users, nil
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