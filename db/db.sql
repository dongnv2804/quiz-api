drop database quizdb;
create database quizdb;
use  quizdb;

create table topics(
    id int auto_increment primary key ,
    name nvarchar(500),
    created_at datetime,
    updated_at datetime
);
create table questions(
    id int auto_increment primary key ,
    content nvarchar(1000),
    score int,
    created_at datetime,
    updated_at datetime,
    topic_id int,
    constraint fk_question_topic foreign key(topic_id) references topics(id)
);
create table answers(
    id int auto_increment primary key ,
    content nvarchar(500),
    is_correct boolean,
    created_at datetime,
    updated_at datetime,
    question_id int,
    constraint fk_answer_question foreign key(question_id) references questions(id)
);