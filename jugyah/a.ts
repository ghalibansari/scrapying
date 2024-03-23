import crypto from "node:crypto";
import fs from "node:fs";
import playwright from "playwright";

//https://www.propi.in/buildings_directory/1-10
const baseUrl = "https://www.propi.in/buildings_directory/";
// let currentPage = 1;
const lastPage = 1300;
const duplicateData: Record<string, any> = {};
const failedPages: number[] = [];

const main = async () => {
  const data = await readJsonData();

  const browser = await playwright.chromium.launch({ headless: false });
  const context = await browser.newContext();
  const page = await context.newPage();

  // for await from 1 to lastPage
  const arr = Array.from({ length: lastPage }, (_, i) => i + 1);
  console.log("arr", arr.length);
  for await (const pageNumber of arr) {
    // start from 1217
    // if (pageNumber < 1217) {
    //   continue;
    // }
    try {
      // wait for random time between 0 to 4000 ms
      await new Promise((resolve) => {
        setTimeout(resolve, Math.floor(Math.random() * 9000));
      });
      // await fetchPage(pageNumber, page, data);
      console.log("pageNumber", pageNumber);
    } catch (e) {
      failedPages.push(pageNumber);
      console.error("error ========>>>> ", pageNumber, e);
    }
  }
  // await writeJsonData(data, "data.json");
  // await writeJsonData(duplicateData, "duplicateData.json");
  // await writeJsonData({ failedPages }, "failedPages.json");

  // await fetchPage(1, page, data);
  // console.log("data", data);

  await browser.close();
};

const fetchPage = async (
  pageNumber: number,
  page: playwright.Page,
  data: Record<string, any>
) => {
  const url = `${baseUrl}${pageNumber}-10`;

  await page.goto(url);
  await page.getByText("Showing").first().waitFor({ timeout: 2_000 });

  const rows = await page.$$(".row");

  // rows.length = 3; // remove this line

  await handleRows(rows, data, pageNumber);
};

const handleRows = async (
  rows: playwright.ElementHandle<SVGElement | HTMLElement>[],
  data: Record<string, any>,
  pageNumber: number
) => {
  for await (const row of rows) {
    await handleRow(row, data, pageNumber);
  }
};

const handleRow = async (
  row: playwright.ElementHandle<SVGElement | HTMLElement>,

  data: Record<string, any>,
  pageNumber: number
) => {
  const cards = await row.$$(".card");
  console.log(`on page ${pageNumber},  found ${cards.length} cards`);
  for await (const card of cards) {
    const cardData = await parseCardData(card);

    if (data[cardData.id]) {
      // console.log(
      //   `#### duplicate card found on page ${pageNumber}, id: ${
      //     cardData.id
      //   } title: ${cardData.title} count: ${data[cardData.id].count}`
      // );
      // data[cardData.id].count += 1;
      if (duplicateData[`${cardData.id}-${cardData.title}`]) {
        duplicateData[`${cardData.id}-${cardData.title}`] += 1;
      }
      duplicateData[`${cardData.id}-${cardData.title}`] = 1;
    } else {
      data[cardData.id] = cardData;
    }
  }
};

const parseCardData = async (
  card: playwright.ElementHandle<SVGElement | HTMLElement>
) => {
  const [anchorTitleEle, shortAddressEle, mainImageEle, shortBhkEle] =
    await Promise.all([
      card.$$("a.text-default"),
      card.$("h6"),
      card.$$("a.btn"),
      card.$$("span.text-muted"),
    ]);

  const [title, detailLink, shortAddress, mainImgLink, shortBhk] =
    await Promise.all([
      anchorTitleEle[1]?.textContent(),
      anchorTitleEle[1]?.getAttribute("href"),
      shortAddressEle?.textContent().then((x) => x?.trim()),
      mainImageEle[0]?.getAttribute("href"),
      shortBhkEle[1]?.textContent(),
    ]);

  const id = generateMD5(JSON.stringify(`${title}${shortAddress}${shortBhk}`));

  return {
    id,
    title,
    detailLink: `https://www.propi.in${detailLink}`,
    shortAddress,
    mainImgLink,
    shortBhk,
  };
};

const readJsonData = async () => {
  const data = fs.readFileSync("data.json");
  return JSON.parse(data.toString());
};

const writeJsonData = async (data: any, fileName: string) => {
  await fs.promises.writeFile(fileName, JSON.stringify(data, null, 2));
};

function generateMD5(input: string): string {
  const hash = crypto.createHash("md5");
  hash.update(input);
  return hash.digest("hex");
}

main();
