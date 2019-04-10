package main

import (
    "github.com/strix-kr/go-standard-project/pkg/databases/rdb"
    "github.com/strix-kr/go-standard-project/pkg/services"
    "log"
)

// application 구동에 필요한 prerequisite 를 설정합니다.
// 예를 들면 database pool 생성, 환경변수 설정 등을 수행할 수 있습니다.
func init() {
    
    // 관계형 데이터베이스를 초기화 합니다.
    // 오류가 발생시 panic 을 발생시킵니다.
    if err := rdb.New(); err != nil {
        panic(err.(error))
    }
    
    // 그 외 필요한 설정을 여기에 추가합니다.
    
}

// application 의 entry point 입니다.
// 이 함수가 실행될 때는 application 구동에 필요한 모든 설정이 완료되었다고 가정합니다.
func main() {
    
    // application 이 종료될 때 필요한 과정을 기술합니다.
    defer func() {
        // application 이 종료될 때 database 도 닫아줍니다.
        if err := rdb.Session().Close(); err != nil {
            log.Fatal(err)
        }
    }()
    
    // 새로운 사용자를 등록합니다.
    newUser := services.User{
        Name: "USER_NAME",
    }
    
    insertedId, err := services.CreateUser(&newUser)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println(insertedId)
    log.Println(newUser)
    
    // 단일 사용자를 조회합니다.
    user, err := services.GetUser(1)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println(user)
    
    // 사용자 정보를 수정합니다.
    rowsAffected, err := services.UpdateUser(&newUser)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println(rowsAffected)
    log.Println(newUser)
    
    // 사용자 목록을 조회합니다.
    users, err := services.ListUsers()
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Println(users)
}
