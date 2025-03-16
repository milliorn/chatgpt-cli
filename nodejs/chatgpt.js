import * as dotenv from "dotenv";
import OpenAI from "openai";
import readline from "readline";

dotenv.config(); // load environment variables from .env file
console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration
// https://platform.openai.com/account/org-settings
// https://platform.openai.com/account/api-keys

// Check if environment variables are set
if (!process.env.OPENAI_API_KEY || !process.env.OPENAI_ORGANIZATION) {
  console.error("Error: Missing required environment variables (API_KEY). Please check your .env file.");
  process.exit(1); // Exit the program if environment variables are missing
}

if (!process.env.OPENAI_ORGANIZATION) {
  console.error("Error: Missing required environment variables (ORGANIZATION). Please check your .env file.");
  process.exit(1); // Exit the program if environment variables are missing
}

// create new instance of OpenAI API using the configuration
const openai = new OpenAI({
  organization: process.env.OPENAI_ORGANIZATION, // set the organization using environment variable
  apiKey: process.env.OPENAI_API_KEY, // set the API key using environment variable
});

// create a UI in the terminal that allows users to type in their questions
const userInterface = readline.createInterface({
  input: process.stdin, // set the input to the standard input
  output: process.stdout, // set the output to the standard output
});

// Prompt user to enter a message
function askQuestion() {
  userInterface.question('Please enter a prompt (type "exit" to quit): ', async (input) => {
    if (input.toLowerCase() === 'exit') {
      console.log('Goodbye!');
      userInterface.close();
      return;
    }

    try {
      const res = await openai.chat.completions.create({
        model: "gpt-3.5-turbo", // set the model to use for the API
        messages: [{ role: "user", content: input }], // set the user's message as input for the API
      });

      if (res.choices && res.choices.length > 0 && res.choices[0].message) {
        console.log("ChatGPT:", res.choices[0].message.content);
      } else {
        console.log("No response content found.");
      }
    } catch (e) {
      console.error("Error:", e); // if there's an error, display it in the console
    }

    askQuestion(); // prompt the user for another message
  });
}

// Initial prompt
askQuestion();