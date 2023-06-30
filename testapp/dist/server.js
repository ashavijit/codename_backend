"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const axios_1 = __importDefault(require("axios"));
const app = (0, express_1.default)();
const port = 3000;
var TOKEN = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjg4NjY2OTYwfQ.gKygcoudX2sDXo54kPCTXNnJtiwUczviCSmhgJqtAU4";
// Body parsing Middleware
app.use(express_1.default.json());
app.use(express_1.default.urlencoded({ extended: true }));
const URL = "http://localhost:8080/admin";
const headers = {
    "Content-Type": "application/json",
    Authorization: TOKEN,
};
app.get("/", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    return res.status(200).send({
        message: "Hello World!",
    });
}));
app.get("/admin", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const response = yield axios_1.default.get(URL, { headers });
        if (response.status === 400) {
            return res.status(400).json({ error: "Bad Request" });
        }
        else if (response.status === 401) {
            return res.status(401).json({ error: "Unauthorized" });
        }
        return res.status(200).json(response.data);
        // save the data to excel sheet
    }
    catch (error) {
        console.error("Error fetching data:", error);
        return res.status(500).json({ error: "Failed to fetch data" });
    }
}));
try {
    app.listen(port, () => {
        console.log(`Connected successfully on port ${port}`);
    });
}
catch (error) {
    console.error(`Error occured`);
}
