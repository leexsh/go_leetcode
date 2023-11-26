
/*
 from 作业帮
 有两个表，user表 user_id user_name, score表user_id score, user表中有50行数据， score表有45行数据，找出score表中score数据为空的user， 
    user:
        |user_id| user_name|
    score:
        |user_id| score|
    */

select *
from user left join score on user.user_id=score.user_id where score.score is null ;
