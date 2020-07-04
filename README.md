# orderbackend
Developer test for packform.io backend

Install and run this application first before installing and running the viewer application

This application was created using:
**Windows 10**
**GO version go1.14.4 windows/amd64**

You will need to have access to a PostgreSQL database

Steps to install this application
1. Clone this repository to your computer
`git clone https://github.com/vanjoechua/orderbackend.git <installation-directory>`

2. Enter the installation directory

3. Edit the .env file and update the variables to access your PostgreSQL database
`DB_HOST=<database-host>`
`DB_DRIVER=postgres`
`DB_USER=<database-user>`
`DD_PASSWORD=<user-password>`
`DB_NAME=<database-name>`
`DB_PORT=5432`

4. In the setup directory, load the developer_test.sql file to create the tables for the database

5. CD to the installation root directory. Run the program:
`go run main.go`

This will run a REST API server on http://localhost:8080