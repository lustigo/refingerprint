import Module from "../interfaces/Module";

/**
 * Module which stores Information about the Browser
 */
export default class BrowserInfo implements Module {

    /**
    * Name of the module
    */
    public readonly name = "Browser";


    /**
    * Will be called, when the Captcha is rendered
    * does nothing
    */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public start(): void { }

    /**
    * Will be called, when the Captcha is solved
    * does nothing
    */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public stop(): void { }


    /**
     * Gets the PluginNames of the Browser
     * @returns Array of Pluginnames
     */
    private getPlugins(): Array<string> {
        const plugins = navigator.plugins;
        const names = new Array<string>();

        for (const plugin of plugins) {
            names.push(plugin.name);
        }
        return names;
    }

    /**
     * Gets the supported Mime-Types of the Browser
     * @returns Array of Mime-Types
     */
    private getMimeTypes(): Array<string> {
        const types = navigator.mimeTypes;
        const names = new Array<string>();

        for (const type of types) {
            names.push(type.type);
        }
        return names;
    }

    /**
     * returns the Information
     * @returns Browser Information
     */
    public getCollectedData(): Record<string, any> {
        return {
            cookieEnabled: navigator.cookieEnabled,
            lang: navigator.languages,
            cpuNo: navigator.hardwareConcurrency,
            mimeTypes: this.getMimeTypes(),
            platform: navigator.platform,
            userAgent: navigator.userAgent,
            dnt: navigator.doNotTrack,
            plugins: this.getPlugins(),
            timezone: new Date().getTimezoneOffset(),
        };
    }


}
