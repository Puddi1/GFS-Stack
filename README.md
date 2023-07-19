# GFS Stack

The quickest and simplest GO monolith stack to bring you from idea to MVP ready for production in days.

Built with:  
[Fiber](https://docs.gofiber.io/)  
[Vite](https://vitejs.dev/)  
[HTMX](https://htmx.org/)  
[Tailwind](https://tailwindcss.com/)  
[DaisyUI](https://daisyui.com/)  
[Supabase](https://supabase.com/)  
[Stripe](https://stripe.com/)
[entr(1)](https://eradman.com/entrproject/)

## Intro and Credits

Thanks for checking out the stack and being interested. Feel free to open issues related to this repo and make adjustements to the stack with pull requests.

Big credit to [Anthony](https://github.com/anthdm) for the philosophy.

## To Start

1. Clone the repository, which is a boilerplate repo to start all of your projects:

    ```git
    git clone https://github.com/Puddi1/GFS-Stack.git
    ```

2. Run the initialization:

    ```bash
    chmod +x init.sh
    ./init.sh
    ```

    You'll be asked to insert your environment variables. Here's a quick overview:
    Variables that are preceeded by `SUPABASE_` can all be found in your supabase's project [dashboard](https://supabase.com/dashboard)

    `SUPABASE_PROJECT_PASSWORD` is your supabase project's password string.  
    `SUPABASE_API_PUBLIC_KEY` is the `anon_key` string.  
    `SUPABASE_API_PRIVATE_KEY` is the `service_role` string.  
    `SUPABASE_DB_HOST` is the full databse URL <https://yourProjectSubdomain.supabase.co>.  
    `SUPABASE_DB_PORT` is the databse port.  
    `SUPABASE_DB_NAME` is the databse name.  
    `SUPABASE_DB_USER` is the databse user.  
    `SUPABASE_DB_SSLMODE` is either disable string or verify-full string to enable it.  
    `SUPABASE_DB_SSLCERT_PATH` if you insert `verify-full`on `SUPABASE_DB_SSLMODE` this is the path to the ssl certificate from the root of the project.

    `STRIPE_API_KEY` is the stripe API key string.

    `PORT` is the port you wish the server listens to, skipping it will default to port 3000
    `DEVELOPMENT` is if you want to run the build on development (Y) or production (N). You can change it after in the .env file: true for development, false for production

3. Run the server:

    If you want to run it as production, thus DEVELOPMENT set to false (make stands for the first make command: `make run`)

    ```bash
    make
    ```

    Note, by using the production environment you make Vite run `vite build`, that creates a static folder `./dist` with compressed assets and miscellania to improve both stability and performances. Then the GO server uses the static folder to render the HTML and related files. It doesn't refresh templates if you change them because they are meant to be used as static, in development it is the opposite.

    If you want to run it as development, thus DEVELOPMENT set to true.

    ```bash
    make dev
    ```

    Note, by using the development environment you and the GO server will work directly in the `./src` folder, with no compressed files and with CDNs. HTML is rendered each time it is served by the GO server. To enhance usage I suggest to use the quick reload escape command for the browser.

    Important: routes from devlopment to production don't change. If your build runs in development it will run even in production.  
    You start development with fast and immediate frontend changes, a quick hot-reload for the GO backend that you don't have to monitor. Then you move in pre-production with a static frontend composed by Vite, where you can still make changes to the go infrastructure, still with hot-reload. In conclusion for production you have your GO server that uses static assets and if prefereed no hot-reload for the server.

4. To change quickly environment:

    <!-- to implement -->

    Simply run

    ```bash
    make environment
    ```

## To Production

To deploy the stack to rpoduction

## Useful Commands

This is a list of pre-configured commands you can use:

-   `make build`  
    Will to complie the GO app in binary executable

-   `make dev`  
    Will first complie the GO app in binary and then execute it

-   `make vite`  
    Will make vite build the static frontend, same as running `npm run build` with build as `vite build --emptyOutDir`

-   `make hotreload-run`  
    to use entr(1) to hot reload if any changes to the active production directories occours (all GO files and src directory, excluded: bin, dist, node_modules, .gitignore, README.md, go.sum).  
    Note: this command requires [entr(1)](https://eradman.com/entrproject/) to be installed.  
    Shut down is being forced. How can you make it shut down gracefully? Make also it do it

-   `make hotreload-dev`  
    Will use entr(1) to hot reload if any changes to the active development directories occours (all GO files, excluded: bin, dist, node_modules, .gitignore, README.md, go.sum and src).  
    Note: this command requires [entr(1)](https://eradman.com/entrproject/) to be installed.  
    Shut down is being forced. How can you make it shut down gracefully? Make also it do it

-   `make clean`  
    Will run `go mod tidy`, cleaning your go.mod and go.sum setup based on packages usage.

## Stack Structure

README.md is used as documents fot this stack.

Makefile is the file to handle all execution, like building and running the app as default actions. You can add more scripts to suit your needs.

.gitignore is the directory to ignore. default: bin && .env

.env is the enviroment variables file, all your secrets will be placed here.

go.mod defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version. It needs to be initialized for every package that is outside the $GOPATH/src. After initialization, 'go get <packagename>', 'go build', 'go test' will handle all go.sum updates, to clean go.sum fron unused package simple run 'make clean', that will run the command 'go mod tidy' (https://go.dev/blog/using-go-modules)

go.sum contains the expected cryptographic hashes of the content of specific module versions. It is handled and auto generated by go CLI, do not modify the file.

main.go is the main execution file, this is the file that is run when you run the application, the core of your app.

bin folder contains the go binary to be run by the machine.

data folder contains all the data you have to define.

stripe folder creates and initialize the connection with stripe and provides a single variable to interact with it.

database folder creates and initialize the connection with your database and provides a single variable to interact with it.

scripts folder contains all single worker scripts you have to run locally on your server.

utils folder contains all shared util functions.

## Basics - How to move around

add cdn: <script src="https://unpkg.com/htmx.org@1.9.2" integrity="sha384-L6OqL9pRWyyFU3+/bjdSri+iIphTN/bvYyM37tICVyOJkWZLpP2vGn6VUEXgzg6h" crossorigin="anonymous"></script>
to use htmx on html page
note: not good i production, how to solve?

note: using htmx with get request and response from backend with playin html, you'll be able to create a quick and simple app

note: js files must be type="module" to be build by vite

note: custom routing added

htmx: simply import script on html
htmx: htmx request's paths must follow your same backend response's paths
htmx: minification eval error, evaluate

5. entr1 graceful shutdown // hard
6. supebase ssl // medium
7. backend related stuff // medium
8. superbase mysequell // ez
9. deployment // medium
10. dns? // medium
11. how to scale // ez
12. Login with third party auth
13. user token management clarify

<!-- github.com/stripe/stripe-go/v74 v74.25.0 // indirect
github.com/sujit-baniya/flash v0.1.8 // indirect -->

https://sujit-baniya.gitbook.io/fiber-boilerplate/additional-libraries

pocketbase -> on stable version

error: miscellania with eval
error: Generated an empty chunk: "login/user".
add note on how to deal also with go work
