"use strict";
// import express, { Application, Request, Response } from "express";
// import { google, sheets_v4 } from "googleapis";
// import axios from "axios";
// import {_getGoogleSheetClient , _readGoogleSheet, _writeGoogleSheet} from "./lib/sheets"
// import {parse} from "json2csv"
// import * as fs from "fs";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// const app: Application = express();
// const port = 3000;
// var TOKEN =
//   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4";
// // Body parsing Middleware
// app.use(express.json());
// app.use(express.urlencoded({ extended: true }));
// const URL= "http://localhost:8080/admin"
// const headers = {
//   "Content-Type": "application/json",
//   Authorization: TOKEN,
// };
// const sheetId = "19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s";
// const tabName = "Sheet1";
// // const range = "A1:B2";
// // const data = [
// //   ["Name", "Age"],
// //   ["Alice", "20"],
// //   ["Bob", "25"],
// // ];
// // async function writeToGoogleSheet() {
// //   const client = await _getGoogleSheetClient();
// //   const sheetId = '19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s';
// //   const tabName = 'Sheet1';
// //   const range = 'A1:B2';
// //   data;
// //   try {
// //     const result = await _writeGoogleSheet(client, sheetId, tabName, range, data);
// //     console.log(result);
// //   } catch (error) {
// //     console.error('Error:', error);
// //   }
// // }
// // writeToGoogleSheet().catch(console.error);
// app.get("/", async (req: Request, res: Response): Promise<Response> => {
//   return res.status(200).send({
//     message: "Hello World!",
//   });
// });
// app.get("/admin", async (req, res) => {
//   try {
//     const response = await axios.get(URL, { headers });
//     if (response.status === 400) {
//       return res.status(400).json({ error: "Bad Request" });
//     } else if (response.status === 401) {
//       return res.status(401).json({ error: "Unauthorized" });
//     }
//     const JSONDATA = JSON.parse(JSON.stringify(response.data));
//     const fields = Object.keys(JSONDATA[0]);
//     const CsvDATA = parse(JSONDATA , {fields});
//     const _csvPath = "./csvdata.csv";
//     fs.writeFileSync(_csvPath, CsvDATA);
//     const client = await _getGoogleSheetClient();
//     const _data = fs.readFileSync("./csvdata.csv");
//     const sheetId = "19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s";
//     const tabName = "Sheet1";
//     const range = "A1:F6";
//     const data =_data.toString().split("\n").map((el)=>el.split(","))
//     const result = await _writeGoogleSheet(client, sheetId, tabName, range, data);
//     fs.unlinkSync(_csvPath);
//     return res.status(200).json({ message: "Success" });
//   } catch (error) {
//     console.log(error);
//     return res.status(500).json({ error: "Internal Server Error" });
//   }
// });
// try {
//   app.listen(port, (): void => {
//     console.log(`Connected successfully on port ${port}`);
//   });
// } catch (error) {
//   console.error(`Error occured`);
// }
const express_1 = __importDefault(require("express"));
const axios_1 = __importDefault(require("axios"));
const sheets_1 = require("./lib/sheets");
const app = (0, express_1.default)();
const port = 3000;
const TOKEN = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4";
const URL = "http://localhost:8080/admin";
const sheetId = "19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s";
const tabName = "Sheet1";
const range = "A1:F6";
app.use(express_1.default.json());
app.use(express_1.default.urlencoded({ extended: true }));
app.get("/", async (req, res) => {
    return res.status(200).send({
        message: "Hello World!",
    });
});
app.get("/admin", async (req, res) => {
    try {
        const response = await axios_1.default.get(URL, { headers: { "Content-Type": "application/json", Authorization: TOKEN } });
        if (response.status === 400) {
            return res.status(400).json({ error: "Bad Request" });
        }
        else if (response.status === 401) {
            return res.status(401).json({ error: "Unauthorized" });
        }
        const jsonData = JSON.parse(JSON.stringify(response.data));
        const fields = Object.keys(jsonData[0]);
        const client = await (0, sheets_1._getGoogleSheetClient)();
        const data = []; // Initialize data array
        // Add header row with field names
        const headerRow = fields.map((field) => [field]);
        data.push(headerRow);
        for (let i = 0; i < jsonData.length; i++) {
            const row = fields.map((field) => jsonData[i][field]);
            data.push(row);
        }
        const range = `A1:${String.fromCharCode(65 + fields.length)}${jsonData.length + 1}`; // Calculate the range dynamically starting from column A
        const result = await (0, sheets_1._writeGoogleSheet)(client, sheetId, tabName, range, data);
        return res.status(200).json({ message: "Success" });
    }
    catch (error) {
        console.log(error);
        return res.status(500).json({ error: "Internal Server Error" });
    }
});
app.listen(port, () => {
    console.log(`Connected successfully on port ${port}`);
});
