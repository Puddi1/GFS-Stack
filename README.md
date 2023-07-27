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

0. Have all external dependencies installed in you system:

    - [entr(1)](https://eradman.com/entrproject/)
      Download and install the package

    - [Flyio](https://fly.io)
      Follow the [installation](https://fly.io/docs/hands-on/install-flyctl/) instruction an [log in to Fly](https://fly.io/docs/getting-started/log-in-to-fly/)

1. Clone the repository, which is a boilerplate repo to start all of your projects:

    ```git
    git clone https://github.com/Puddi1/GFS-Stack.git
    ```

2. Run the initialization:

    ```bash
    make init
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

4. To change quickly between production and development environment:

    Simply run

    ```bash
    make e
    ```

5. To run the build with hot reload

    Simply run:

    ```bash
    make hot
    ```

    Remember that you'll need to have [entr(1)](https://eradman.com/entrproject/) to be installed to use this command.  
    You can use it both interchangibly for development or if you have a last minute production change to hot reload when any backend or environment changes occur.  
    To change files that are being watched you have to modify `./scripts/bash/hotReloadGo.sh`
    <!-- Shut down is being forced. How can you make it shut down gracefully? Make also it do it -->

## To Production

To deploy the stack to poduction as smooth as possible you'll need to have [Flyio](https://fly.io) CLI installed, workflow:

0. If you aren't logged in, do so:

    Login:

    ```bash
    fly auth login
    ```

    Signup:

    ```bash
    fly auth signup
    ```

1. Launch the app:

    ```bash
    flyctl launch
    ```

2. Deploy the app:

    ```bash
    flyctl deploy
    ```

Otherwise to deploy on your custom server, simply clone your repo, inititalize the project as production and run the app:

1. Clone:

    ```git
    git clone https://github.com/<YourUsername/<RepositoryName>.git
    ```

2. Initialize:

    ```bash
    make init
    ```

3. Run:

    ```bash
    make
    ```

## Useful Commands

This is a list of pre-configured commands you can use:

-   `make build`  
     Will to complie the GO app in binary executable

-   `make dev`  
     Will first complie the GO app in binary and then execute it

-   `make vite`  
     Will make vite build the static frontend, same as running `npm run build` with build as `vite build --emptyOutDir`

-   `make clean`  
     Will run `go mod tidy`, cleaning your go.mod and go.sum setup based on packages usage.

## Stack Structure

<!--  -->

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

Quick overview about features and how to leverage them

-   Stripe
    Stripe is the easies, fastest, and most reliable method of payment that the market can offer. The stack implements its [checkout session feature](https://stripe.com/docs/payments/checkout) as it is the quickest checkout workflow possible, while still having one of the highest conversions rates, offered by its brand familiarity and easiness of use.  
    To modify the checkout page go to your [stripe checkout settings](https://dashboard.stripe.com/settings/checkout).  
    Moreover, to help your users manage their profile and subscriptions we have impllemented the [stripe customer dashboard](https://stripe.com/docs/no-code/customer-portal).  
    To modify the dashboard page go to your [stripe customer dashboard settings](https://dashboard.stripe.com/settings/billing/portal).  
    Any creation or modification of any of your stripe products is done programattically. That's because to create an MVP where speed matters doing it by hand can be fatser.  
    Our stripe implementation regarding [products and relative subcategories](https://stripe.com/docs/api/products?lang=go) will give you the possibility to fetch infos so that you can update the UI programmatically, search, update, create and delete Products, Prices, Coupons, Promotion codes and Discounts.  
    [Webhooks](https://stripe.com/docs/webhooks) are essential to coordinate movements between your app and the stripe balance. You can check all [events](https://stripe.com/docs/api/events/types) that can be triggered and choose to which one subscribe with our [webhooks API implementation](https://stripe.com/docs/api/webhook_endpoints).  
    If you wish to add features you are completely free to do so.

<!--  -->

-   Fiber
    Test

-   Supabase
    Test

-   HTMX
    Remember that with htmx you'll also need to handle any request you make in the backend with a fiber route.

-   Vite
    To make sure that vite takes all the additional js scripts imported in any of your html pages you'll need to be sure to pass the script tag with `type="module"`, otherwise it won't be passed to the production build

## Notes

If running a multi-follder go environment, at the main root you will need to add a `go.mod` file with the go version and the path to handle these environments, example:

```work
go 1.20

use (
    ./environment1
    ./environment2
)
```

Notes:

-   htmx
    We use htmx CDN on html page not good i production, how to solve?  
    Minification eval error, evaluate

-   Vite
    check vite dont add CDNs to build,, try add htmx as static in production too
    Error: empty chunk, what does it mean, is it affecting the build? If not, can you silenc it?

-   To check
    https://sujit-baniya.gitbook.io/fiber-boilerplate/additional-libraries

Needed:

-   supebase ssl
-   backend integration stuff
-   superbase postgress utilities
-   check Fiber middelwares
-   Login with third party auth
-   user session token management utils
-   check chmod works for everyone or better solution
-   stripe webhooks, idempotency and verify signature
-   deployment flow test
-   Integrate [netadata](https://www.netdata.cloud/)

Optionals:

-   dns?
-   how to scale?
