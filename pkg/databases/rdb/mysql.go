package rdb

import (
    "database/sql"
    "errors"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/strix-kr/go-standard-project/sxerrors"
    "net/url"
    "os"
    "time"
)

// connection pool instance 입니다.
var database *sql.DB

// 새로운 관계형 데이터베이스 connection pool 을 생성합니다.
// connection pool 을 생성할 때 연결할 데이터베이스의 timezone 을 고려해야합니다.
func New() (err error) {
    
    // 아래의 값을 받아와야 합니다.
    // 이 값들을 설정하기 위한 방법을 여러개가 있습니다.
    // docker 를 활용할 경우 환경변수로 활용해도 좋습니다.
    // viper 를 이용하는 방법도 있습니다.
    // 필요에 맞게 선택하면 됩니다.
    var user, password, host, port, db string
    
    var has bool
    
    if host, has = os.LookupEnv("DB_HOST"); !has {
        return sxerrors.MissingEnvironmentVariable{Message: "DB_HOST_NOT_FOUND"}
    }
    if port, has = os.LookupEnv("DB_PORT"); !has {
        return sxerrors.MissingEnvironmentVariable{Message: "DB_PORT_NOT_FOUND"}
    }
    if user, has = os.LookupEnv("DB_USER"); !has {
        return sxerrors.MissingEnvironmentVariable{Message: "DB_USER_NOT_FOUND"}
    }
    if password, has = os.LookupEnv("DB_USER_PASSWORD"); !has {
        return sxerrors.MissingEnvironmentVariable{Message: "DB_USER_PASSWORD_NOT_FOUND"}
    }
    if db, has = os.LookupEnv("DB_NAME"); !has {
        return sxerrors.MissingEnvironmentVariable{Message: "DB_NAME_NOT_FOUND"}
    }
    
    // parseTime option 은 mysql 의 date 또는 datetime 의 값을 golang 의 time.Time type 으로 변환합니다.
    // loc option 은 golang 의 time.Time 의 timezone 을 설정합니다. 이 option 은 실제 database 의 timezone 을 변경하지 않습니다.
    // 만약 loc option 을 사용하고 배포 artifact 가 binary 일 경우 application 이 설치되는 host machine 에 golang 이 설치되어있거나 zoneinfo.zip 파일을 올바른 경로에 복사해야 합니다.
    // 자세한 내용은 다음을 참조하세요. https://golang.org/src/time/zoneinfo.go
    // 이 외 다른 설정도 가능합니다. 더 자세한 내용은 https://github.com/go-sql-driver/mysql 를 참조하세요.
    dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s", user, password, host, port, db, url.QueryEscape("Asia/Seoul"))
    
    if db, err := sql.Open("mysql", dataSource); err != nil {
        // 데이터베이스가 초기화 과정에서 문제가 발생하면 error 를 반환합니다.
        return err
        
    } else {
        
        // 데이터베이스 instance 입니다
        database = db
        
        // connection pool 에 있는 connection 의 lifetime 을 설정합니다.
        // 이 값을 설정할 때, 데이터베이스의 설정값을 참조하세요.(mysql 의 경우 wait_timeout 등)
        database.SetConnMaxLifetime(1 * time.Minute)
        // 최대로 연결할 수 있는 connection 수 입니다.
        // 이 값을 설정할 때, 데이터베이스의 설정값을 참조하세요.
        database.SetMaxOpenConns(10)
        // connection pool 에서 유지할 유후상태의 connection 수 입니다.
        // 이 값을 설정할 때, max open connection 값보다 클 수 없습니다.
        database.SetMaxIdleConns(5)
    }
    
    return
    
}

// 데이터베이스 connection 을 반환합니다.
func Session() *sql.DB {
    
    // 데이터베이스가 초기화되지 않았다면 panic 을 발생시킵니다.
    if database == nil {
        panic(errors.New("DatabaseNotInitialized"))
    }
    
    return database
}
