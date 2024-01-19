# Text-to-Speech Converter

This is a command-line application that converts text to speech using the VoiceRSS API.

## Installation

Clone the repository and navigate to the project directory. Then, install the necessary dependencies:

```bash
go install
```

## Get API Keys

You must signup to both RapidAPI and VoiceRSS to get their API keys to run this application. Be sure to click `Activate` API on your account details in VoiceRSS.

VoiceRSS = [http://www.voicerss.org/registration.aspx](http://www.voicerss.org/registration.aspx)
RapidAPI = [https://rapidapi.com/auth/sign-up](https://rapidapi.com/auth/sign-up)

## Usage
The application provides a `convert` command that converts the given text to speech. You can specify the language and the text to convert using the `--language` (or `-l`) and `--text` (or `-t`) flags, respectively.

Here's an example:

`go run . convert --language en-us --text "Hello, world!"`

This will convert the text "Hello, world!" to speech in US English.