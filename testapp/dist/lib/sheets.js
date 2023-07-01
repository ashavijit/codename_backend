'use strict';
var __importDefault =
      (this && this.__importDefault) ||
      function (mod) {
            return mod && mod.__esModule ? mod : { default: mod };
      };
Object.defineProperty(exports, '__esModule', { value: true });
exports._clearGoogleSheetData =
      exports._getExistingSheetData =
      exports._writeGoogleSheet =
      exports._readGoogleSheet =
      exports._getGoogleSheetClient =
            void 0;
const googleapis_1 = require('googleapis');
const keys_json_1 = __importDefault(require('../keys.json'));
async function _getGoogleSheetClient() {
      const auth = new googleapis_1.google.auth.GoogleAuth({
            credentials: keys_json_1.default,
            scopes: ['https://www.googleapis.com/auth/spreadsheets'],
      });
      const authClient = await auth.getClient();
      return googleapis_1.google.sheets({
            version: 'v4',
            auth: authClient,
      });
}
exports._getGoogleSheetClient = _getGoogleSheetClient;
async function _readGoogleSheet(
      googleSheetClient,
      sheetId,
      tabName,
      range
) {
      const res = await googleSheetClient.spreadsheets.values.get({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
      });
      return res.data.values;
}
exports._readGoogleSheet = _readGoogleSheet;
async function _writeGoogleSheet(
      googleSheetClient,
      sheetId,
      tabName,
      range,
      data
) {
      const res = await googleSheetClient.spreadsheets.values.append({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
            valueInputOption: 'USER_ENTERED',
            requestBody: {
                  values: data,
            },
      });
      return res.data;
}
exports._writeGoogleSheet = _writeGoogleSheet;
async function _getExistingSheetData(
      googleSheetClient,
      sheetId,
      tabName,
      range
) {
      const res = await googleSheetClient.spreadsheets.values.get({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
      });
}
exports._getExistingSheetData = _getExistingSheetData;
async function _clearGoogleSheetData(
      googleSheetClient,
      sheetId,
      tabName,
      range
) {
      const res = await googleSheetClient.spreadsheets.values.clear({
            spreadsheetId: sheetId,
            range: `${tabName}!${range}`,
      });
}
exports._clearGoogleSheetData = _clearGoogleSheetData;
