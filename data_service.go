package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type DataService interface {
    Init() error
    GetReportItems() []ReportItem
    GetReportItem(id int) ReportItem
    QueryRow(sql string, dest ...interface{}) error
    Close() error
}

type DataServiceMysql struct {
    Url string
    Db  *sql.DB
}

func (dsMysql *DataServiceMysql) Init() error {
    db, err := sql.Open("mysql", dsMysql.Url)
    if err != nil {
        return err
    }
    dsMysql.Db = db
    return nil
}

func (dsMysql *DataServiceMysql) Close() error {
    return dsMysql.Db.Close()
}

func (dsMysql *DataServiceMysql) GetReportItems() []ReportItem {
    result := []ReportItem{}
    rows, err := dsMysql.Db.Query(`select id, title, description,
                                   table_name, handler_key, create_time
                                   from report_conf`)
    if err != nil {
        log.Println("Error:", err)
        return result
    }
    defer rows.Close()
    var id, title, desc, createTime, tableName, handlerKey string
    for rows.Next() {
        err := rows.Scan(&id, &title, &desc, &tableName, &handlerKey, &createTime)
        if err != nil {
            log.Println("Error:", err)
        } else {
            log.Println("id:", id, "title:", title, "time:", createTime)
            reportItem := ReportItem{id, title, desc, createTime, tableName, handlerKey}
            result = append(result, reportItem)
        }
    }

    return result
}

func (dsMysql *DataServiceMysql) GetReportItem(id int) ReportItem {
    var idStr, title, desc, createTime, tableName, handlerKey string
    err := dsMysql.Db.QueryRow(`select id, title, description, table_name, handler_key, create_time
        from report_conf where id = ?`, id).
            Scan(&idStr, &title, &desc, &tableName, &handlerKey, &createTime)
    if err != nil {
        log.Println("Error:", err)
        return ReportItem{}
    } else {
        return ReportItem{idStr, title, desc, createTime, tableName, handlerKey}
    }
}

func (dsMysql *DataServiceMysql) QueryRow(sql string, dest ...interface{}) error {
    err := dsMysql.Db.QueryRow(sql).Scan(dest...)
    if err != nil {
        return err
    }
    return nil
}
