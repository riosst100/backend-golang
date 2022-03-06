CREATE TABLE IF NOT EXISTS Business_Address(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    DesaID INT,
    KecamatanID INT,
    Street VARCHAR (255),
    BusinessID INT,
    FOREIGN KEY (DesaID) REFERENCES Address_Desa(ID),
    FOREIGN KEY (KecamatanID) REFERENCES Address_Kecamatan(ID),
    FOREIGN KEY (BusinessID) REFERENCES Business(ID),
    PRIMARY KEY (ID)
)