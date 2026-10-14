export interface ICard {
  id: string;
  columnId: string;
  title: string;
  description: string;
  status: string;
  created_at?: string;
  updated_at?: string;
  completed_at?: string;
}
