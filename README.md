# Harry Potter CLI App

## Requirements 
- Random spell generator
- Character lookup 
    - Can search via name -> return info 
    - Can give list of characters based on criteria e.g. house, gender, age etc


## CLI Examples 

### Spells
`hogwarts spells` : base command  

Flags:   
`random` : return a random spell   
`lookup --name` : return a spell with the exact name

### Characters 
`hogwarts character` : base command

Flags:  
`random` : return a random character   
`lookup --name` : return the character with the exact name  
`list [flags]` : return a list of all characters that fit the criteria 

- `--house string`: filter by house
- `--gender string`: filter by gender
- `--patronus string` : filter by patronus
- `--alive bool` : filter by life status
