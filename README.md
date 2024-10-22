# Simple Go Struct-Interface Method Example

### Overview
This Go program demonstrates the use of interface-typed objects that are declared to link some functions to their specified struct. This technique, although using the property of interface-typed objects, is not the same as demonstrated [here](https://github.com/pkx8326/simple_go_struct-interface_example/tree/main) where interface-typed arguments are passed into functions. As for the code in this repository,  We can use functions as methods for the created structs since we are not passing interface-typed objects as conventional argument, but we are using them as receiver arguments instead.

### Program manual
Since this program is very similar to the one in the repository in this [link](https://github.com/pkx8326/simple_go_struct-interface_example/tree/main), I'm just going to copy-paste the program manual from there as the following:

When run, the program asks the user to input the following information in the following order:

Note title
Note content
Todo
Then, the program will show messages containing the input information and, if there's no error, will notify the user that the note and todo are successfully saved (in json file format).

There is no input validation for this program because every piece of information are in the form of free text. However, the program is designed to catch error when saving the files. The user will be notified if there's any error while saving each file. In case of error, the program will stops after displaying the error message.

### Code structure

Although, the program in this project works exactly the same as the one from this [link](https://github.com/pkx8326/simple_go_struct-interface_example/tree/main), the code structure was designed differently in order to demonstrate another way of using interface to link common methods to different struct types from different packages.

The project comprises the ```main.go``` file which contains the code of the main program, and the codes which make up the ```note``` and ```todo``` packages. The ```main.go``` file contains the code declaring interface objects that link the ```save``` and ```display``` functions in the mentioned packages to their native structs. Those functions can, in turn, be used as the structs' methods.

### Program flow

Since this program works exactly as another one in a different repository as mentioned in several places above, I'm just going to copy an paste the same program flow from there as the following:

1. The user inputs the note title as a string
2. The user inputs the note content as a string
3. The program takes those inputs to create a struct which stores the note title, content, and the timestamp at its creation
4. The user inputs the todo text
5. The program takes the todo text input to create another struct which stores the todo text (without any title)
6. The program displays messages to confirm the note's title and its content from the inputs
7. The program displays a message to confirm the todo
8. The program displays a message to notify the user that it's saving the note
9. The program attempts to save the inputs as a json file with json field names according to the struct tags given in the code
10. The program displays the message that it saves the file successfully
11. The program repeat the same process from 8. for the todo (the todo's file name is already hard-coded in the program and can't be changed)
