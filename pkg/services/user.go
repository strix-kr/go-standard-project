package services

import (
    "github.com/strix-kr/go-standard-project/pkg/databases/rdb"
    "github.com/strix-kr/go-standard-project/utils/datetimeutils"
    "log"
    "time"
)

// user model
type User struct {
    ID        uint64    `json:"id"`         // 아이디
    Name      string    `json:"name"`       // 이름
    CreatedAt time.Time `json:"created_at"` // 생성일자
    UpdatedAt time.Time `json:"updated_at"` // 수정일자
}

// user 생성
func CreateUser(user *User) (insertedId int64, err error) {
    
    query := `insert into users(name, created_at) values (?, ?)`
    
    // 현재 시각을 설정합니다.
    user.CreatedAt = datetimeutils.Now()
    
    // 새로운 user 를 생성합니다.
    // 같은 sql 의 반복작업이 많을 경우 prepared statement 를 사용해도됩니다.
    result, err := rdb.Session().Exec(query, user.Name, user.CreatedAt)
    if err != nil {
        return
    }
    
    // 영향받은 rows 수 를 반환합니다. 일반적으로 1 을 반환받습니다.
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Printf("rows inserted : %d", rowsAffected)
    
    // id 가 pk 이고 auto increment 일 경우 그 값을 반환받습니다.
    insertedId, err = result.LastInsertId()
    if err != nil {
        return
    }
    
    user.ID = uint64(insertedId)
    
    return
}

// user 수정
func UpdateUser(user *User) (rowsAffected int64, err error) {
    
    query := `update users set id = ?, name = ?, updated_at = ? where id = ?`
    
    user.UpdatedAt = datetimeutils.Now()
    
    result, err := rdb.Session().Exec(query, user.ID, user.Name, user.UpdatedAt, user.ID)
    if err != nil {
        return
    }
    
    rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Printf("rows updated : %d", rowsAffected)
    
    return
    
}

// user 삭제
func DeleteUser(id uint64) (rowsAffected int64, err error) {
    
    query := `delete from users where id = ?`
    
    result, err := rdb.Session().Exec(query, id)
    if err != nil {
        return
    }
    
    rowsAffected, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Printf("rows deleted : %d", rowsAffected)
    
    return
    
}

// user 조회
func GetUser(id uint64) (user User, err error) {
    
    query := `select * from users where id = ?`
    
    row := rdb.Session().QueryRow(query, id)
    
    // 조회한 모든 데이터를 return type 에 맞게 생성합니다.
    var userId uint64
    var name string
    var createdAt time.Time
    if err = row.Scan(&userId, &name, &createdAt); err != nil {
        log.Fatal(err)
        return
    }
    
    user = User{
        ID:        userId,
        Name:      name,
        CreatedAt: createdAt,
    }
    
    return
}

// user 목록 조회
func ListUsers() (users []User, err error) {
    
    query := `select * from users`
    
    rows, err := rdb.Session().Query(query)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    defer func() {
        // 모든 읽기 작업이 끝나면 rows connection 을 닫아줍니다.
        // 특히나 예상치 못한 오류가 발생했을 경우 connection 의 누수를 예방합니다.
        if err := rows.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    
    // 조회한 모든 데이터를 return type 에 맞게 생성합니다.
    for rows.Next() {
        
        var id uint64
        var name string
        var createdAt time.Time
        err = rows.Scan(&id, &name, &createdAt)
        
        users = append(users, User{
            ID:        id,
            Name:      name,
            CreatedAt: createdAt,
        })
    }
    
    return
}
