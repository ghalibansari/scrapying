import fs from "node:fs";

const fileName = "agent";

const readJsonData = async () => {
  const data: Record<string, any> = fs.readFileSync(`${fileName}.json`);
  return JSON.parse(data.toString());
};

// convert json to csv

const convertJsonToCsv = async (data: Record<string, Record<string, any>>) => {
  const headers = Object.keys(data[Object.keys(data)[0]]);
  const csv = [headers.join(",")];

  for (const key in data) {
    const row = data[key];
    const values = headers.map((header) => row[header]);
    csv.push(values.join(","));
  }

  return csv.join("\n");
};

// write csv to file
const writeCsvToFile = async (csv: any) => {
  fs.writeFileSync(`${fileName}.csv`, csv);
};

const main = async () => {
  const data = await readJsonData();
  console.log("read data file, total records: ", Object.keys(data).length);

  const csv = await convertJsonToCsv(data);
  console.log("converted data to csv");

  await writeCsvToFile(csv);
  console.log("written csv to file");
};

main();
