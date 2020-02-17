import Module from "./interfaces/Module";

/**
 * Data Collector
 */
export default class Collector {

    /**
     * Modules that should be initialized when a new Captcha will be rendered
     */
    private modules: Array<{ new(): Module }>;

    /**
     * Module Instances for the next or current Captcha
     */
    private collectors: Array<Module> = new Array();


    /**
     * Initializes the Constructor with the data collection modules
     * Add the Listener to the Document
     */
    public constructor(modules: Array<{ new(): Module }>) {
        this.modules = modules;
        this.reset();
        document.getRootNode().addEventListener("captcha-created", this.start.bind(this));
        document.getRootNode().addEventListener("captcha-solved", this.stop.bind(this));
    }

    /**
     * Will be called when the Captcha is rendered.
     * Will start all Modules
     */
    public start() {
        this.collectors.forEach(col => col.start());
    }

    /**
     * Will be called when the Captcha was solved.
     * Will stop all Modules and collect the data
     */
    public stop() {
        this.collectors.forEach(col => col.stop());
        const data = this.collectData();
        this.saveData(data);
        this.reset();
    }

    /**
     * Collects the data gathered by the Modules and builds a single Object
     */
    private collectData(): Object {
        const data = Object.create(null);
        this.collectors.forEach(col => data[col.name] = col.getCollectedData());
        return data;
    }

    /**
     * Resets the Collectors after a Captcha is solved 
     * or for the first time during Construction.
     */
    private reset() {
        this.collectors = this.modules.map(Mod => new Mod());
    }

    /**
     * Send the collected data to the background script to save it on disk
     * @param data Data to save
     */
    private saveData(data: Object) {
        if (chrome) {
            chrome.runtime.sendMessage(data);
        } else {
            browser.runtime.sendMessage({
                message: data
            });
        }
    }

}
