import Module from "../interfaces/Module";

/**
 * Module which collects the SessionId
 */
export default class Id implements Module {

    /**
     * Name of the module
     */
    public readonly name = "SessionId";

    /**
     * Id of the Session
     */
    private id = "UNKNOWN";

    /**
     * Gets the Id of the Session.
     */
    public constructor() {
        this.getId();
    }

    /**
     * Will be called, when the Captcha is rendered
     * Check if the id is set, if not, try again
     */
    public start(): void {
        if (this.id == "UNKNOWN") {
            this.getId();
        }
    }

    /**
    * Will be called, when the Captcha is solved
    * does nothing
    */
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    public stop(): void { }

    /**
     * Returns the SessionId
     * @return Id of the Session
     */
    public getCollectedData(): string {
        return this.id;
    }

    /**
     * Gets the SessionId from the DOM and stores it in this.id
     */
    private getId(): void {
        const sessionId = document.getElementById("sessionId");
        this.id = sessionId ? sessionId.innerHTML : "UNKNOWN";
    }

}
