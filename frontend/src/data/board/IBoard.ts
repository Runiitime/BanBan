import type { IMap, MapI } from "../structure";
import type { IColumn } from "./IColumn.ts";

export interface IBoard {
  id: string;
  title: string;
  description?: string;
  columns: MapI<IColumn>;
}
