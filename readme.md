# 2016 UNITOSS tech study and repo
 This is repository of 2016 UNITOSS tech trend study.
 Feel Free to commit or pull request what you study and developement.

## UNITOSS Home repo
 [UNITOSS-HOME](https://github.com/hanzin/unitoss)

### [REDIS Install tutorial](http://redis.io/topics/quickstart)

### REFERENCES
 * [redis.io](http://redis.io/)
 * [redis-go-cluster](https://github.com/chasex/redis-go-cluster)
 * [cassandra official home](http://cassandra.apache.org/)
 * [cassandra data structure 설명 블로그](http://meetup.toast.com/posts/58)
 * [cassandra CQL go driver](https://github.com/gocql/gocql)

 * [elastic home page](https://www.elastic.co/)
 * [Logstash Reference (Getting Started)](https://www.elastic.co/guide/en/logstash/current/index.html)
 * [Elasticsearch Reference (Getting Started)](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html)
 * [Kibana Reference (Getting Started)](https://www.elastic.co/guide/en/kibana/current/index.html)
 * [ELK 실제 적용 예제](http://brantiffy.axisj.com/archives/418)

### CQL instructions
 sql 유사한 부분이 많으며, sql 에 익숙하면 쉽게 적응 가능.
```
DESCRIBE KEYSPACES; -- show list of all keyspaces
USE keyspace-name;
DESCRIBE TABLES; -- show list of all tables in current keyspace-name
SELECT * FROM table-name; -- sql 의 select 문과 동일.
```
