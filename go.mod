module gitee.com/panda8xy/gin-blog

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qiniu/api.v7/v7 v7.6.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.61.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	gitee.com/panda8xy/gin-blog/api => ./api
	gitee.com/panda8xy/gin-blog/middleware => ./middleware
	gitee.com/panda8xy/gin-blog/model => ./model
	gitee.com/panda8xy/gin-blog/routes => ./routes
	gitee.com/panda8xy/gin-blog/upload => ./upload
	gitee.com/panda8xy/gin-blog/utils => ./utils
	gitee.com/panda8xy/gin-blog/utils/errors => ./utils/errors
	gitee.com/panda8xy/gin-blog/web => ./web
)
