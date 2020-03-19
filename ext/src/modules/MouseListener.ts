import Module from "../interfaces/Module";

/**
 * Module which tracks the Mouse Events
 */
export default abstract class MouseListener implements Module {

    /**
     * Handlers to remove the Listener
     */
    protected removeHandler: Array<() => void> = new Array<() => void>();

    /**
     * Name of the Module
     */
    public name = "Mouse";

    /**
     * will be called when the Captcha is rendered
     * does nothing
     */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public start(): void { }

    /**
     * will be called when the Captcha is solved
     * does nothing
     */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public stop(): void { }

    /**
     * returns the collected Data
     * @returns collected Data
     */
    public abstract getCollectedData(): Array<any>

    /**
     * Executes and removes all RemoveHandlers
     */
    protected stopTracking(): void {
        while (this.removeHandler.length > 0) {
            const f = this.removeHandler.pop();
            if (f) {
                f();
            }
        }
    }
    // eslint-disable-next-line semi
}
