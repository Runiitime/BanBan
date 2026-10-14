export interface IColumn {
  id: string;
  boardId: string;
  title: string;
  description?: string;

  position: number;

  created_at?: string;
  updated_at?: string;

  color: string;

  itemsOrder: string[];
}
