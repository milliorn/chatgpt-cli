import { Configuration, OpenAIApi } from "openai";
import readline from "readline";
import * as dotenv from "dotenv"; // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import

dotenv.config();
// console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration

const configuration = new Configuration({
  organization: process.env.ORGANIZATION,
  apiKey: process.env.API_KEY,
});

// create new instance of OpenAI API
const openai = new OpenAIApi(configuration);

// creates a UI in the terminal that allows users to type in their questions.
const userInterface = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

// prompt the user to enter a message
userInterface.prompt();

// user gives prompt and hits Enter it triggers a callback
// which passes as input which is now used as content.
// when response is displayed the user is prompted for another message.
userInterface.on("line", async (input) => {
  await openai
    .createChatCompletion({
      model: "gpt-3.5-turbo",
      messages: [{ role: "user", content: input }],
    })
    .then((res) => {
      console.log(res.data.choices[0].message.content);
      userInterface.prompt();
    })
    .catch((e) => {
      console.error(e);
    });
});
