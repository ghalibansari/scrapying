import fs from "node:fs";

const readJsonData = async () => {
  const data = fs.readFileSync("agent.json");
  return JSON.parse(data.toString());
};

const main = async () => {
  const data = await readJsonData();
  // console.log(data);
  let lagestCount = 0;

  for (const key in data) {
    console.log(`key: ${key}, pageNumber: ${data[key].pageNumber}`);
    if (data[key].pageNumber > lagestCount) {
      lagestCount = data[key].pageNumber;
    }
  }

  console.log(`largest count: ${lagestCount}`);
};

main();
