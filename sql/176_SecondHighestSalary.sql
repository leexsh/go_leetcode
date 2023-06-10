/*

第二高的薪水
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
id 是这个表的主键。
表的每一行包含员工的工资信息。

*/

select
    (
        select Distinct salary from Employee order by salary DESC LIMIT 1 offset 1
        ) as SecondHighestSalary