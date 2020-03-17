import jsSHA from "jssha";

/**
 * Disable the Shelf for Chrome (Footer that shows what was downloaded)
 */
if (chrome.downloads.setShelfEnabled) {
    chrome.downloads.setShelfEnabled(false);
}

/**
 * Remove the Google Cookie, so that there should be a Captcha Quest
 */
function onCookieUpdate(): void {
    chrome.cookies.remove({ name: "NID", "url": "https://google.com" });
    chrome.cookies.remove({ name: "ANID", "url": "https://google.com" });
}


/**
 * When a message is received, the Data should be encoded in a Blob.
 * The Blob will be encoded in an URL and that URL will be downloaded.
 */
function downloadData(req: Record<string, any>): void {
    const data = JSON.stringify(req);
    const blob = new Blob([data]);
    const url = window.webkitURL.createObjectURL(blob);

    // Hash of Data (only first 10 chars) as Filename
    const hasher = new jsSHA("SHA-512", "TEXT");
    hasher.update(data);
    const filename = hasher.getHash("B64").substr(0, 10).replace("/", "").replace("\\", "") + ".json";

    chrome.downloads.download({
        url,

        filename,

        // Do not ask where to save file
        saveAs: false,

        // If file name does already exist, do not overwrite
        conflictAction: "uniquify",
    });
}

/**
 * When a message is received from the Content Script
 */
chrome.runtime.onMessage.addListener(downloadData);

/**
 * When the cookies are changed, remove the Goolge Cookies
 */
chrome.cookies.onChanged.addListener(onCookieUpdate);
