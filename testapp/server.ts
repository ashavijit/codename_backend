import express, { Application, Request, Response } from "express";
import axios from "axios";

const app: Application = express();
const port = 3000;
var TOKEN = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4"
// Body parsing Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
const URL = "http://localhost:8080/admin";
const headers = {
    "Content-Type": "application/json",
    Authorization: TOKEN, //the token is a variable which holds the token
};
app.get("/", async (req: Request, res: Response): Promise<Response> => {
    return res.status(200).send({
        message: "Hello World!",
    });
});

app.get("/admin", async (req, res) => {
    try {
      const response = await axios.get(URL, { headers: headers });
      return res.status(200).json(response.data);
    } catch (error) {
      console.error("Error fetching data:", error);
      return res.status(500).json({ error: "Failed to fetch data" });
    }
  });
try {
    app.listen(port, (): void => {
        console.log(`Connected successfully on port ${port}`);
    });
} catch (error) {
    console.error(`Error occured`);
}
