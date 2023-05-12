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
