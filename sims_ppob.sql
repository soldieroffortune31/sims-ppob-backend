/* start user_m */
CREATE TABLE user_m(
	user_id INTEGER PRIMARY KEY AUTO_INCREMENT,
	email VARCHAR(100) NOT NULL,
	nama_depan VARCHAR(100) NOT NULL,
	nama_belakang VARCHAR(100) NOT NULL,
	photo text NULL,
	password VARCHAR(255) NOT NULL,
	token TEXT NULL
)ENGINE=INNODB;

ALTER TABLE user_m ADD COLUMN created_at DATETIME;
ALTER TABLE user_m ADD COLUMN updated_at DATETIME;
ALTER TABLE user_m ADD COLUMN deleted_at DATETIME;

/* end user_m*/

/* start userbalance_m*/
CREATE TABLE userbalance_m (
	userbalance_id INTEGER PRIMARY KEY AUTO_INCREMENT,
	user_id INTEGER UNIQUE,
	balance BIGINT,
	created_at DATETIME,
	updated_at DATETIME
)ENGINE=INNODB

ALTER TABLE userbalance_m ADD COLUMN deleted_at DATETIME;

/* end userbalance_m */

/* start jenis transaksi */

CREATE TABLE jenistransaksi_m (
	jenistransaksi_id INTEGER PRIMARY KEY AUTO_INCREMENT,
	jenis_transaksi VARCHAR(100)
)ENGINE=INNODB

/* end jenis transaksi*/


/* start transaksi */
CREATE TABLE transaksi_t(
	transaksi_id INTEGER PRIMARY KEY AUTO_INCREMENT,
	userbalance_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	saldo_terakhir BIGINT NOT NULL,
	saldo_masuk BIGINT DEFAULT 0,
	saldo_keluar BIGINT DEFAULT 0,
	saldo_sekarang BIGINT NOT NULL,
	jenistransaksi_id INTEGER,
	tgl_transaksi DATETIME,
	created_at DATETIME,
	update_at DATETIME,
	delete_at DATETIME
)ENGINE=INNODB;

ALTER TABLE transaksi_t
ADD CONSTRAINT fk_transaksi_userbalance
FOREIGN KEY (userbalance_id) REFERENCES userbalance_m(userbalance_id)
ON DELETE RESTRICT
ON UPDATE RESTRICT

ALTER TABLE transaksi_t
ADD CONSTRAINT fk_transaksi_user
FOREIGN KEY (user_id) REFERENCES user_m(user_id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;

ALTER TABLE transaksi_t
ADD CONSTRAINT fk_transaksi_jenistransaksi
FOREIGN KEY (jenistransaksi_id) REFERENCES jenistransaksi_m(jenistransaksi_id)
ON DELETE RESTRICT
ON UPDATE RESTRICT;

/* end transaksi */