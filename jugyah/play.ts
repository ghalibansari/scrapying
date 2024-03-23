import playwright from "playwright";

const main = async () => {
  const chrome = await playwright.chromium.launch({ headless: false });
  const context = await chrome.newContext();
  const page = await context.newPage();
  await page.goto("https://www.google.com");
  // search for "playwright"
  await page.fill('input[name="q"]', "playwright");
  await page.keyboard.press("Enter");

  //scroll down
  await page.evaluate(() => {
    window.scrollBy(0, window.innerHeight);
  });
  await page.screenshot({ path: "example.png" });

  console.log(">>>>>>> END <<<<<<<<");
};

main();
