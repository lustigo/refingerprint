import { WidgetType } from "../enums/WidgetType";
/**
 * Describes the structure of a Widget.
 */
export default interface WidgetDescription {
  readonly type: WidgetType;
  readonly structure: Object;
  readonly required: boolean;
}
