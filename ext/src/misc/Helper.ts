/**
 * Helper Functions
 */

/**
* Returns a Promise, that will be called after time
* @param time Time to wait in ms
*/
async function delay(time: number): Promise<void> {
    return new Promise((resolve) => {
        setTimeout(resolve, time);
    });
}

export { delay };
