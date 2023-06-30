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
const googleapis_1 = require("googleapis");
const keys_json_1 = __importDefault(require("../keys.json"));
const sheetId = '19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s';
const tabName = 'codename_app';
const range = "A:E";
function _getGoogleSheetClient() {
    return __awaiter(this, void 0, void 0, function* () {
        const auth = new googleapis_1.google.auth.GoogleAuth({
            credentials: keys_json_1.default,
            scopes: ['https://www.googleapis.com/auth/spreadsheets'],
        });
        const authClient = yield auth.getClient();
        return googleapis_1.google.sheets({
            version: 'v4',
            auth: authClient,
        });
    });
}
function _readGoogleSheet(googleSheetClient, sheetId, tabName, range) {
    return __awaiter(this, void 0, void 0, function* () {
        const res = yield googleSheetClient.spreadsheets.values.get({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
        });
        return res.data.values;
    });
}
function _writeGoogleSheet(googleSheetClient, sheetId, tabName, range, data) {
    return __awaiter(this, void 0, void 0, function* () {
        const res = yield googleSheetClient.spreadsheets.values.append({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
            valueInputOption: 'USER_ENTERED',
            requestBody: {
                values: data,
            },
        });
        return res.data;
    });
}
