const data = require("./data.json");
const mysql = require("mysql");
var connection = mysql.createConnection({
  host: "localhost",
  user: "root",
  password: "dongnv2804",
  database: "quizdb",
});
connection.connect();
const topics = data.topics.map(({ name, questions }) => {
  return {
    name: name,
    questions,
    created_at: new Date(),
    updated_at: new Date(),
  };
});

const insertTopic = (topic) => {
  const { name, created_at, updated_at } = topic;
  connection.query(
    "INSERT INTO topics SET ?",
    { name, created_at, updated_at },
    function (error, results, fields) {
      if (error) throw error;
      insertQuestion(results.insertId, topic.questions);
    }
  );
};
const insertQuestion = (topicId, questions) => {
  questions.map((question) => {
    const newQuestion = {
      content: question.name,
      score: question.score,
      topic_id: topicId,
      created_at: new Date(),
      updated_at: new Date(),
    };
    connection.query(
      "INSERT INTO questions SET ?",
      newQuestion,
      function (error, results, fields) {
        if (error) throw error;
        insertAnswer(results.insertId, question.answers);
      }
    );
  });
};
const insertAnswer = (questionId, answers) => {
  answers.map((answer) => {
    const newAnswer = {
      question_id: questionId,
      content: answer.content,
      is_correct: answer.isCorrect,
      created_at: new Date(),
      updated_at: new Date(),
    };
    connection.query(
      "INSERT INTO answers SET ?",
      newAnswer,
      function (error, results, fields) {
        if (error) throw error;
        console.log("The solution is: ", results);
      }
    );
  });
};
topics.map((value) => {
  insertTopic(value);
});
