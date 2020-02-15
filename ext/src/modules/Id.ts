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
    private id: string;

    /**
     * Gets the Id of the User.
     */
    constructor() {
        const userId = document.getElementById("userId");
        this.id = userId ? userId.innerHTML : "UNKNOWN";
    }

    /**
     * Will be called, when the Captcha is rendered
     * does nothing
     */
    start(): void { }

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

}
