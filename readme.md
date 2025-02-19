# Gylym: Kazakhstan IT Courses Platform

## About the Project
**Gylym** is an educational platform for learning programming. The project aims to create a convenient learning environment with courses, interactive assignments, and tests.

## Technologies
### **Frontend**
- HTML + CSS
- React + Vite
- MUI (Material UI)
- Axios
- React Router

### **Backend**
- Golang
- PostgresSQL
- net/http
- goose (Migrations)
- testify, go-mock (Testing)
- Redis
- Docker
- GitHub Actions (CI/CD)
- Swagger (API Documentation)

## Key Features
- **Forum**: A discussion forum where students can ask and answer questions on various IT topics (Frontend, Backend, etc.).
- **Course Catalog**: View a list of courses with filtering and search functionality.
- **Dynamic Course Data Loading**:
  - The course page loads information via API.
  - Course sections include textual materials, quizzes, and practical tasks.
  - **Integrated IDE**: Practical assignments feature an embedded coding environment for hands-on experience.
  - The course page loads information via API.
  - Course sections include textual materials, quizzes, and practical tasks.
- **Authentication**:
  - Login and registration through API.
  - Uses JWT tokens.
- **User Page**:
  - View user information (email, name, avatar, course progress).
- **Interactive Interface**:
  - Sidebar for topic navigation.
  - Semi-header with course title and future progress bar.


## CI/CD
We use **GitHub Actions** for continuous integration and deployment:
- **Linting**: Automatically lints the codebase using golangci-lint. 
- **Testing**: Automatically runs tests using go test.
- **Migrations**: Automatically generates and runs database migrations.
- **Deployment**: Automatically deploys the latest version to Railway after a successful build.
  The project has been deployed using Docker on Railway and is available at [https://gylym-base.up.railway.app/](https://gylym-base.up.railway.app/).

## API
API is documented in Swagger. To access, open `/swagger/index.html` after starting the backend.

## Contact
Authors: [**Rassul Turutlov**](https://github.com/Rasikrr), [**Amir Kurmanbekov**](https://github.com/Arh0rn), [**Kaminur Orynbek**](https://github.com/kaminurorinbek)

