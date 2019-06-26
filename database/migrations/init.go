package migrations

import (
	pd "github.com/gomsa/nats/proto/nats"
	db "github.com/gomsa/nats/providers/database"
)

func init() {
	nat()
	seeds()
}

// user 用户数据迁移
func nat() {
	nat := &pd.Nat{}
	if !db.DB.HasTable(&nat) {
		db.DB.Exec(`
			CREATE TABLE nats (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,	
			event varchar(32) NOT NULL,
			name varchar(64) DEFAULT NULL,
			type varchar(64) DEFAULT NULL,
			templateCode varchar(128) DEFAULT NULL,
			templateValue varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY event (event)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// seeds 填充文件
func seeds() {
	db.DB.Exec(`
		INSERT INTO nats ( event, name, type, templateCode, templateValue ) VALUES ('register_verify','用户注册验证码','sms','SMS_135275049','')
	`)
}
