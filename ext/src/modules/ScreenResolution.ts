import Module from "../interfaces/Module";

/**
 * Module which stores the Screen Resolution and Dimensions
 */
export default class ScreenResolution implements Module {

    /**
     * Name of the module
     */
    public readonly name = "Screen";

    /**
     * Difference between ClientX and ScreenX
     */
    private dx = 0;

    /**
     * Difference between ClientY and ScreenY
     */
    private dy = 0;


    /**
     * Will be called, when the Captcha is rendered
     * registers handler to get dx and dy
     */
    public start(): void {
        document.addEventListener('mousemove', this.calcDelta.bind(this), { once: true });
    }

    /**
     * Will be called, when the Captcha is solved
     * does nothing
     */
    public stop(): void { }

    /**
     * Calculates dx and dy
     * @param ev Event of MouseMove
     */
    private calcDelta(ev: MouseEvent): void {
        this.dx = ev.screenX - ev.clientX;
        this.dy = ev.screenY - ev.clientY;
    }

    /**
     * returns the Screen Resolution and Dimensions
     * @returns Screen Resolution and Dimensions
     */
    public getCollectedData(): Object {
        return {
            width: screen.availWidth,
            height: screen.availHeight,
            density: screen.pixelDepth,
            innerw: window.innerWidth,
            innerh: window.innerHeight,
            deltax: this.dx,
            deltay: this.dy,
        }
    }

}
