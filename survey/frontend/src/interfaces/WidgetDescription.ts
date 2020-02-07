import { WidgetType } from "../enums/WidgetType";
import MatrixDescription from "./MatrixDescription";
import DropdownDescription from "./DropdownDescription";
/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
  readonly type: WidgetType;
  readonly structure: MatrixDescription | DropdownDescription;
  readonly required: boolean;
}
