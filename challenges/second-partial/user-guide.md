User Guide
==========
To setup the program you first must be in the second-partial folder. Then open two bash terminals,
one to run the program and the other to run the following commands.

To login use the command:
curl -u username:password http://localhost:8080/login
(Must be logged in to use any of the following commands)

To get the status use the command:
curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status
(replace <ACCESS_TOKEN> with the token you're given when you login)

To upload the image use the command:
curl -F "data=@path/to/local/image.png" -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/upload
(replace "data=@path/to/local/image.png" with the actual path of the image and <ACCESS_TOKEN> with your token)

To log out use the command:
curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout
(replace <ACCESS_TOKEN> with the token you're given when you login).