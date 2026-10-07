export interface ICard {
  id: string;
  title: string;
  description: string;
  status: string;
  position: number;
  created_at?: string;
  updated_at?: string;
  completed_at?: string;
}
