const express = require('express')
const path = require('path')

const app = express()
const port = process.env.PORT || 8888

app.get('/', (request, response) => {
  response.sendFile(path.resolve(__dirname, 'index.html'))
})

app.listen(port)

console.log("server started on port http://localhost:" + port)
