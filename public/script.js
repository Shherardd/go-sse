const ID = Math.random().toString()
let source = new EventSource(`http://localhost:8080/notify?id=${ID}`)

source.addEventListener('open', function (e) {
    console.log(`Connection opened with ID: ${ID}`)
})

source.addEventListener('saludo', function (e) {
    console.log(e)
    console.log(e.lastEventId)
})

source.addEventListener('bye', function (e) {
    console.log(e.data)
})

source.addEventListener('error', function (e) {
    if (e.readyState === EventSource.CLOSED) {
        console.log('Connection was closed.')
    }
})

function bye(source) {
    source.close()
}