module gitee.com/panda8xy/gin-blog

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qiniu/api.v7/v7 v7.7.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/ugorji/go v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	gitee.com/panda8xy/gin-blog/api => ./api
	gitee.com/panda8xy/gin-blog/middleware => ./middleware
	gitee.com/panda8xy/gin-blog/model => ./model
	gitee.com/panda8xy/gin-blog/routes => ./routes
	gitee.com/panda8xy/gin-blog/upload => ./upload
	gitee.com/panda8xy/gin-blog/utils => ./utils
	gitee.com/panda8xy/gin-blog/utils/errors => ./utils/errors
	gitee.com/panda8xy/gin-blog/utils/validate => ./utils/validate
	gitee.com/panda8xy/gin-blog/web => ./web
)
