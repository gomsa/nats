package migrations

import (
	pd "github.com/gomsa/nats/proto/template"
	db "github.com/gomsa/nats/providers/database"
)

func init() {
	template()
	seeds()
}

// user 用户数据迁移
func template() {
	template := &pd.Template{}
	if !db.DB.HasTable(&template) {
		db.DB.Exec(`
			CREATE TABLE templates (
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
		INSERT INTO templates ( event, name, type, templateCode, templateValue ) VALUES ('register_verify','用户注册验证码','sms','SMS_135275049','')
	`)
}
