# Simple examples using `webgl-go`

- Copy `wasm_exec.js` from the standard `Go` installation in this folder (only once).

    ```
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    ```

- Compile the code

  - A colored triangle

    ```
    GOARCH=wasm GOOS=js go build -o demo.wasm demos/triangle.go demos/shader_helpers.go
    ```

  - A colored triangle using a texture generated from data

    ```
    GOARCH=wasm GOOS=js go build -o demo.wasm demos/generate_texture.go demos/shader_helpers.go
    ```

- Serve the web page

    ```
    go run server.go
    ```

- Open the following URL in a browser: <http://localhost:8080/>
