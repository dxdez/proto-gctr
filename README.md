# Proto GO CRUD Task Rendering (GCTR)

Proto GCTR is a lightweight web application designed for experimenting with GO and HTMX. It simplifies the creation of dynamic web interfaces by combining the efficiency of GO with the interactivity of HTMX, without the need for bulky front-end libraries. This application makes use of basic net/http and CRUD functionality. This project is for experimental purposes only.

### Features

- **Chi Router:** Proto GCTR leverages the Chi router for GO, providing a robust and efficient routing system.
- **TailwindCSS:** Enjoy a sleek and modern UI design with TailwindCSS, ensuring a responsive and visually appealing user experience.
- **Templating with GO:** The application utilizes GO's templating capabilities for dynamic content generation, offering flexibility and simplicity in development.
- **HTMX Integration:** By integrating HTMX, Proto GCTR streamlines front-end functionality, allowing for seamless interactions without the overhead of heavy JavaScript frameworks.

Proto GCTR is perfect for building item lists, resembling traditional to-do applications, and serves as an excellent starting point for experimenting with GO and HTMX.

### Running Proto GCTR Locally

#### Prerequisites

Before running Proto GCTR locally, ensure that you have GO installed on your system. If GO is not installed, follow the official installation instructions for your operating system:

- [GO Installation Guide](https://golang.org/doc/install)

#### Getting Started

1. Clone the Proto GCTR repository to your local machine:

   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory:

   ```bash
   cd proto-gctr
   ```

3. Run the application using the `go run` command:

   ```bash
   go run .
   ```

4. Once the application is running, open your web browser and navigate to `http://localhost:8080` to access Proto GCTR.

### Running Application in Docker

The application runs within a docker container. See the official documentation for docker to set it up on your system [here](https://docs.docker.com/engine/install/ "here").

Execute the following command to run the application:
`````
docker compose up -d
`````

Execute the following command to shutdown:
`````
docker compose down -v
`````
