# Simple examples using `webgl-go`

## A colored triangle

- Copy `wasm_exec.js` from the standard `Go` installation in this folder (only once).

    ```
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    ```

- Compile the code

    ```
    GOARCH=wasm GOOS=js go build -o demo.wasm demos/triangle.go demos/shader_helpers.go
    ```

- Serve the web page

    ```
    go run server.go
    ```

- Open the following URL in a browser: <http://localhost:8080/>
