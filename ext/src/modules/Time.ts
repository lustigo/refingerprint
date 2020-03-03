import Module from "../interfaces/Module";

/**
 * Module which stores the Start- and Endtime
 */
export default class Time implements Module {
    /**
     * Name of the module
     */
    public readonly name = "Time";

    /**
     * Start-Time
     */
    private startTime: Date = new Date();

    /**
     * End-Time
     */
    private endTime: Date = new Date();

    /**
     * Will be called, when the Captcha is rendered
     * Stores the time
     */
    start(): void {
        this.startTime = new Date();
    }

    /**
     * Will be called, when the Captcha is solved
     * Stores the time
     */
    stop(): void {
        this.endTime = new Date();
    }

    /**
     * Returns the Times
     * @return Start and End-Time as Unix Milliseconds
     */
    getCollectedData(): Object {
        return {
            start: this.startTime.getTime(),
            end: this.endTime.getTime(),
        };
    }

}
