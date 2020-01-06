# Letter counter 
Discord letter counter bot

## Presentation

This discord bot count letters in users message.  
It also count message of users.

## Usage

When you write messages, the bot count letters in it.  
If you want to know your score, just call the bot with the command `§score`.

The output will be some thing like this:

```txt
User: username#discriminator

Written characters: number
Caracter level: number

Messages sent: number
Message level: number

Ratio (written caracters/messages sent): floating number

Rank: #number
```

A hight ratio mean you write a lot by message.
A low ratio mean you send a lot message with little content.

The Rank is your place with all other people in the world.

## Implementation

### Level

Levels are calculed as follow for both characters and messages:

- Level 1: 1 written (message / character)
- Level 2: 10 written (messages / characters)
- Level 3: 50 written (messages / characters)
- Level 4: 100 written (messages / characters)
- Level 5: 500 written (messages / characters)
- Level 6: 1 000 written (messages / characters)
- Level 7: 3 250 written (messages / characters)
- Level 8: 5 500 written (messages / characters)
- Level 9: 7 750 written (messages / characters)
- Level 10: 10 000 written (messages / characters)
- Level 11: 50 000 written (messages / characters)
- Level 12: 100 000 written (messages / characters)
- Level 13: 1 000 000 written (messages / characters)
- Level 14: 2 000 000 written (messages / characters)
- Level 15: 3 000 000 written (messages / characters)
- Level 16: 4 000 000 written (messages / characters)
- Level 17: 5 000 000 written (messages / characters)
- Level 18: 7 000 000 written (messages / characters)
- Level 19: 9 000 000 written (messages / characters)
- Level 20: 10 000 000 written (messages / characters)
- Level 21: 100 000 000 written (messages / characters)
- Level 22: 200 000 000 written (messages / characters)
- Level 23: 300 000 000 written (messages / characters)
- Level 24: 400 000 000 written (messages / characters)
- Level 25: 500 000 000 written (messages / characters)
- Level 26: 600 000 000 written (messages / characters)
- Level 27: 700 000 000 written (messages / characters)
- Level 28: 800 000 000 written (messages / characters)
- Level 29: 900 000 000 written (messages / characters)
- Level max: over 1 000 000 000 written (messages / characters)
