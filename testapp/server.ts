/* eslint-disable @typescript-eslint/no-unused-vars */
import express, { Application, Request, Response } from 'express';
import { google, sheets_v4 } from 'googleapis';
import axios from 'axios';
import {
      _getGoogleSheetClient,
      _writeGoogleSheet,
} from './lib/sheets';
import { parse } from 'json2csv';
import fs from 'fs';

const app: Application = express();
const port = 3002;
const TOKEN =
      'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4';
const URL = 'http://localhost:8080/admin';
const sheetId = '19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s';
const tabName = 'Sheet3';
const range = 'A1:F6';

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get(
      '/',
      async (req: Request, res: Response): Promise<Response> => {
            return res.status(200).send({
                  message: 'Hello World!',
            });
      }
);

const headers = {
      'Content-Type': 'application/json',
      Authorization: TOKEN,
};

axios.get(URL, { headers })
      .then(async response => {
            const JSONDATA = JSON.parse(
                  JSON.stringify(response.data)
            );
            console.log(JSONDATA);
      })
      .catch(error => {
            console.log(error);
      });

app.listen(port, (): void => {
      console.log(`Connected successfully on port ${port}`);
});
