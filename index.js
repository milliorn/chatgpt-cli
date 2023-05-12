import { Configuration, OpenAIApi } from "openai";
import readline from "readline";
import * as dotenv from "dotenv"; // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import

dotenv.config();
console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration

const configuration = new Configuration({
  organization: process.env.ORGANIZATION,
  apiKey: process.env.API_KEY,
});

// create new instance of OpenAI API
const openai = new OpenAIApi(configuration);

// calls the createChatCompletion, triggers an endpoint (https://api.openai.com/v1/chat/completions).
// function accepts an object of arguments, the model of chatGPT in use, an array of messages between user and AI.
// each message is an object containing the role, sent the message.
// value can be assistant, if from the AI or user,
// when the message originates user, and the content.
// the code prints the response from the AI.
openai
  .createChatCompletion({
    model: "gpt-3.5-turbo",
    messages: [{ role: "user", content: "Hello" }],
  })
  .then((res) => {
    console.log(res.data.choices[0].message.content);
  })
  .catch((e) => console.error(e));
