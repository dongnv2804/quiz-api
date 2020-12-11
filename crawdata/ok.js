const categories = data.map(({ category, image, topics }) => {
  return {
    name: category,
    image,
    created_at: new Date(),
    updated_at: new Date(),
    topics,
  };
});
const insertCategory = (category) => {
  const { name, image, created_at, updated_at } = category;
  connection.query(
    'INSERT INTO categories SET ?',
    { name, image, created_at, updated_at },
    function (error, results, fields) {
      if (error) throw error;
      console.log('The solution is: ', results);
      insertTopic(results.insertId, category.topics);
    }
  );
};
const insertTopic = (categoryId, topics) => {
  topics.map((topic) => {
    const newTopic = {
      name: topic.name,
      category_id: categoryId,
      created_at: new Date(),
      updated_at: new Date(),
    };
    connection.query(
      'INSERT INTO topics SET ?',
      newTopic,
      function (error, results, fields) {
        if (error) throw error;
        insertQuestion(results.insertId, topic.questions);
      }
    );
  });
};
const insertQuestion = (topicId, questions) => {
  questions.map((question) => {
    const newQuestion = {
      content: question.name,
      topic_id: topicId,
      created_at: new Date(),
      updated_at: new Date(),
    };
    connection.query(
      'INSERT INTO questions SET ?',
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
      'INSERT INTO answers SET ?',
      newAnswer,
      function (error, results, fields) {
        if (error) throw error;
        console.log('The solution is: ', results);
      }
    );
  });
};
categories.map((category) => {
  insertCategory(category);
});