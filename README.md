
## 宜信报表图形化展示

1. 编译执行：  
    * 编译：go build  
    * 执行：./report-viewer, 然后在浏览器localhost:9000查看结果。  
    * 数据准备：建表语句以及数据导入脚本在 shell文件夹中。  


2. 功能描述：  
    * 支持灵活的数据格式，指定表名，字段名，然后采用约定的格式存储数据，就可以自定义报表展示项。  
    * 支持饼图，柱图，折线图等。
    * 支持timeline展示，即可以展示同事数据不同时间的图表（如一周内每天数据的变化），进行对比。

