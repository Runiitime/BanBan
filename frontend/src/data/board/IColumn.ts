import type { ICard } from "./ICard.ts";

export interface IColumn {
  id: string;
  title: string;
  description?: string;

  position: number;

  created_at?: string;
  updated_at?: string;

  color: string;
  cards: ICard[];
}
