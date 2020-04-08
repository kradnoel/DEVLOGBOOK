const Servue = require("servue")
const express = require("express")
const path = require("path")
const dotenv = require("dotenv").config()
const axios = require("axios")
const debug = require('debug')('express')

const port = process.env.PORT || 3000

var app = express()

app.use(express.static(path.join(__dirname, 'public')))

// Sets default data from .env
app.set('host', process.env.HOST)
app.set('port', port)

var servue = new Servue()

//Sets views folder path which contains .vue .css .less .js and etc
servue.resources = path.resolve(__dirname, "src")

servue.mode = "production"

//path of node_modules
servue.nodemodules = path.resolve(__dirname, 'node_modules')

//custom html body template
servue.template = (content, context, bundle) => (`
<!DOCTYPE html>
<html${ context.htmlattr ? ' ' + context.htmlattr : '' }>
  <head>
    ${context.head ? context.head : ''}
    ${context.renderResourceHints()}
    ${context.renderStyles()}
    ${context.renderState({windowKey: '__INITIAL_STATE__', contextKey: "data"})}
    ${context.renderState({windowKey: '__INITIAL_ROOTSTATE__', contextKey: "state"})}
  </head>
  <body class="hold-transition sidebar-mini layout-fixed pace-primary">
    ${content}
    <script>${bundle.client}</script>
  </body>
</html>
`)

// Home 
app.get('/', async (req, res) => {
  debug('app:express')
  res.send(await servue.render('pages/dashboard'))
})

// Project
app.get('/projects', async(req, res) => {
  res.send(await servue.render('pages/project'))
})

// Task
app.get('/projects/:id', async(req, res) => {
  var id = req.params.id
  let data = { "id": id }
  res.send(await servue.render('pages/task', {data}))
})

// Subtask
app.get('/projects/:pid/:tid', async(req, res) => {
  var pid =  req.params.pid
  var tid = req.params.tid
  let data = { "pid": pid, "tid": tid}
  res.send(await servue.render('pages/subtask', {data}))
})

app.listen(app.get('port'), () => console.log(`listening on port ${port}!`))
