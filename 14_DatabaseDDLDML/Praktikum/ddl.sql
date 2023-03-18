--1
ALTER TABLE `alta_online_shop`.`user` 
--2
ALTER TABLE `alta_online_shop`.`user` 
ADD COLUMN `Transaction detail` INT NULL AFTER `Transaction`,
CHANGE COLUMN `product` `Product` VARCHAR(255) NULL DEFAULT NULL ,
CHANGE COLUMN `product_type` `Product_Type` VARCHAR(255) NULL DEFAULT NULL ,
CHANGE COLUMN `operator` `Operators` VARCHAR(255) NULL DEFAULT NULL ,
CHANGE COLUMN `product description` `Product_Description` VARCHAR(255) NULL DEFAULT NULL ,
CHANGE COLUMN `Payment_method` `Payment_Method` INT NULL DEFAULT NULL ,
CHANGE COLUMN `transaction` `Transaction` INT NULL DEFAULT NULL ;
--3
CREATE TABLE kurir (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
--4
ALTER TABLE kurir ADD ongkos_dasar DECIMAL(10,2) NOT NULL DEFAULT 0;
--5
ALTER TABLE kurir RENAME TO shipping;
--6
## Hapus/Drop TABLE
--7
--A.1-to-1: payment method description.
CREATE TABLE payment_method_description (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    payment_method_id INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id) ON DELETE CASCADE
);

-- Contoh insert data ke tabel payment_method_description
INSERT INTO payment_method_description (payment_method_id, description)
VALUES (1, 'Metode pembayaran ini dapat digunakan untuk transaksi dengan nominal di bawah 10 juta.');

--B.1-to-many: user dengan alamat
CREATE TABLE alamat (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    nama_alamat VARCHAR(255) NOT NULL,
    jalan VARCHAR(255) NOT NULL,
    kota VARCHAR(255) NOT NULL,
    provinsi VARCHAR(255) NOT NULL,
    kode_pos VARCHAR(10) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

-- Contoh insert data ke tabel alamat
INSERT INTO alamat (user_id, nama_alamat, jalan, kota, provinsi, kode_pos)
VALUES (1, 'Rumah', 'Jalan Merdeka No. 10', 'Jakarta Selatan', 'DKI Jakarta', '12345');

-- C.many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_detail (
    user_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, payment_method_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id) ON DELETE CASCADE
);

-- Contoh insert data ke tabel user_payment_method_detail
INSERT INTO user_payment_method_detail (user_id, payment_method_id)
VALUES (1, 1);
INSERT INTO user_payment_method_detail (user_id, payment_method_id)
VALUES (1, 2);