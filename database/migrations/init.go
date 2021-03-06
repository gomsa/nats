package migrations

import (
	pd "github.com/gomsa/nats/proto/template"
	db "github.com/gomsa/nats/providers/database"
)

func init() {
	template()
	seeds()
}

// template 模板数据迁移
func template() {
	template := &pd.Template{}
	if !db.DB.HasTable(&template) {
		db.DB.Exec(`
			CREATE TABLE templates (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,	
			event varchar(32) NOT NULL,
			name varchar(64) DEFAULT NULL,
			type varchar(64) DEFAULT NULL,
			template_code varchar(128) DEFAULT NULL,
			template_value varchar(128) DEFAULT NULL,
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
		INSERT INTO templates ( event, name, type, template_code, template_value ) VALUES ('register_verify','用户注册验证码','sms','453946','')
	`)
}
