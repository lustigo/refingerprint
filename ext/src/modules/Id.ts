import Module from "../interfaces/Module";

/**
 * Module which collects the UserId
 */
export default class Id implements Module {

    /**
     * Name of the module
     */
    public readonly name = "UserId";

    /**
     * Id of the User
     */
    private id: string = "UNKNOWN";

    /**
     * Gets the Id of the User.
     */
    constructor() {
        this.getId();
    }

    /**
     * Will be called, when the Captcha is rendered
     * Check if the id is set, if not, try again
     */
    start(): void {
        if (this.id == "UNKNOWN") {
            this.getId();
        }
    }

    /**
    * Will be called, when the Captcha is solved
    * does nothing
    */
    stop(): void { }

    /**
     * Returns the UserId
     * @return Id of the User
     */
    getCollectedData(): string {
        return this.id;
    }

    /**
     * Gets the UserId from the DOM and stores it in this.id
     */
    private getId() {
        const userId = document.getElementById("userId");
        this.id = userId ? userId.innerHTML : "UNKNOWN";
    }

}
