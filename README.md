## Pain Bot

**Contributors:** Jake Sigman, Jacob Koziej  
**Email:** <jsigman04@gmail.com>  
**Created:** August 10, 2021

### Description

This bot serves to help me learn Go.

### Initialization

1. Create a file in the same directory of `main.go` called `init.txt`
2. The `init.txt` file should look exactly as follows (ignore any square brackets):

    ```
    TOKEN: [Bot token]

    IGNORE
    [Server ID 1]
    [Server ID 2]

    AUTORESPONSES
    [Prompt 1]/[Response 1]
    [Prompt 2]/[Response 2]
    ```

3. Create a file in the same directory of `main.go` called `count.txt`.
4. The bot will also use a file that responds to DMs with messages in `messages.txt`. The file should be in the same directory as `main.go` and should look as follows (ignore any square brackets):

    ```
    [Message 1 Content]
    ---END MESSAGE---
    [Message 2 Content]
    ---END MESSAGE---
    ```

### Commands

**$help:** Opens list of commands.   
**$ping:** Pings the bot.   
**$paincount:** Informs user how many times "pain" was said.   
**$git:** Returns git repository information.   
**$pain:** Pain.   
**$autoresponse `{prompt}` `{response}`:** Adds autoresponse to bot.    
**$autoresponses**: Sends list of all current autoresponses.    
**$remove `{prompt}` :** Removes autoresponse from bot with provided prompt.     
