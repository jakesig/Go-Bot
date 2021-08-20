/* git.js
** Pain Bot
** Author: Jake Sigman
** This file contains the code for the git function.
 */


// Library imports

const Discord = require('discord.js');
const fs = require('fs');
const git = require('git-last-commit');

// Constant and function declarations

const client = new Discord.Client();
console.log = function(){};

// Log to console
console.info(`JS: Reading config file...`);

// Read init.txt file

fs.readFile('init.txt', 'utf8', function(err, data) {
  if (err) throw err;
  const lines = data.split("\n");
  console.log(lines)

  //Determines token, then logs in bot

  const token = lines[0].split(': ')[1];
  client.login(token);
});

// When client is ready, log that the client logged in successfully.

client.on('ready', () => {
    console.info(`JS: Logged in as ${client.user.tag}`);
});

// Code for $git command.

client.on('message', (msg) => {
  if (msg.content.toLowerCase() == "$git") {
    //Send embed with name of latest commit

    var commitname=git.getLastCommit(function(err, commit) {
      
      if (err)
        throw err;
      msg.channel.bulkDelete(1);

      //Determine commit information, including the date

      var commitid=Object.values(commit)[1]
      var name=Object.values(commit)[2];
      var branch=Object.values(commit)[10];
      var s = new Date(Object.values(commit)[6]*1000).toLocaleDateString("en-US", {timeZone: "America/New_York"});

      //Construct and send embed

      const embed = new Discord.MessageEmbed()
        .setColor('#f1c40f')
        .setTitle("GitHub Repository: jakesig/Pain-Bot")
        .setDescription("https://github.com/jakesig/Pain-Bot\n\n__**Latest Commit**__\n**Message: **"+name+"\n**Branch: **"+branch+"\n**Date: **"+s+"\n**ID: **"+commitid)
        .setThumbnail("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRAi68fw3hBkE6l-vGLWYB9aRoSV5DWJ0zKJtAzpjYTMD83DwP5WU4D1N7eHx1ucPcZle8&usqp=CAU");
        
      msg.channel.send(embed);
    });
  }
});