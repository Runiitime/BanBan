import type { IMap } from "../structure";
import type { IColumn } from "./IColumn.ts";

export interface IBoard {
  id: string;
  title: string;
  description?: string;
  columns?: IMap<IColumn>;
}
