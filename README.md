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

- Level 1: 1 written (message / charater)
- Level 2: 10 written (messages / charaters)
- Level 3: 50 written (messages / charaters)
- Level 4: 100 written (messages / charaters)
- Level 5: 500 written (messages / charaters)
- Level 6: 1 000 written (messages / charaters)
- Level 7: 3 250 written (messages / charaters)
- Level 8: 5 500 written (messages / charaters)
- Level 9: 7 750 written (messages / charaters)
- Level 10: 10 000 written (messages / charaters)
- Level 11: 50 000 written (messages / charaters)
- Level 12: 100 000 written (messages / charaters)
- Level 13: 1 000 000 written (messages / charaters)
- Level 14: 2 000 000 written (messages / charaters)
- Level 15: 3 000 000 written (messages / charaters)
- Level 16: 4 000 000 written (messages / charaters)
- Level 17: 5 000 000 written (messages / charaters)
- Level 18: 6 000 000 written (messages / charaters)
- Level 19: 7 000 000 written (messages / charaters)
- Level 20: 9 000 000 written (messages / charaters)
- Level 21: 10 000 000 written (messages / charaters)
- Level 22: 100 000 000 written (messages / charaters)
- Level 23: 200 000 000 written (messages / charaters)
- Level 24: 300 000 000 written (messages / charaters)
- Level 25: 400 000 000 written (messages / charaters)
- Level 26: 500 000 000 written (messages / charaters)
- Level 27: 600 000 000 written (messages / charaters)
- Level 28: 700 000 000 written (messages / charaters)
- Level 29: 800 000 000 written (messages / charaters)
- Level 30: 900 000 000 written (messages / charaters)
- Level max: over 1 000 000 000 written (messages / charaters)
