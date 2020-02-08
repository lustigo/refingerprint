import { WidgetType } from "../enums/WidgetType";
import CaptchaDescription from "./CaptchaDescription";
import DropdownDescription from "./DropdownDescription";
import MatrixDescription from "./MatrixDescription";
import TextFieldDescription from "./TextFieldDescription";
import TextWidgetDescription from "./TextWidgetDescription";
import VectorDescription from './VectorDescription';
/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
    readonly required: boolean;
    readonly structure: CaptchaDescription | DropdownDescription | MatrixDescription | TextFieldDescription | TextWidgetDescription | VectorDescription;
    readonly type: WidgetType;
    // eslint-disable-next-line semi
}
