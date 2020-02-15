/**
 * A module that collects data while the Captcha is rendered
 */
export default interface Module {

    /**
     * Name of the module
     */
    readonly name: string;

    /**
     * Will be called, when the Captcha appears on the page.
     * The collection should be started.
     */
    start(): void;

    /**
     * Will be called, when the Captcha was solved. 
     * The collection should be stopped.
     */
    stop(): void;

    /**
     * This method is a Getter for the collected data.
     */
    getCollectedData(): Object | string | number;

}