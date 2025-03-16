import * as dotenv from "dotenv";
import OpenAI from "openai";
import readline from "readline";

dotenv.config(); // load environment variables from .env file
// console.log(process.env); // remove this after you've confirmed it is working

// OpenAI configuration
// https://platform.openai.com/account/org-settings
// https://platform.openai.com/account/api-keys

// Check if environment variables are set
if (!process.env.OPENAI_API_KEY || !process.env.OPENAI_ORGANIZATION) {
  console.error("Error: Missing required environment variables (OPENAI_API_KEY or OPENAI_ORGANIZATION). Please check your .env file.");
  process.exit(1); // Exit the program if environment variables are missing
}

// Create new instance of OpenAI API using the configuration
const openai = new OpenAI({
  organization: process.env.OPENAI_ORGANIZATION, // set the organization using environment variable
  apiKey: process.env.OPENAI_API_KEY, // set the API key using environment variable
});

// Create a UI in the terminal that allows users to type in their questions
const userInterface = readline.createInterface({
  input: process.stdin, // set the input to the standard input
  output: process.stdout, // set the output to the standard output
});

// Configurable retry parameters
const maxRetries = 3;
const initialDelay = 2000; // Initial delay in milliseconds (2 seconds)

// Function to handle retries with exponential backoff
async function handleRequest(input, retries = 0) {
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
    // Log the error for monitoring
    console.error(`Error on attempt ${retries + 1}:`, e);

    if (retries < maxRetries) {
      const delay = initialDelay * Math.pow(2, retries); // Exponential backoff
      console.log(`Error occurred. Retrying in ${delay / 1000} seconds...`);
      setTimeout(() => handleRequest(input, retries + 1), delay); // Retry after delay
    } else {
      console.error("Error: Maximum retry attempts reached. Please try again later.");
    }
  }
}

// Prompt user to enter a message
function askQuestion() {
  userInterface.question('Please enter a prompt (type "exit" to quit): ', (input) => {
    if (input.toLowerCase() === 'exit') {
      console.log('Goodbye!');
      userInterface.close(); // Close the interface and stop further prompts
      return;
    }

    if (!input.trim()) {  // Check for empty input
      console.log("Input cannot be empty. Please enter a valid prompt.");
      askQuestion(); // prompt again if input is empty
      return;
    }

    handleRequest(input) // Call function with retry mechanism

      // Only after the request is handled and output is displayed, ask for another prompt
      .then(() => askQuestion()); // Ask for the next prompt after the current one is processed
  });
}

// Initial prompt
askQuestion();
