package email_processor

import "github.com/magiconair/properties"

func GetStringByKey(keyName string) string {
	p := properties.MustLoadFile("/home/miguel/go-hello/app.properties", properties.UTF8)
	return p.MustGetString(keyName)
}
