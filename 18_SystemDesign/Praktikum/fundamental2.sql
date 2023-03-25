--1
HGETALL users //redis
MATCH (u:users)
--2
RETURN u.* //neo4j
--3
SELECT * FROM users; //cassandra