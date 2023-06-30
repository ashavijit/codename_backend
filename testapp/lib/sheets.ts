import { google, sheets_v4 } from "googleapis";
import serviceAccountKeyFile from "../keys.json"
  
  async function _getGoogleSheetClient() {
    const auth = new google.auth.GoogleAuth({
      credentials: serviceAccountKeyFile,
      scopes: ['https://www.googleapis.com/auth/spreadsheets'],
    });
    const authClient = await auth.getClient();
    return google.sheets({
      version: 'v4',
      auth: authClient,
    });
  }
  
  async function _readGoogleSheet(googleSheetClient: sheets_v4.Sheets, sheetId: string, tabName: string, range: string) {
    const res = await googleSheetClient.spreadsheets.values.get({
      spreadsheetId: sheetId,
      range: `${tabName}!${range}`,
    });
  
    return res.data.values;
  }
  
  async function _writeGoogleSheet(googleSheetClient: sheets_v4.Sheets, sheetId: string, tabName: string, range: string, data: any[]) {
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

  export {_getGoogleSheetClient, _readGoogleSheet, _writeGoogleSheet}