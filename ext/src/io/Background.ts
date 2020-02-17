/**
 * Disable the Shelf for Chrome (Footer that shows what was downloaded)
 */
if (chrome.downloads.setShelfEnabled) {
    chrome.downloads.setShelfEnabled(false);
}

/**
 * When a message is received, the Data should be encoded in a Blob.
 * The Blob will be encoded in an URL and that URL will be downloaded.
 */
function downloadData(req: Object) {
    const blob = new Blob([JSON.stringify(req)]);
    const url = window.webkitURL.createObjectURL(blob);
    chrome.downloads.download({
        url,
        // (Unix-)Timestamp as Filename
        filename: new Date().getTime().toString() + ".json",
        // Do not ask where to save file
        saveAs: false
    });
}

/**
 * When a message is received from the Content Script
 */
chrome.runtime.onMessage.addListener(downloadData);
