import { google, sheets_v4 } from "googleapis";
import serviceAccountKeyFile from "../keys.json"
const sheetId = '19nwije38ve_YkzMkwY8lmDr7jKkjnXUSN6K4KFGey3s'
const tabName = 'codename_app'
const range = "A:E"
  
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