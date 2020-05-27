Sources used for these exercises:

- gobyexample.com/reading-files
- http://terokarvinen.com/2020/go-programming-course-2020-w22/

This program gets the password as an input and compares it to the specified password list. There is one password list in this repo found from danielmiessler/SecLists repo that can be found in GitHub.
Link to the password list: https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt

You can add the password by typing --p *password*

You can specify the password list by typing --l *file name*

If the password can be found from the list the program says that *inputted password* has been found from the list.

If the password is not on the list the program says that: "Password is not on the list"
