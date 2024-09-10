# Harry Potter CLI App  

A Harry Potter CLI tool that knows all spells and characters of the Wizarding World.  


## CLI Examples 

### Spells
`hogwarts spells` : base command  

Flags:   
`random` : return a random spell   
`lookup --name` : return a spell with the exact name

![Hogwarts Spells examples](images/hogwarts_spells.png)

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
