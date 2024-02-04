# Discord_bot
Overview
This Discord bot is a multi-tool that provides various functionalities, including setting reminders, creating polls, and fetching current weather information.
m
<img src="https://github.com/Sulafit/Discord_bot/assets/95477101/c9e9ae08-c7ad-4527-852b-12ab104eeecf" data-canonical-src="https://gyazo.com/eb5c5741b6a9a16c692170a41a49c858.png" width="400" height="400" />

Features
1. Reminder System
Users can set reminders with a specified message and duration. The bot will notify users when the time elapses.

Usage: !remindme <message> <duration>

Example: !remindme Take a break! 1h

2. Poll Creation
Users can create polls with a title and multiple options. The bot displays the poll and allows users to vote by reacting to the options.

Usage: !createpoll <title> <option1> <option2> ...

Example: !createpoll Favorite Color red blue green

3. Weather Information
Users can fetch the current weather information for a specified location. The bot retrieves the temperature and description of the weather.

Usage: !currentweather <location>

Example: !currentweather London

Installation
Clone this repository to your local machine.

Install dependencies:

arduino
Copy code
go get github.com/bwmarrin/discordgo
Obtain a Discord bot token and replace "YOUR_DISCORD_BOT_TOKEN_HERE" in the code with your token.

Optionally, obtain an OpenWeatherMap API key for fetching weather information and replace "YOUR_OPENWEATHERMAP_API_KEY_HERE" in the code with your API key.

Build and run the bot:

bash
Copy code
go build
./multi_tool_bot
Invite the bot to your Discord server using the OAuth2 URL generated by Discord.

Usage
Invite the bot to your Discord server.

Use the provided commands in any text channel where the bot is present.

Follow the command syntax and examples provided in the Features section.

Contributions
Contributions are welcome! Feel free to open issues or pull requests to suggest improvements or report bugs.

