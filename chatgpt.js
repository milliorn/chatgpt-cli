import { Configuration, OpenAIApi } from "openai";
import readline from "readline";
import * as dotenv from "dotenv";

dotenv.config(); // load environment variables from .env file
// console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration
// https://platform.openai.com/account/org-settings
// https://platform.openai.com/account/api-keys
const configuration = new Configuration({
  organization: process.env.ORGANIZATION, // set the organization using environment variable
  apiKey: process.env.API_KEY, // set the API key using environment variable
});

// create new instance of OpenAI API using the configuration
const openai = new OpenAIApi(configuration);

// creates a UI in the terminal that allows users to type in their questions.
const userInterface = readline.createInterface({
  input: process.stdin, // set the input to the standard input
  output: process.stdout, // set the output to the standard output
});

// prompt the user to enter a message
userInterface.prompt();

// when the user gives a prompt and hits Enter, 
// it triggers a callback which takes the input and uses it as content for the OpenAI API
// when the response is displayed, the user is prompted for another message
userInterface.on("line", async (input) => {
  await openai
    .createChatCompletion({
      model: "gpt-3.5-turbo", // set the model to use for the API
      messages: [{ role: "user", content: input }], // set the user's message as input for the API
    })
    .then((res) => {
      console.log(res.data.choices[0].message.content); // display the response from the API
      userInterface.prompt(); // prompt the user for another message
    })
    .catch((e) => {
      console.error(e); // if there's an error, display it in the console
    });
});
