import Module from "../interfaces/Module";

/**
 * Module which stores the Screen Resolution
 */
export default class ScreenResolution implements Module {

    /**
    * Name of the module
    */
    public readonly name = "Screen";


    /**
    * Will be called, when the Captcha is rendered
    * does nothing
    */
    public start(): void { }

    /**
    * Will be called, when the Captcha is solved
    * does nothing
    */
    public stop(): void { }

    /**
     * returns the Screen Resolution
     * @returns Screen Resolution
     */
    public getCollectedData(): Object {
        return {
            width: screen.availWidth,
            height: screen.availHeight,
            density: screen.pixelDepth,
        }
    }

}
