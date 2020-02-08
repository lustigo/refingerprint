import { WidgetType } from "../enums/WidgetType";
import CaptchaDescription from "./CaptchaDescription";
import DropdownDescription from "./DropdownDescription";
import MatrixDescription from "./MatrixDescription";
import TextFieldDescription from "./TextFieldDescription";
import TextWidgetDescription from "./TextWidgetDescription";
/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
    readonly type: WidgetType;
    readonly structure: CaptchaDescription | DropdownDescription | MatrixDescription | TextFieldDescription | TextWidgetDescription;
    readonly required: boolean;
    // eslint-disable-next-line semi
}
