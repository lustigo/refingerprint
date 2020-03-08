/**
 * Helper Functions
 */

/**
* Returns a Promise, that will be called after time
* @param time Time to wait in ms
*/
async function delay(time: number): Promise<void> {
    return new Promise((resolve, _) => {
        setTimeout(resolve, time);
    });
}

/**
 * Checks if the given Rectangle collides with the given point
 * @param rx X Coordinate of Rectangle
 * @param ry Y Coordinate of Rectangle
 * @param rw width of Rectangle
 * @param rh height of Rectangle
 * @param x X Coordinate of Point
 * @param y Y Coordinate of Point
 * @return wether there is a collision or not
 */
function checkRectCollision(rx: number, ry: number, rw: number, rh: number, x: number, y: number): boolean {
    return (rx <= x && rx + rw >= x) && (ry + rh >= y && ry <= y);
}

export { delay };
