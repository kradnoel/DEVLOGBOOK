# DEVLOGBOOK

Simple Project Management WebApp for Developers

## Built with
- server: Golang + Sqlite3 + [Gorm](https://gorm.io/) + [Chi](https://github.com/go-chi/chi)
- client: [Node.js](https://nodejs.org/) + [Express](https://expressjs.com/) + [servue](https://github.com/futureaus/servue) + [Vue](https://vuejs.org/) + [Bootstrap](https://getbootstrap.com/) + [AdminLTE](https://adminlte.io/) + Javascript

## Features
- Add, Update or Delete Project;
- Add, Update or Delete Task;
- Add, Update or Delete Subtask;
- Get Status on Projects (Finished or Unfinished);
- Get time spent on Project;
- Get statistics on Projects (Total Projects, Completed, Avarage time spent);

## Set up

### Requirements
- [Node.js](https://nodejs.org/)
- [Golang](https://golang.org/)

### Local development

1. Clone this repository and `cd` into it

```bash
git clone https://github.com/kradnoel/DEVLOGBOOK.git
cd DEVLOGBOOK

```

For the server:

2. `cd` into server and install dependencies

```bash
cd server
go mod download
```

3. Run the server application

```bash
go run cmd/devlb.go
```

For the client:

4. `cd` into client and install dependencies

```bash
cd client
npm install
```

5. Run the client application

```bash
npm start
```

6.  Navigate to [http://localhost:3000](http://localhost:3000)

## License

[MIT](LICENSE)

Here's some Images:

![DEVLOGBOOK_1](https://github.com/kradnoel/portfolio_images/blob/master/DEVLOGBOOK/DEVLOGBOOK_1.png "DEVLOGBOOK 1")
![DEVLOGBOOK_2](https://github.com/kradnoel/portfolio_images/blob/master/DEVLOGBOOK/DEVLOGBOOK_2.png "DEVLOGBOOK 2")
![DEVLOGBOOK_3](https://github.com/kradnoel/portfolio_images/blob/master/DEVLOGBOOK/DEVLOGBOOK_3.png "DEVLOGBOOK 3")
