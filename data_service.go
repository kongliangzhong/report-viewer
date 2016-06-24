package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "strconv"
)

type DataService interface {
    Init() error
    GetReportItems() []ReportItem
    GetReportItem(id int) ReportItem
    QueryRow(sql string, dest ...interface{}) error
    GetDataByTimeline(tableName, dataField, timelineField string) ([]string, []string, error)
    SaveReportItem(ri ReportItem) error
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
                                   create_time from report_conf`)
    if err != nil {
        log.Println("Error:", err)
        return result
    }
    defer rows.Close()
    var id, title, createTime string
    var descBs []byte

    for rows.Next() {
        err := rows.Scan(&id, &title, &descBs, &createTime)
        if err != nil {
            log.Println("Error:", err)
        } else {
            log.Println("id:", id, "title:", title, "time:", createTime)
            reportItem := ReportItem{Id: id, Title: title, Desc: string(descBs), CreateTime: createTime}
            result = append(result, reportItem)
        }
    }

    return result
}

func (dsMysql *DataServiceMysql) SaveReportItem(ri ReportItem) error {
    stmt, err := dsMysql.Db.Prepare(`insert into report_conf(title, table_name, data_field,
        data_format, chart_type, legend, show_timeline, timeline_field)
        values(?,?,?,?,?,?,?,?)`)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(ri.Title, ri.TableName, ri.DataField, ri.DataFormat, ri.ChartType,
        ri.Legend, ri.ShowTimeLine, ri.TimeLineField)
    return err
}

func (dsMysql *DataServiceMysql) GetReportItem(id int) ReportItem {
    var idStr, title, createTime, tableName, timeLineFiled, legend,
        dataField, dataFormat, chartType string
    var showTimeLine bool
    var descBs, handlerKeyBs []byte
    err := dsMysql.Db.QueryRow(`select id, title, legend, description, table_name, data_field,
        data_format, chart_type, handler_key, show_timeline, timeline_field, create_time
        from report_conf where id = ?`, id).
        Scan(&idStr, &title, &legend, &descBs, &tableName, &dataField, &dataFormat, &chartType,
            &handlerKeyBs, &showTimeLine, &timeLineFiled, &createTime)
    if err != nil {
        log.Println("Error:", err)
        return ReportItem{}
    } else {
        return ReportItem{idStr, title, legend, string(descBs), createTime, tableName, dataField, dataFormat,
            chartType, string(handlerKeyBs), showTimeLine, timeLineFiled}
    }
}

func (dsMysql *DataServiceMysql) QueryRow(sql string, dest ...interface{}) error {
    err := dsMysql.Db.QueryRow(sql).Scan(dest...)
    if err != nil {
        return err
    }
    return nil
}

func (dsMysql *DataServiceMysql) GetDataByTimeline(tableName, dataField, timelineField string) (datas []string, timeLines []string, err error) {
    var genSql = func() string {
        maxDataLen := 7
        sql := "select " + dataField + ", " + timelineField + " from " + tableName
        sql = sql + " order by " + timelineField + " desc "
        sql = sql + " limit " + strconv.Itoa(maxDataLen)
        return sql
    }

    sql := genSql()
    rows, err := dsMysql.Db.Query(sql)
    if err != nil {
        return
    }
    defer rows.Close()

    var dataBs, timeLineBs []byte
    for rows.Next() {
        err := rows.Scan(&dataBs, &timeLineBs)
        if err != nil {
            log.Println("Error:", err)
        } else {
            datas = append(datas, string(dataBs))
            timeLines = append(timeLines, string(timeLineBs))
        }
    }
    return
}
