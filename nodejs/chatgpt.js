import * as dotenv from "dotenv";
import OpenAI from "openai";
import readline from "readline";

dotenv.config(); // load environment variables from .env file
// console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration
// https://platform.openai.com/account/org-settings
// https://platform.openai.com/account/api-keys

// Check if environment variables are set
if (!process.env.API_KEY || !process.env.ORGANIZATION) {
  console.error("Error: Missing required environment variables (API_KEY or ORGANIZATION). Please check your .env file.");
  process.exit(1); // Exit the program if environment variables are missing
}

// create new instance of OpenAI API using the configuration
const openai = new OpenAI({
  organization: process.env.ORGANIZATION, // set the organization using environment variable
  apiKey: process.env.API_KEY, // set the API key using environment variable
});

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
    .chat.completions.create({
      model: "gpt-3.5-turbo", // set the model to use for the API
      messages: [{ role: "user", content: input }], // set the user's message as input for the API
    })
    .then((res) => {
      // Accessing the first choice's message content
      if (res.choices && res.choices.length > 0 && res.choices[0].message) {
        console.log("ChatGPT:", res.choices[0].message.content);
      } else {
        console.log("No response content found.");
      }

      userInterface.prompt(); // prompt the user for another message
    })
    .catch((e) => {
      console.error("Error:", e); // if there's an error, display it in the console
    });
});
