import MouseData from "./MouseData";

/**
 * Represents the Key which was pressed
 */
enum MouseKey {
    LEFT = 0,
    RIGHT = 2,
    MIDDLE = 1,
}

/**
 * Represents a MouseClick Event
 */
interface ClickData extends MouseData {
    key: MouseKey;
    release: boolean; //Wheter it is a Press or Release Event
}

export { MouseKey, ClickData };
