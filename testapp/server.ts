import express, { Application, Request, Response } from "express";
import { google, sheets_v4 } from "googleapis";
import axios from "axios";



const app: Application = express();
const port = 3000;
var TOKEN =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4";
// Body parsing Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
const URL = process.env.BASE_URL
const headers = {
  "Content-Type": "application/json",
  Authorization: TOKEN,
};
app.get("/", async (req: Request, res: Response): Promise<Response> => {
  return res.status(200).send({
    message: "Hello World!",
  });
});

app.get("/admin", async (req, res) => {
  try {
    const response = await axios.get(URL, { headers });
    if (response.status === 400) {
      return res.status(400).json({ error: "Bad Request" });
    } else if (response.status === 401) {
      return res.status(401).json({ error: "Unauthorized" });
    }
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
