import { WidgetType } from "../enums/WidgetType";
import MatrixDescription from "./MatrixDescription";
/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
  readonly type: WidgetType;
  readonly structure: MatrixDescription;
  readonly required: boolean;
}
