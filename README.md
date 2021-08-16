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

    AUTORESPONSES
    [Prompt 1]/[Response 1]
    [Prompt 2]/[Response 2]
    ```

3. Create a file in the same directory of `main.go` called `count.txt`.

### Commands

**$help:** Opens list of commands.   
**$ping:** Pings the bot.   
**$paincount:** Informs user how many times "pain" was said.   
**$git:** Returns git repository information.   
**$pain:** Pain.   
**$autoresponse `{prompt}` `{response}`:** Adds autoresponse to bot.    
**$autoresponses**: Sends list of all current autoresponses.      
