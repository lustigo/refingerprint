import { WidgetType } from "../enums/WidgetType";
import Keyed from "./Keyed";
/**
 * Describes the Data of a filled-in widget.
 */
export default interface WidgetData {
    readonly type: WidgetType;
    value: Keyed<number|string>;
    // eslint-disable-next-line semi
}
