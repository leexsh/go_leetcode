/*
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| email       | varchar |
+-------------+---------+
id 是该表的主键列。
此表的每一行都包含一封电子邮件。电子邮件不包含大写字母。


编写一个 SQL 查询来报告所有重复的电子邮件。 请注意，可以保证电子邮件字段不为 NULL。
*/
select email from Person group by email having count(email) > 1