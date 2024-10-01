# Prompt Engineering Project

This project demonstrates the use of Prediction Guard's API for question answering based on provided context.

## Prerequisites

- Go programming language
- Git
- Prediction Guard API key

## Setup

1. Clone the repository:
```
git clone https://github.com/mesutoezdil/prompt_engineering_02.git
```

2. Navigate to the project directory:
```
cd prompt_engineering_02
```

3. Initialize the Go module:
```
go mod init prompt_engineering_02
```

4. Install the Prediction Guard Go client:
```
go get github.com/predictionguard/go-client
```

5. Set up your Prediction Guard API key as an environment variable:
```
export PGKEY="your_api_key_here"
```

6. Create a context file (e.g., `context1.txt`) and add some content to it:
```
touch context1.txt
```
Add content to `context1.txt`, for example:
```
Bitcoin is the first cryptocurrency created by Satoshi Nakamoto in 2009. It is decentralized and uses blockchain technology.
```

7. Create the main Go file:
```
touch main.go
```

8. Copy the provided code into `main.go`. Make sure to update the model name to "Hermes-2-Pro-Llama-3-8B" and adjust the `MaxTokens` and `Temperature` fields as shown in the previous updates.

## Running the Program

To run the program, use the following command:

```
go run main.go context1.txt
```

The program will start and prompt you to enter questions. Type your questions and press Enter. The AI will respond based on the context provided in `context1.txt`.

To exit the program, type "exit" and press Enter.

## Troubleshooting

If you encounter any issues:

1. Ensure your API key is correctly set as an environment variable.
2. Make sure you have the latest version of the Prediction Guard Go client:
```
go get -u github.com/predictionguard/go-client
```
3. Check that your `context1.txt` file contains relevant information for the questions you're asking.

## Example Usage

```
🧑: what is bitcoin?
🤖: Bitcoin is the first cryptocurrency created by Satoshi Nakamoto in 2009. It is decentralized and uses blockchain technology.
```

Remember to replace "your_api_key_here" with your actual Prediction Guard API key when setting up the environment variable.
