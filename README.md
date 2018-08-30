# Telemetry and Unicorns🦄


Here is a preview of the project with data recorded from Benji riding his bike in SF 🤓
There is also a HQ preview [here](preview/preview_hq.gif)
![Gif preview](preview/preview.gif)

The goal of the project is to be able to receive telemetry data from the excellent [Sensorlog for iOS](https://itunes.apple.com/us/app/sensorlog/id388014573?mt=8).

The main goal of the project was to learn about Go. The server listens to incoming HTTP Post requests from Sensorlog and saves the data in Postgres using Gorm. It also broadcasts the newly received points to the web interface using WebSockets.


The server is written in Go and uses:
* [Gin](https://github.com/gin-gonic/gin) framework to handle HTTP requests
* [Gorm](https://github.com/jinzhu/gorm) for data persistence in Postgres
* [Gorilla WebSocket](https://github.com/gorilla/websocket/) to support WebSocket and pushing data to the browser as soon as the POST request is received on the server.

The web page uses:
* [Bootstrap](https://getbootstrap.com/) for the (very basic) layout
* [Chart.js](https://www.chartjs.org/) for the charts
* [Leaflet](https://leafletjs.com/) and [Mapbox](https://www.mapbox.com/) for the map
* [Cornify](https://www.cornify.com/) for the unicorns

